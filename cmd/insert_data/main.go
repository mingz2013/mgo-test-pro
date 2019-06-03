package main

import (
	"fmt"
	"github.com/mingz2013/mgo-test-pro/services"
	"log"
	"os"
	"strconv"
)

const defaultCount = 10
const defaultConcurrentCount = 1

func main() {

	fmt.Println("main...")
	log.Println("main....")

	var err error
	count := defaultCount
	s := os.Getenv("INSERT_COUNT")
	if s == "" {
		count = defaultCount
	} else {
		count, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("error INSERT_COUNT", err)
		}
	}

	concurrentCount := defaultConcurrentCount
	s = os.Getenv("CONCURRENT_COUNT")
	if s == "" {
		concurrentCount = defaultConcurrentCount
	} else {
		concurrentCount, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("error CONCURRENT_COUNT", err)
		}
	}

	log.Println(count, concurrentCount)

	//
	services.InsertAllUserData(count, concurrentCount)
}
