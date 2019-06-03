package main

import (
	"github.com/mingz2013/mgo-test-pro/services"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	services.TestConcurrent(10, 10)
}
