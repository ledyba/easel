.PHONY: test run-test inst proto cl

GIT_REV=$(shell git log -1 | base64)
NOW=$(shell date -u "+%Y/%m/%d %H:%M:%S")

run-test: easel-server
	./easel-server

test:
	go test .

test-rpc: easel-server easel-client
	./easel-server
	./easel-client


inst:
	go get -u "github.com/go-gl/gl/v4.1-core/gl"
	go get -u "github.com/go-gl/glfw/v3.2/glfw"
	go get -u "github.com/Sirupsen/logrus"
	go get -u "github.com/golang/protobuf/protoc-gen-go"
	go get -u "google.golang.org/grpc"
	go get -u "github.com/chai2010/webp"

proto:
	cd proto && PATH=$(GOPATH)/bin:$(PATH) protoc --go_out=plugins=grpc:. easel_service.proto

easel-server: $(shell find . -type f -name '*.go')
	go build -o easel-server \
					-ldflags "-v -X 'cli.gitRev=$(GIT_REV)' -X 'cli.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/server"

easel-client: $(shell find . -type f -name '*.go')
	go build -o easel-client \
					-ldflags "-v -X 'cli.gitRev=$(GIT_REV)' -X 'cli.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/client"

cl:
	@find . -type f -name \*.go | xargs wc -l
	@echo $(shell git log --pretty=oneline | wc -l) commits.
