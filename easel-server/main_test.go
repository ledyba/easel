package main

import (
	"context"
	"log"
	"net"
	"runtime"
	"sync"
	"testing"

	"google.golang.org/grpc"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/ledyba/easel/easel-server/impl"
	"github.com/ledyba/easel/proto"
)

func startup(t *testing.T) {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		t.Fatal("Failed to init glfw", t)
	}
}

func shutdown() {
	glfw.Terminate()
	runtime.UnlockOSThread()
}

func TestServerWithRPC(t *testing.T) {
	em := impl.NewEaselMaker()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		startup(t)
		defer shutdown()
		em.Start()
	}()

	lis, err := net.Listen("tcp", ":8192")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := impl.NewServer(em)
	gserver := grpc.NewServer()
	go server.StartGC()
	go (func() {
		defer wg.Done()
		proto.RegisterEaselServiceServer(gserver, server)
		gserver.Serve(lis)
	})()

	conn, err := grpc.Dial(":8192", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := proto.NewEaselServiceClient(conn)
	_, err = cli.DeleteEasel(context.Background(), &proto.DeleteEaselRequest{EaselId: "po"})
	if err != impl.ErrEaselNotFound {
		log.Fatal(err, "!=", impl.ErrEaselNotFound)
	}

	em.Stop()
	gserver.Stop()
	wg.Wait()
}
