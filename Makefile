.PHONY: run-test inst

run-test: easel
	./easel

inst:
	go get -u "github.com/go-gl/gl/v4.1-core/gl"
	go get -u "github.com/go-gl/glfw/v3.2/glfw"
	go get -u "github.com/Sirupsen/logrus"

easel: $(shell find . -type f -name '*.go')
	go build -o easel "github.com/ledyba/easel/runner"
