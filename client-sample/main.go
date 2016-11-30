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
	eresp, err := serv.NewEasel(context.Background(), &proto.NewEaselRequest{
		EaselId: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Easel Created: %s", eresp.EaselId)
	defer func() {
		serv.DeleteEasel(context.Background(), &proto.DeleteEaselRequest{
			EaselId: eresp.EaselId,
		})
		log.Printf("Easel Deleted: %s", eresp.EaselId)
	}()
	presp, err := serv.NewPalette(context.Background(), &proto.NewPaletteRequest{
		EaselId: eresp.EaselId,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Palette Created: (%s > %s)", presp.EaselId, presp.PaletteId)
	defer func() {
		serv.DeletePalette(context.Background(), &proto.DeletePaletteRequest{
			EaselId:   eresp.EaselId,
			PaletteId: presp.PaletteId,
		})
		log.Printf("Palette Deleted: (%s > %s)", presp.EaselId, presp.PaletteId)
	}()

	_, err = serv.UpdatePalette(context.Background(), &proto.UpdatePaletteRequest{
		EaselId:   presp.EaselId,
		PaletteId: presp.PaletteId,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Palette Updated: (%s > %s)", presp.EaselId, presp.PaletteId)
}
