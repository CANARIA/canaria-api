__GITHUB_ORG_REPO=github.com/CANARIA/canaria-api

deps:
	glide install

migrate:
	go get bitbucket.org/liamstask/goose/cmd/goose
	cd $(shell pwd)/src/$(__GITHUB_ORG_REPO) && goose up

fmt:
	go fmt $(shell go list ./... | grep -v vendor)

run dev:
	go run server.go

init gae:
	GOPATH=$(shell pwd)/gae/gopath go run gae/app/init.go

symbol: core-symbol core-vendor

core-symbol:
	cd $(shell pwd)/gae/gopath/src/$(__GITHUB_ORG_REPO) && ln -s $(shell pwd)/core core

core-vendor:
	cd $(shell pwd)/gae/gopath && ln -s $(shell pwd)/vendor vendor

build:
	GOOS=linux GOARCH=amd64 go build
