__GITHUB_ORG_REPO=github.com/CANARIA/canaria-api

deps:
	glide install

migrate:
	go get bitbucket.org/liamstask/goose/cmd/goose
	cd core && goose up

fmt:
	go fmt $(shell go list ./... | grep -v vendor)

run dev:
	goapp serve app/local.yaml

build:
	GOOS=linux GOARCH=amd64 go build
