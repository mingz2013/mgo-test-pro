.PHONY: help
help:
	@echo '                                                                          '
	@echo 'Makefile help                                                             '
	@echo '                                                                          '
	@echo 'Usage:                                                                    '
	@echo '   make build                         构建所有二进制文件                    '
	@echo '                                                                          '
	@echo '   make docker-compose-build                        构建所有image       '
	@echo '   make docker-compose-up                           创建并启动所有容器       '
	@echo '   make docker-compose-stop                         停止容器               '
	@echo '   make docker-compose-start                        启动                   '
	@echo '   make docker-compose-down                         删除                   '
	@echo '   make docker-compose-ps                           查看状态               '
	@echo '   make docker-compose-logs                         查看日志              '
	@echo '                                                                          '
	@echo '   make docker                        构建所有image，用docker-compose-build替代     '
	@echo '                                                                          '
	@echo '   make run                           编译并启动容器，build, up             '
	@echo '                                                                          '
	@echo '                                                                          '
	@echo '                                                                          '


.PHONY: build
build:
	for d in build; do \
		echo $$d; \
		for f in $$d/*; do \
			echo $$f; \
			cd $$f; make build; cd ../../; \
		done; \
	done; \


.PHONY: docker-compose-build
docker-compose-build: build
	docker-compose build

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose up -d



.PHONY: docker-compose-start
docker-compose-start:
	docker-compose start


.PHONY: docker-compose-stop
docker-compose-stop:
	docker-compose stop

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose down


.PHONY: docker-compose-ps
docker-compose-ps:
	docker-compose ps


.PHONY: docker-compose-logs
docker-compose-logs:
	docker-compose logs


.PHONY: docker
docker:
	for d in build; do \
		echo $$d; \
		for f in $$d/*; do \
			echo $$f; \
			cd $$f; make docker; cd ../../; \
		done \
	done


.PHONY: run
run: docker-compose-build docker-compose-up

