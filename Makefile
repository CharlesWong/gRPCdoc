.PHONY: build proto test

default: build

proto:
	protoc -I=./model/ --go_out=plugins=grpc:./model/ ./model/*.proto

build: proto 
	GOARCH=amd64 GOOS=linux go build -v -o bin/grpcdoc-amd64-linux grpcdoc.go
	GOARCH=amd64 GOOS=darwin go build -v -o bin/grpcdoc-amd64-darwin grpcdoc.go

test:
	go test ./...
