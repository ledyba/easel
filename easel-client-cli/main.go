package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
	filters "github.com/ledyba/easel/image-filters"
	"github.com/ledyba/easel/proto"
	"github.com/ledyba/easel/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

/* Server to work with */
var server *string = flag.String("server", "localhost:3000", "server to connect")

/* Filter Flags */
var filter *string = flag.String("filter", "lanczos", "applied filter name.")
var lobes *int = flag.Int("filter_lobes", 10, "lobes parameter")
var scale *float64 = flag.Float64("scale", 2.0, "scale")
var quality *float64 = flag.Float64("quality", 95.0, "quality")

/* General */
var help *bool = flag.Bool("help", false, "Print help and exit")

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
	var input []byte
	var src image.Image
	switch *filter {
	case "lanczos":
		err = filters.UpdateLanczos(serv, presp.EaselId, presp.PaletteId, *lobes)
		if err != nil {
			log.Fatal(err)
		}
		input, src, err = util.LoadImage(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		widthf := *scale * float64(src.Bounds().Dx())
		heightf := *scale * float64(src.Bounds().Dy())
		output, err = filters.RenderLanczos(serv, presp.EaselId, presp.PaletteId, input, src, int(widthf), int(heightf), float32(*quality))
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
