.DEFAULT_GOAL := docker-run

clean:
	rm QuestionBankAPI

setup:
	go get github.com/golang/dep/cmd/dep
#	source ./.env

build: setup
	dep ensure
	env GOOS=linux GOARCH=amd64 go build

docker-clean:
	docker stop $$(docker ps -a -q) || true
	docker rm $$(docker ps -a -q) || true
	docker image rm questionbankapi_api || true

compose: docker-clean build
	docker-compose up -d
