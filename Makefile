SHELL = /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)
SERVICE = editor-service

build: 
	@[ -d bin ] || mkdir bin
	( /bin/rm -f bin/* )
	( go build -o bin/editor-service src/main.go )

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/$(SERVICE) src/main.go

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get github.com/darrylwest/cassava-logger/logger
	go get github.com/darrylwest/go-unique/unique
	go get github.com/gorilla/websocket

format:
	( gofmt -s -w src/*.go src/edit/*.go test/unit/*.go )

lint:
	@( golint src/... && golint test/... )

test:
	@( go vet src/edit/*.go && go vet src/edit/*.go && go vet src/*.go && cd test/unit/ && go test -cover )
	@( make lint )

run:
	@make build && ./bin/$(SERVICE)

watch:
	go-watcher --loglevel=5

edit:
	vi -O3 src/*/*.go test/*/*.go src/*.go

.PHONY: format test watch examples
