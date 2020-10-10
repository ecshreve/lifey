mod-init:
	go mod init github.com/ecshreve/lifey

mod-tidy:
	go mod tidy
	
build:
	go build -o bin/lifey github.com/ecshreve/lifey/cmd/lifey

install:
	go install -i github.com/ecshreve/lifey/cmd/lifey

run-only:
	bin/lifey

run: build run-only

test:
	go test github.com/ecshreve/lifey/...

testv:
	go test -v github.com/ecshreve/lifey/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/lifey/...