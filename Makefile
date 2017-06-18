deps:
	cd $(shell pwd)/src/github.com/CANARIA/canaria-api && glide install

migrate:
	go get bitbucket.org/liamstask/goose/cmd/goose
	goose up

fmt:
	go fmt $(shell go list ./... | grep -v vendor)

run dev:
	GOPATH=$(shell pwd) go run app/server.go

build:
	GOOS=linux GOARCH=amd64 go build
