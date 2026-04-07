.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

lint:fmt
	golint ./...
.PHONY:lint

vet:lint
	go vet ./...
.PHONY:vet

run:vet
	go run main.go
.PHONY:run

test:vet
	go test -v ./...
.PHONY:test
