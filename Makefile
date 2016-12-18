.PHONY: all linux test inst cl clean

GIT_REV=$(shell git log -1 | base64)
NOW=$(shell date -u "+%Y/%m/%d %H:%M:%S")

all: bin/server bin/client-daemon bin/client-cli

linux: bin.linux/server bin/client-daemon bin/client-cli

test:
	go test .

inst:
	go get -u "github.com/go-gl/gl/v4.1-core/gl"
	go get -u "github.com/go-gl/glfw/v3.2/glfw"
	go get -u "github.com/Sirupsen/logrus"
	go get -u "github.com/golang/protobuf/protoc-gen-go"
	go get -u "google.golang.org/grpc"
	go get -u "github.com/chai2010/webp"
	go get -u "github.com/go-sql-driver/mysql"

proto/easel_service.pb.go: proto/easel_service.proto
	cd proto && PATH=$(GOPATH)/bin:$(PATH) protoc --go_out=plugins=grpc:. easel_service.proto

bin/server: proto/easel_service.pb.go $(shell find server image-filters proto util -type f -name '*.go')
	@mkdir -p bin
	go build -o bin/server \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/server"

bin/client-cli: proto/easel_service.pb.go $(shell find client-cli image-filters proto util -type f -name '*.go')
	@mkdir -p bin
	go build -o bin/client-cli \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/client-cli"


bin/client-daemon: proto/easel_service.pb.go $(shell find client-daemon image-filters proto util -type f -name '*.go')
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o bin.linux/client-daemon \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/client-daemon"


bin.linux/server: proto/easel_service.pb.go $(shell find server image-filters proto util -type f -name '*.go')
	@mkdir -p bin.linux
	GOOS=linux GOARCH=amd64 go build -o bin.linux/server \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/server"

bin.linux/client-cli: proto/easel_service.pb.go $(shell find client-cli image-filters proto util -type f -name '*.go')
	@mkdir -p bin.linux
	GOOS=linux GOARCH=amd64 go build -o bin.linux/client-cli \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/client-cli"


bin.linux/client-daemon: proto/easel_service.pb.go $(shell find client-daemon image-filters proto util -type f -name '*.go')
	@mkdir -p bin.linux
	go build -o bin/client-daemon \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/client-daemon"

clean:
	rm -rf bin bin.linux

cl:
	@find . -type f -name \*.go | xargs wc -l
	@echo $(shell git log --pretty=oneline | wc -l) commits.
