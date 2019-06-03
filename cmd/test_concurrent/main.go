package main

import (
	"fmt"
	"github.com/mingz2013/mgo-test-pro/services"
	"os"
	"runtime"
	"strconv"
)

const defaultCount = 10
const defaultConcurrentCount = 1

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	maxUserId := 99999
	minUserId := 10000
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

	s = os.Getenv("MAX_USER_ID")
	if s == "" {
	} else {
		maxUserId, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("error MAX_USER_ID", err)
		}
	}

	s = os.Getenv("MIN_USER_ID")
	if s == "" {
	} else {
		minUserId, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("error MIN_USER_ID", err)
		}
	}

	services.TestConcurrent(count, concurrentCount, maxUserId, minUserId)
}
