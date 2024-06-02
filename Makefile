.PHONY: build test run clean all

build:
	@go build -o todo-app cmd/main.go

test:
	@go test -v ./...

run:
	@go run cmd/main.go run

clean:
	@rm -f ./todo-app

all: build run clean
