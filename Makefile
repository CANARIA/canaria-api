deps:
	go get bitbucket.org/liamstask/goose/cmd/goose
	glide install

migrate:
	goose up

fmt:
	go fmt $(shell go list ./... | grep -v vendor)
