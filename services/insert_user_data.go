package services

import (
	"fmt"
	"github.com/mingz2013/mgo-test-pro/dao"
	"log"
	"sync"
	"time"
)

var userDataC *dao.UserDataC

//func init() {
//	//userDataC = dao.NewUserDataC()
//}

func InsertUserData(waitGroup *sync.WaitGroup, tokens chan<- int, index int) {
	//log.Println("InsertUserData..")

	defer waitGroup.Done()

	data := dao.NewUserData(index)

	startTime := time.Now()

	defer func() {
		endTime := time.Now()

		interval := endTime.Sub(startTime)

		fmt.Println("interval: ", interval, int(interval))
		log.Println("log interval: ", interval, int(interval))
		tokens <- 1
	}()

	//c := dao.NewUserDataC()
	//defer c.Close()

	if userDataC == nil {
		userDataC = dao.NewUserDataC()
	}

	err := userDataC.Insert(data)
	if err != nil {
		fmt.Println("err: ", err)
	}

}

func InsertAllUserData(count int, concurrentCount int) {
	//fmt.Println("InsertAllUserData...", count, concurrentCount)
	log.Println("InsertAllUserData....")

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
			go InsertUserData(&waitGroup, tokens, i)

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

	log.Println("InsertAllUserData....end...")
}
