package main

import (
	"flag"
	"net"
	"runtime"

	"github.com/ledyba/easel/proto"
	impl "github.com/ledyba/easel/server-impl"
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

/* Serving */
var listen = flag.String("listen", ":3000", "listen addr")

/* General */
var help *bool = flag.Bool("help", false, "Print help and exit")

func startServer(lis net.Listener, em *impl.EaselMaker) {
	log.Infof("Now listen at %s", *listen)
	gserver := grpc.NewServer()
	server := impl.NewServer(em)
	go server.StartGC()
	proto.RegisterEaselServiceServer(gserver, server)
	gserver.Serve(lis)
}

func main() {
	printLogo()
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	var err error
	err = glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()
	log.Debug("Initialized.")

	printStartupBanner()

	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	em := impl.NewEaselMaker()
	go startServer(lis, em)
	em.Start()
}
