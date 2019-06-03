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
	@echo '   START_USER_ID             起始的userid， 整数，  默认值：10000                                         '
	@echo '   INSERT_COUNT              插入的数据条数，整数，  默认值：10                                         '
	@echo '   CONCURRENT_COUNT          并发数，       整数，  默认值：1                                         '
	@echo '   DB_HOST                   数据库地址，    字符串，默认值：127.0.0.1:27017                                   '
	@echo '   DB_NAME                   数据库名称，    字符串，默认值：idlethree                                         '
	@echo '   USER_DATA_COLLECTION_NAME 集合名称，     字符串， 默认值：user_data                                         '
	@echo '                                                                          '



.PHONY: build
build:
	@echo ${GOPATH}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o insert_data ${GOPATH}/src/github.com/mingz2013/mgo-test-pro/cmd/insert_data/main.go

.PHONY: build-macos
build-macos:
	go build -o insert_data ${GOPATH}/src/github.com/mingz2013/mgo-test-pro/cmd/insert_data/main.go


.PHONY: docker
docker: build
	docker build . -t insert_data:latest


.PHONY: run
run:
	time START_USER_ID=10000 INSERT_COUNT=3000000 CONCURRENT_COUNT=4 DB_HOST=127.0.0.1:27017 DB_NAME=idlethree  USER_DATA_COLLECTION_NAME=user_data ./insert_data > 1.log


.PHONY: run-docker
run-docker: docker
	docker run insert_data -e START_USER_ID=30000 INSERT_COUNT=30000 CONCURRENT_COUNT=1 DB_HOST=127.0.0.1:27017 DB_NAME=idlethree  USER_DATA_COLLECTION_NAME=user_data
