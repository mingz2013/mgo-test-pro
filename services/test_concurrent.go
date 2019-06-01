package services

//
//import (
//	"fmt"
//	"github.com/mingz2013/mgo-test-pro/dao"
//	"time"
//)
//
//func FindUserData(tokens chan<- int, userId int) {
//
//	c := dao.UserDataC{}
//
//	startTime := time.Now()
//
//	defer func() {
//		endTime := time.Now()
//
//		interval := endTime.Sub(startTime)
//
//		fmt.Println("interval: ", interval)
//
//		tokens <- 1
//	}()
//
//	data, err := c.FindByUserId(userId)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(data)
//
//}
//
//func TestConcurrent(count int, concurrentCount int) {
//
//	tokens := make(chan int, concurrentCount+10)
//
//	for i := 0; i < concurrentCount; i++ {
//		tokens <- 1
//	}
//
//	for i := 0; i < count; i++ {
//		select {
//		case <-tokens:
//			// 拿到一个token，
//			go FindUserData(tokens, i)
//			break
//
//		case <-time.After(time.Microsecond * 1):
//			break
//		}
//	}
//
//}
