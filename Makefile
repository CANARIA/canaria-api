deps:
	glide install

migrate:
	go get bitbucket.org/liamstask/goose/cmd/goose
	goose up

fmt:
	go fmt $(shell go list ./... | grep -v vendor)

dev run:
	docker-compose restart api && docker-compose logs -ft api
