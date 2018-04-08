SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( /bin/rm -f bin/* )
	( go build -o bin/editor-service src/main.go )

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/editor-service src/main.go

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/oklog/ulid
	go get github.com/franela/goblin

format:
	( gofmt -s -w src/*.go src/edit/*.go test/unit/*.go )

lint:
	@( golint src/... && golint test/... )

test:
	@( go vet src/edit/*.go && go vet src/edit/*.go && go vet src/*.go && cd test/unit/ && go test -cover )
	@( make lint )

run:
	go run src/main.go

watch:
	./watcher.js

edit:
	vi -O3 src/*/*.go test/*.go src/*.go

open:
	(cd public && python -m SimpleHTTPServer)

.PHONY: format test watch examples
