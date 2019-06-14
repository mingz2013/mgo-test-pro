package main

import (
	"fmt"
	"github.com/mingz2013/mgo-test-pro/services"
	"log"
	"os"
	"runtime"
	"strconv"
)

const defaultCount = 10
const defaultConcurrentCount = 1
const defaultCollectionCount = 1

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

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

	collectionCount := defaultCollectionCount
	s = os.Getenv("COLLECTION_COUNT")
	if s == "" {
		collectionCount = defaultCollectionCount
	} else {
		collectionCount, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("err COLLECTION_COUNT", err)
		}
	}

	//
	services.InsertAllUserData(count, concurrentCount, collectionCount)
}
