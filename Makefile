SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)
TARGET=/usr/local/bin

build: 
	@[ -d bin ] || mkdir bin
	( /bin/rm -f bin/* )
	( go build -o bin/unique src/main.go )
	( go build -o bin/unique-tcp src/unique-tcp.go )

install:
	@make build
	cp -f bin/unique $(TARGET)/unique
	ln -f $(TARGET)/unique $(TARGET)/ulid
	ln -f $(TARGET)/unique $(TARGET)/uuid
	ln -f $(TARGET)/unique $(TARGET)/guid
	ln -f $(TARGET)/unique $(TARGET)/tsid
	ln -f $(TARGET)/unique $(TARGET)/txid
	ln -f $(TARGET)/unique $(TARGET)/cuid
	ln -f $(TARGET)/unique $(TARGET)/xuid

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linux/unique src/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linux/unique-tcp src/unique-tcp.go

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linux/unique-tcp src/unique-tcp.go
	( cd linux && ./build.sh )

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/oklog/ulid
	go get github.com/franela/goblin

format:
	( gofmt -s -w src/*.go src/unique/*.go test/*.go )

lint:
	@( golint src/... && golint test/... )

test:
	@( go vet src/unique/*.go && go vet src/unique/*.go && go vet src/*.go && cd test/ && go test -cover )
	@( make lint )

qtest:
	@( cd test && go test -cover )

run:
	go run src/main.go

run-tcp:
	go run src/unique-tcp.go

examples:
	javac examples/UniqueClient.java

watch:
	./watcher.js

edit:
	vi -O3 src/*/*.go test/*.go src/*.go

.PHONY: format test watch examples
