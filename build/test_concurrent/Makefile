GOPATH:=$(shell go env GOPATH)
PRO_PATH:=/Users/zhaojm/company/dfprojects/idlethree/trunk/tools/mgo_test
GOPATH:=${GOPATH}:${PRO_PATH}
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o test_concurrent ${PRO_PATH}/src/mgo_test/cmd/test_concurrent/main.go

.PHONY: docker
docker:
	docker build . -t test_concurrent:latest