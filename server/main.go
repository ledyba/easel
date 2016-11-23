package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"

	"github.com/ledyba/easel/proto"
	"google.golang.org/grpc"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	log "github.com/Sirupsen/logrus"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

var port = flag.Int("port", 14514, "port to listen")

func main() {
	var err error
	err = glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()
	log.Debug("Initialized.")

	printStartupBanner()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	proto.RegisterEaselServiceServer(server, newServer())
	log.Infof("Now listen at :%d", *port)
	server.Serve(lis)

}
