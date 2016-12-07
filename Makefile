.PHONY: all test inst cl clean

GIT_REV=$(shell git log -1 | base64)
NOW=$(shell date -u "+%Y/%m/%d %H:%M:%S")

all: bin/server bin/client bin/cli

test:
	go test .

inst:
	go get -u "github.com/go-gl/gl/v4.1-core/gl"
	go get -u "github.com/go-gl/glfw/v3.2/glfw"
	go get -u "github.com/Sirupsen/logrus"
	go get -u "github.com/golang/protobuf/protoc-gen-go"
	go get -u "google.golang.org/grpc"
	go get -u "github.com/chai2010/webp"

proto/easel_service.pb.go: proto/easel_service.proto
	cd proto && PATH=$(GOPATH)/bin:$(PATH) protoc --go_out=plugins=grpc:. easel_service.proto

bin/server: proto/easel_service.pb.go $(shell find server-daemon filters proto util -type f -name '*.go')
	@mkdir -p bin
	go build -o bin/server \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/server-daemon"

bin/cli: proto/easel_service.pb.go $(shell find client-cli filters proto util -type f -name '*.go')
	@mkdir -p bin
	go build -o bin/cli \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/client-cli"


bin/client: proto/easel_service.pb.go $(shell find client-daemon filters proto util -type f -name '*.go')
	@mkdir -p bin
	go build -o bin/client \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/client-daemon"

clean:
	rm -rf bin

cl:
	@find . -type f -name \*.go | xargs wc -l
	@echo $(shell git log --pretty=oneline | wc -l) commits.
