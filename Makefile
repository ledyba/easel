.PHONY: all linux test inst cl clean

all: \
	.bin/easel-client-daemon \
	.bin/easel-client-cli \
	.bin/easel-server

linux: \
	.bin/easel-client-daemon \
	.bin/easel-client-cli \
	.bin.linux/easel-server

test:
	go test -cover "github.com/ledyba/easel/..."

inst:
	go get -u "github.com/go-gl/gl/v4.1-core/gl"
	go get -u "github.com/go-gl/glfw/v3.2/glfw"
	go get -u "github.com/Sirupsen/logrus"
	go get -u "github.com/golang/protobuf/protoc-gen-go"
	go get -u "google.golang.org/grpc"
	go get -u "github.com/chai2010/webp"
	go get -u "github.com/go-sql-driver/mysql"
	go get -u "golang.org/x/image/tiff"

proto/easel_service.pb.go: proto/easel_service.proto
	go generate github.com/ledyba/easel/proto

####### Executables #######

.bin/easel-server: proto/easel_service.pb.go $(shell find easel-server image-filters proto util -type f -name '*.go')
	@mkdir -p .bin
	go generate github.com/ledyba/easel/easel-server
	go build -o .bin/easel-server "github.com/ledyba/easel/easel-server"

.bin/easel-client-cli: proto/easel_service.pb.go $(shell find easel-client-cli image-filters proto util -type f -name '*.go')
	@mkdir -p .bin
	go generate github.com/ledyba/easel/easel-client-cli
	go build -o .bin/easel-client-cli "github.com/ledyba/easel/easel-client-cli"


.bin/easel-client-daemon: proto/easel_service.pb.go $(shell find easel-client-daemon image-filters proto util -type f -name '*.go')
	@mkdir -p .bin
	go generate github.com/ledyba/easel/easel-client-daemon
	go build -o .bin/easel-client-daemon "github.com/ledyba/easel/easel-client-daemon"

####### Linux Targets #######

.bin.linux/easel-server: proto/easel_service.pb.go $(shell find easel-server image-filters proto util -type f -name '*.go')
	@mkdir -p .bin.linux
	go generate github.com/ledyba/easel/easel-server
	GOOS=linux GOARCH=amd64 go build -o .bin.linux/easel-server "github.com/ledyba/easel/easel-server"

.bin.linux/easel-client-cli: proto/easel_service.pb.go $(shell find easel-client-cli image-filters proto util -type f -name '*.go')
	@mkdir -p .bin.linux
	go generate github.com/ledyba/easel/easel-client-cli
	GOOS=linux GOARCH=amd64 go build -o .bin.linux/easel-client-cli "github.com/ledyba/easel/easel-client-cli"


.bin.linux/easel-client-daemon: proto/easel_service.pb.go $(shell find easel-client-daemon image-filters proto util -type f -name '*.go')
	@mkdir -p .bin.linux
	go generate github.com/ledyba/easel/easel-client-daemon
	GOOS=linux GOARCH=amd64 go build -o .bin.linux/easel-client-daemon "github.com/ledyba/easel/easel-client-daemon"

####### Misc #######

clean:
	rm -rf .bin .bin.linux **/*.gen.go proto/*.pb.go ./.DS_Store **/.DS_Store

cl:
	@find . -type f -name \*.go | xargs wc -l
	@echo $(shell git log --pretty=oneline | wc -l) commits.
