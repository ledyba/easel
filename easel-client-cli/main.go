package main

import (
	"crypto/tls"
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
	"google.golang.org/grpc/credentials"
)

/* Server to work with */
var server = flag.String("server", "localhost:3000", "server to connect")
var cert = flag.String("cert", "", "cert file")
var certKey = flag.String("cert_key", "", "private key file")

/* Filter Flags */
var filter = flag.String("filter", "lanczos", "applied filter name.")
var lobes = flag.Int("filter_lobes", 10, "lobes parameter")
var scale = flag.Float64("scale", 2.0, "scale")
var quality = flag.Float64("quality", 95.0, "quality")
var mimeType = flag.String("mime_type", "image/png", "output format. One of: ['image/png', 'image/jpg', 'image/webp']")

/* General */
var help = flag.Bool("help", false, "Print help and exit")
var ping = flag.Bool("ping", false, "Test ping and exit")

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage of %s:
	%s [OPTIONS] FILES...
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
	var dialOpt grpc.DialOption
	if len(*cert) > 0 && len(*certKey) > 0 {
		var cred tls.Certificate
		cred, err = tls.LoadX509KeyPair(*cert, *certKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Auth with x509:")
		log.Infof("    cert: %s", *cert)
		log.Infof("     key: %s", *certKey)
		dialOpt = grpc.WithTransportCredentials(credentials.NewServerTLSFromCert(&cred))
	} else {
		log.Warn("No keypair provided. Insecure.")
		dialOpt = grpc.WithInsecure()
	}
	conn, err := grpc.Dial(*server, dialOpt)
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

	if *ping {
		_, err = serv.Ping(context.Background(), &proto.PingRequest{
			EaselId:   eresp.EaselId,
			PaletteId: presp.PaletteId,
		})
		if err != nil {
			log.Errorf("Ping Failed: %v", err)
		} else {
			log.Info("Ping OK.")
		}
		return
	}

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
		/**** Render Image ****/
		for _, fname := range flag.Args() {
			input, src, err = util.LoadImage(fname)
			if err != nil {
				log.Fatal(err)
			}
			widthf := *scale * float64(src.Bounds().Dx())
			heightf := *scale * float64(src.Bounds().Dy())
			output, err = filters.RenderLanczos(serv, presp.EaselId, presp.PaletteId, input, src, int(widthf), int(heightf), float32(*quality), *mimeType)
			if err != nil {
				log.Fatal(err)
			}
			log.Infof("Rendered: (%s > %s) %s", presp.EaselId, presp.PaletteId, fname)
			outFilename := fmt.Sprintf("%s.out.png", strings.TrimSuffix(fname, path.Ext(fname)))
			err = ioutil.WriteFile(outFilename, output, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			log.Infof("Saved to %s, %d bytes", outFilename, len(output))
		}
	default:
		log.Fatalf("Unknown filter: %s", *filter)
	}

}
