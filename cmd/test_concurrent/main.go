package main

import "github.com/mingz2013/mgo_test/services"

func main() {
	services.TestConcurrent(10, 10)
}
