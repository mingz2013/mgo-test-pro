package main

import "github.com/mingz2013/mgo-test-pro/services"

func main() {
	services.TestConcurrent(10, 10)
}
