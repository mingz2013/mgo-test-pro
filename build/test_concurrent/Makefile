GOPATH:=$(shell go env GOPATH)
# PRO_PATH:=/Users/zhaojm/company/dfprojects/idlethree/trunk/tools/mgo_test
# GOPATH:=${GOPATH}:${PRO_PATH}



.PHONY: help
help:
	@echo '                                                                          '
	@echo 'Makefile help                                                             '
	@echo '                                                                          '
	@echo 'Usage:                                                                    '
	@echo '   make build                         构建所有二进制文件                    '
	@echo '   make build-macos                   构建macos上的二进制                                                    '
	@echo '   make docker                         构建docker镜像             '
	@echo '   make run                            执行二进制                                           '
	@echo '   make run-docker                     执行docker容器                                                  '
	@echo '                                                                        '
	@echo 'ENV:                                                                  '
	@echo '   MIN_USER_ID             最小的userid， 整数，  默认值：99999                                         '
	@echo '   MAX_USER_ID             最大的userid， 整数，  默认值：10000                                         '
	@echo '   FIND_COUNT                查询的数据次数，整数，  默认值：10                                         '
	@echo '   CONCURRENT_COUNT          并发数，       整数，  默认值：1                                         '
	@echo '   DB_HOST                   数据库地址，    字符串，默认值：127.0.0.1:27017                                   '
	@echo '   DB_NAME                   数据库名称，    字符串，默认值：idlethree                                         '
	@echo '   USER_DATA_COLLECTION_NAME 集合名称，     字符串， 默认值：user_data                                         '
	@echo '                                                                          '



.PHONY: build
build:
	@echo ${GOPATH}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o test_concurrent ${GOPATH}/src/github.com/mingz2013/mgo-test-pro/cmd/test_concurrent/main.go

.PHONY: build-macos
build-macos:
	go build -o test_concurrent ${GOPATH}/src/github.com/mingz2013/mgo-test-pro/cmd/test_concurrent/main.go


.PHONY: docker
docker: build
	docker build . -t test_concurrent:latest


.PHONY: run
run:
	time MIN_USER_ID=10000 MAX_USER_ID=99999 FIND_COUNT=1000 CONCURRENT_COUNT=4 DB_HOST=127.0.0.1:27017 DB_NAME=idlethree  USER_DATA_COLLECTION_NAME=user_data ./test_concurrent > 1.log


.PHONY: run-docker
run-docker: docker
	docker run test_concurrent -e MIN_USER_ID=10000 MAX_USER_ID=99999 FIND_COUNT=1000 CONCURRENT_COUNT=1 DB_HOST=127.0.0.1:27017 DB_NAME=idlethree  USER_DATA_COLLECTION_NAME=user_data
