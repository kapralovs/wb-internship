.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run-main
run-main:
	go run ./cmd/app

.PHONY: run-nats
run-nats:
	go run ./cmd/nats

.DEFAULT_GOAL := build