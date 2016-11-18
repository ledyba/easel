.PHONY: test run-test inst proto cl

GIT_REV=$(shell git log -1 | base64)
NOW=$(shell date -u "+%Y/%m/%d %H:%M:%S")

run-test: easel-server
	./easel-server

test:
	go test .

inst:
	go get -u "github.com/go-gl/gl/v4.1-core/gl"
	go get -u "github.com/go-gl/glfw/v3.2/glfw"
	go get -u "github.com/Sirupsen/logrus"
	go get -u "github.com/golang/protobuf/protoc-gen-go"

proto:
	cd server/proto && PATH=$(GOPATH)/bin:$(PATH) protoc --go_out=plugins=grpc:. easel_service.proto

easel-server: $(shell find . -type f -name '*.go')
	go build -o easel-server \
					-ldflags "-v -X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/server"

cl:
	find . -type f -name \*.go | xargs wc -l
