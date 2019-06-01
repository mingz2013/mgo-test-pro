package services

import (
	"fmt"
	"github.com/mingz2013/mgo-test-pro/dao"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func NewUserData(index int) (data *dao.UserData) {

	data = &dao.UserData{}

	startUserId := 10000
	var err error
	s := os.Getenv("START_USER_ID")
	if s != "" {
		startUserId, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("error START_USER_ID", err)
		}
	}

	//data.UserId = string(datastore.GetKeyId())
	data.UserId = strconv.Itoa(index + startUserId)
	data.Achievement = ""
	data.ActiveTimestamp = ""
	data.EnterGameAt = ""
	data.Exp = 10
	data.FightPower = 10
	data.Level = 10
	data.McDailyRewardAtDay = ""
	data.McExpireAt = ""
	data.WcExpireNotice = false
	data.FightPower = 1
	data.Merit = ""
	data.MilitaryRank = ""
	data.Newbie = map[string]interface{}{
		"aaa": "aaa",
	}
	data.OfflineAt = ""
	data.OnlineAt = ""
	data.Package = map[string]interface{}{
		"aaa": "aaa",
	}

	data.ServerId = 1
	data.SToken = "aaaaaaaaaa"
	data.VipLevel = 1
	data.VipExp = 1

	return
}

var userDataC *dao.UserDataC

//func init() {
//	//userDataC = dao.NewUserDataC()
//}

func InsertUserData(waitGroup *sync.WaitGroup, tokens chan<- int, index int) {
	log.Println("InsertUserData..")

	defer waitGroup.Done()

	data := NewUserData(index)

	startTime := time.Now()

	defer func() {
		endTime := time.Now()

		interval := endTime.Sub(startTime)

		fmt.Println("interval: ", interval)

		tokens <- 1
	}()

	//c := dao.NewUserDataC()
	//defer c.Close()

	if userDataC == nil {
		userDataC = dao.NewUserDataC()
	}

	err := userDataC.Insert(data)
	if err != nil {
		fmt.Println(err)
	}

}

func InsertAllUserData(count int, concurrentCount int) {
	//fmt.Println("InsertAllUserData...", count, concurrentCount)
	//log.Println("InsertAllUserData....")

	tokens := make(chan int, concurrentCount+10)

	for i := 0; i < concurrentCount; i++ {
		tokens <- 1
	}

	//fmt.Println("tokens: ", tokens)

	var waitGroup sync.WaitGroup

	var i int

	i = 0

	for {
		//log.Println("for...count...")
		select {
		case <-tokens:
			// 拿到一个token，
			waitGroup.Add(count)
			go InsertUserData(&waitGroup, tokens, i)

			i++

			if i >= count {
				break
			}

			continue

		case <-time.After(time.Microsecond * 1):
			continue

		}
	}

	//w:= make(chan int)

	//<-w
	//for i range <-tokens{
	//
	//}

	waitGroup.Wait()

}
