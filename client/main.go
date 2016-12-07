package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ledyba/easel/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var server *string = flag.String("server", "localhost:3000", "server to connect")
var filter *string = flag.String("filter", "lanczos", "Filternames.")
var lobes *int = flag.Int("lobes", 10, "lobes parameter")
var help *bool = flag.Bool("help", false, "Print help and exit")
var scale *float64 = flag.Float64("scale", 2.0, "scale")

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage of %s:
	%s [OPTIONS] IN OUT
Options:
`, os.Args[0], os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	args := flag.Args()
	printStartupBanner()
	if len(args) <= 0 || *help {
		usage()
		return
	}
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

	/**** Create Easel ****/
	log.Printf("Easel Created: %s", eresp.EaselId)
	defer func() {
		serv.DeleteEasel(context.Background(), &proto.DeleteEaselRequest{
			EaselId: eresp.EaselId,
		})
		log.Printf("Easel Deleted: %s", eresp.EaselId)
	}()

	/**** Create Palette ****/
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

	/**** Update Palette ****/
	var output []byte
	switch *filter {
	case "lanczos":
		err = UpdateLanczos(serv, presp.EaselId, presp.PaletteId, *lobes)
		if err != nil {
			log.Fatal(err)
		}
		output, err = RenderLanczos(serv, presp.EaselId, presp.PaletteId, flag.Arg(0), *scale)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Unknown filter: %s", *filter)
	}

	/**** Render Image ****/
	fname := flag.Arg(0)
	err = ioutil.WriteFile(fmt.Sprintf("%s.out.png", strings.TrimSuffix(fname, path.Ext(fname))), output, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Rendered: (%s > %s) %d bytes", presp.EaselId, presp.PaletteId, len(output))

}
