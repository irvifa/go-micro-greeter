GOPATH:=$(shell go env GOPATH)
GIT_COMMIT:=$(shell git rev-parse --short HEAD)
APP_NAME:=go-micro-greeter

include telepresence.mk

.PHONY: clean proto test docker

clean:
	rm $(APP_NAME)

proto:
	protoc --go_out=.  --micro_out=. proto/saying/saying.proto

build: proto	
	CGO_ENABLED=0 go build -a -ldflags "-X github.com/irvifa/go-micro-greeter/version.GitCommit=$(GIT_COMMIT)" -o $(APP_NAME) srv/main.go
