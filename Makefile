.PHONY: all linux test inst cl clean

GIT_REV=$(shell git log -1 | base64 | tr -d '[:space:]')
NOW=$(shell date -u "+%Y/%m/%d %H:%M:%S")

all: \
	.bin/easel-client-daemon \
	.bin/easel-client-cli \
	.bin/easel-server

linux: \
	.bin/easel-client-daemon \
	.bin/easel-client-cli \
	.bin.linux/easel-server

test:
	go test "github.com/ledyba/easel/..."

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

####### Executables #######

.bin/easel-server: proto/easel_service.pb.go $(shell find easel-server image-filters proto util -type f -name '*.go')
	@mkdir -p .bin
	go build -o .bin/easel-server \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/easel-server"

.bin/easel-client-cli: proto/easel_service.pb.go $(shell find easel-client-cli image-filters proto util -type f -name '*.go')
	@mkdir -p .bin
	go build -o .bin/easel-client-cli \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/easel-client-cli"


.bin/easel-client-daemon: proto/easel_service.pb.go $(shell find easel-client-daemon image-filters proto util -type f -name '*.go')
	@mkdir -p .bin
	go build -o .bin/easel-client-daemon \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/easel-client-daemon"

####### Linux Targets #######

.bin.linux/easel-server: proto/easel_service.pb.go $(shell find easel-server image-filters proto util -type f -name '*.go')
	@mkdir -p .bin.linux
	GOOS=linux GOARCH=amd64 go build -o .bin.linux/easel-server \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/easel-server"

.bin.linux/easel-client-cli: proto/easel_service.pb.go $(shell find easel-client-cli image-filters proto util -type f -name '*.go')
	@mkdir -p .bin.linux
	GOOS=linux GOARCH=amd64 go build -o .bin.linux/easel-client-cli \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/easel-client-cli"


.bin.linux/easel-client-daemon: proto/easel_service.pb.go $(shell find easel-client-daemon image-filters proto util -type f -name '*.go')
	@mkdir -p .bin.linux
	GOOS=linux GOARCH=amd64 go build -o .bin.linux/easel-client-daemon \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/easel-client-daemon"

####### Misc #######

clean:
	rm -rf .bin .bin.linux

cl:
	@find . -type f -name \*.go | xargs wc -l
	@echo $(shell git log --pretty=oneline | wc -l) commits.
