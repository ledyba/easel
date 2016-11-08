.PHONY: run-test inst

GIT_REV=$(shell git log -1 | base64)
NOW=$(shell date -u "+%Y/%m/%d %H:%M:%S")

run-test: easel
	./easel

inst:
	go get -u "github.com/go-gl/gl/v4.1-core/gl"
	go get -u "github.com/go-gl/glfw/v3.2/glfw"
	go get -u "github.com/Sirupsen/logrus"

easel: $(shell find . -type f -name '*.go')
	go build -o easel \
					-ldflags "-X 'main.gitRev=$(GIT_REV)' -X 'main.buildAt=$(NOW)'" \
					"github.com/ledyba/easel/runner"
