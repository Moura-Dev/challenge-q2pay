PKGS ?= $(shell go list ./...)
.PHONY: all services clean

ensure-dependencies:
	go mod tidy
down:
	docker-compose down && docker volume prune -f

mod:
	go mod tidy

run:
	go run main.go
services:
	docker-compose up -d
