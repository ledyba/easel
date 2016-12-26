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
	"sync"

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
var list = flag.Bool("list", false, "Listup canvas/easels and exit")
var bench = flag.Bool("bench", false, "Benchmark mode. We does not save image.")
var benchN = flag.Int("benchn", 10, "How many easels will be created.")
var benchM = flag.Int("benchm", 10, "How many duplicated images will be sent.")

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
	if (len(args) <= 0 && !*ping && !*list) || *help {
		usage()
		return
	}

	if *list {
		listup()
		return
	}

	if *ping {
		doPing()
		return
	}

	if *bench {
		var wg sync.WaitGroup
		wg.Add(*benchN)
		for i := 0; i < *benchN; i++ {
			go processImage(&wg, flag.Args(), true)
		}
		wg.Wait()
	} else {
		processImage(nil, flag.Args(), false)
	}

}

func listup() {
	var err error
	conn, serv := makeConnection()
	defer conn.Close()
	var resp *proto.ListupResponse
	resp, err = serv.Listup(context.Background(), &proto.ListupRequest{})
	if err != nil {
		log.Fatalf("Failed to listup easels: %v", err)
	}
	if len(resp.Easels) == 0 {
		log.Info("Currently, there are no easels.")
		return
	}
	for _, easel := range resp.Easels {
		log.Infof("Easel: %s (%s)", easel.Id, easel.UpdatedAt)
		if len(easel.Palettes) == 0 {
			log.Infof("  <no palettes>")
			continue
		}
		for _, palette := range easel.Palettes {
			log.Infof("  - Palette: %s ()", palette.Id, palette.UpdatedAt)
		}
	}
}

func doPing() {
	var err error
	conn, serv := makeConnection()
	defer conn.Close()
	easelID, paletteID, closer := makeEaselAndPalette(serv)
	defer closer()
	_, err = serv.Ping(context.Background(), &proto.PingRequest{
		EaselId:   easelID,
		PaletteId: paletteID,
	})
	if err != nil {
		log.Errorf("Ping Failed: %v", err)
	} else {
		log.Info("Ping OK.")
	}
}

func makeConnection() (*grpc.ClientConn, proto.EaselServiceClient) {
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
	serv := proto.NewEaselServiceClient(conn)
	return conn, serv
}

func makeEaselAndPalette(serv proto.EaselServiceClient) (string, string, func()) {
	var err error
	eresp, err := serv.NewEasel(context.Background(), &proto.NewEaselRequest{
		EaselId: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	easelID := eresp.EaselId

	/**** Create Easel ****/
	log.Printf("Easel Created: %s", easelID)
	closeEasel := func() {
		serv.DeleteEasel(context.Background(), &proto.DeleteEaselRequest{
			EaselId: easelID,
		})
		log.Printf("Easel Deleted: %s", easelID)
	}

	/**** Create Palette ****/
	presp, err := serv.NewPalette(context.Background(), &proto.NewPaletteRequest{
		EaselId: easelID,
	})
	if err != nil {
		closeEasel()
		log.Fatal(err)
	}
	paletteID := presp.PaletteId
	log.Printf("Palette Created: (%s > %s)", easelID, paletteID)
	closePalette := func() {
		serv.DeletePalette(context.Background(), &proto.DeletePaletteRequest{
			EaselId:   easelID,
			PaletteId: paletteID,
		})
		log.Printf("Palette Deleted: (%s > %s)", easelID, paletteID)
	}
	return easelID, paletteID, func() {
		closePalette()
		closeEasel()
	}
}

func processImage(wg *sync.WaitGroup, fnames []string, bench bool) {
	if wg != nil {
		defer wg.Done()
	}
	var err error
	conn, serv := makeConnection()
	defer conn.Close()
	easelID, paletteID, closer := makeEaselAndPalette(serv)
	defer closer()
	/**** Update Palette ****/
	switch *filter {
	case "lanczos":
		err = filters.UpdateLanczos(serv, easelID, paletteID, *lobes)
		if err != nil {
			log.Fatal(err)
		}
		/**** Render Image ****/
		var output []byte
		var input []byte
		var src image.Image
		m := 1
		if bench {
			m = *benchM
		}
		for i := 0; i < m; i++ {
			for _, fname := range fnames {
				input, src, err = util.LoadImage(fname)
				if err != nil {
					log.Fatal(err)
				}
				widthf := *scale * float64(src.Bounds().Dx())
				heightf := *scale * float64(src.Bounds().Dy())
				output, err = filters.RenderLanczos(serv, easelID, paletteID, input, src, int(widthf), int(heightf), float32(*quality), *mimeType)
				if err != nil {
					log.Fatal(err)
				}
				log.Infof("Rendered: (%s > %s) %s", easelID, paletteID, fname)
				outFilename := fmt.Sprintf("%s.out.%s", strings.TrimSuffix(fname, path.Ext(fname)), strings.TrimPrefix(*mimeType, "image/"))
				if !bench {
					err = ioutil.WriteFile(outFilename, output, os.ModePerm)
					if err != nil {
						log.Fatal(err)
					}
					log.Infof("Saved to %s, %d bytes", outFilename, len(output))
				}
			}
		}
	default:
		log.Fatalf("Unknown filter: %s", *filter)
	}
}
