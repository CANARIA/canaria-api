__GITHUB_ORG_REPO=github.com/CANARIA/canaria-api

deps:
	glide install

migrate:
	go get bitbucket.org/liamstask/goose/cmd/goose
	cd $(shell pwd)/src/$(__GITHUB_ORG_REPO) && goose up

fmt:
	go fmt $(shell go list ./... | grep -v vendor)

run dev:
	goapp serve app/local.yaml

build:
	GOOS=linux GOARCH=amd64 go build
