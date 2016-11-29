package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/ledyba/easel/proto"
	"golang.org/x/net/context"
)

import "google.golang.org/grpc"

var server *string = flag.String("server", "localhost:3000", "server to connect")

func main() {
	flag.Parse()
	log.Printf("*** Easel Client Example ***")
	var err error
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	serv := proto.NewEaselServiceClient(conn)
	resp, err := serv.NewEasel(context.Background(), &proto.NewEaselRequest{
		EaselId: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Easel Created: %s", resp.EaselId)
	defer func() {
		serv.DeleteEasel(context.Background(), &proto.DeleteEaselRequest{
			EaselId: resp.EaselId,
		})
		log.Printf("Easel Deleted: %s", resp.EaselId)
	}()
}
