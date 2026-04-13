.BUILD_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

goimports: fmt
	goimports -local=cli-tool -w .
.PHONY:goimports

build:goimports
	go build -o cli-tool cmd/main.go
.PHONY:build

run:build
	./cli-tool
.PHONY:run

test:vet
	go test -v ./...
.PHONY:test
