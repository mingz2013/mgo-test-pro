package services

import (
	"fmt"
	"github.com/mingz2013/mgo-test-pro/dao"
	"k8s.io/apimachinery/pkg/util/rand"
	"log"
	"sync"
	"time"
)

func FindUserData(waitGroup *sync.WaitGroup, tokens chan<- int, userId int) {
	defer waitGroup.Done()

	c := dao.UserDataC{}

	startTime := time.Now()

	defer func() {
		endTime := time.Now()

		interval := endTime.Sub(startTime)

		fmt.Println("interval: ", interval, int(interval))
		log.Println("log interval: ", interval, int(interval))

		tokens <- 1
	}()

	_, err := c.FindByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(data)

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
			go FindUserData(&waitGroup, tokens, findId)

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
