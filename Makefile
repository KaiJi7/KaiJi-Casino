.PHONY: build all

build:
	go build

build-linux:
	GOOS=linux GOARCH=amd64 go build
