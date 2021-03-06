package services

import (
	"fmt"
	"mgo-test-pro/dao"
	"log"
	"sync"
	"time"
)

var userDataC *dao.UserDataC

//func init() {
//	//userDataC = dao.NewUserDataC()
//}

func InsertUserData(waitGroup *sync.WaitGroup, tokens chan<- int, index, collectionIndex int) {
	//log.Println("InsertUserData..")

	defer waitGroup.Done()

	data := dao.NewUserData(index)

	//c := dao.NewUserDataC()
	//defer c.Close()

	//if userDataC == nil {
	userDataC = dao.NewUserDataC(collectionIndex)
	//}

	startTime := time.Now()

	defer func() {
		endTime := time.Now()

		interval := endTime.Sub(startTime)

		a := []interface{}{
			"insert_data", "|", index, "|", collectionIndex, "|", interval, "|", int(interval),
		}

		fmt.Println(a...)
		log.Println(a...)

		tokens <- 1
	}()

	err := userDataC.Insert(data)
	if err != nil {
		fmt.Println("err: ", err)
	}

}

func getCollectionIndex(i, count, collectionCount int) (index int) {

	sub := count / collectionCount

	return i / sub

}

func InsertAllUserData(count int, concurrentCount, collectionCount int) {
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
			collectionIndex := getCollectionIndex(i, count, collectionCount)
			log.Println("collectionIndex", collectionIndex)
			go InsertUserData(&waitGroup, tokens, i, collectionIndex)

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
