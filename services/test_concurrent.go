package services

import (
	"fmt"
	"mgo-test-pro/dao"
	"k8s.io/apimachinery/pkg/util/rand"
	"log"
	"sync"
	"time"
)

func RunOps(waitGroup *sync.WaitGroup, tokens chan<- int, userId int) {
	defer waitGroup.Done()

	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(4)

	logStr := "Find"
	if num == 1 {
		logStr = "insert"
	} else if num == 2 {
		logStr = "update"
	} else {
		logStr = "Find"
	}

	c := dao.UserDataC{}

	data := dao.NewUserData(userId)

	startTime := time.Now()

	defer func() {
		endTime := time.Now()

		interval := endTime.Sub(startTime)

		a := []interface{}{
			logStr, " | ", interval, " | ", int(interval),
		}

		fmt.Println(a...)
		log.Println(a...)

		tokens <- 1
	}()

	//_, err := c.FindByUserId(userId)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(data)

	if num == 1 {
		data.UserId *= 3
		c.Insert(data)
	} else if num == 2 {
		c.TestUpdate(userId, data)
	} else {

		_, err := c.FindByUserId(userId)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func TestConcurrent(count int, concurrentCount int, maxUserId, minUserId int) {

	tokens := make(chan int, concurrentCount+10)

	for i := 0; i < concurrentCount; i++ {
		tokens <- 1
	}

	//fmt.Println("tokens: ", tokens)

	var waitGroup sync.WaitGroup

	var i int

	i = 0

	var breakFor = false

	for {
		//log.Println("for...count...")
		select {
		case <-tokens:
			// 拿到一个token，
			waitGroup.Add(1)
			findId := rand.IntnRange(minUserId, maxUserId)
			fmt.Println("findId: ", findId)
			log.Println("findId: ", findId)
			go RunOps(&waitGroup, tokens, findId)

			i++

			if i >= count {
				log.Println("i > count...")
				breakFor = true
				break
			}

		case <-time.After(time.Microsecond * 1):
			break
		default:
			break

		}

		if breakFor {
			break
		}
	}

	//w:= make(chan int)

	//<-w
	//for i range <-tokens{
	//
	//}

	log.Println("befor wait..")
	waitGroup.Wait()

	log.Println("TestConcurrent....end...")

}
