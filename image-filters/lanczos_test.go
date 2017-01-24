package filters

import (
	"context"
	"log"
	"os"
	"sync"
	"testing"

	"google.golang.org/grpc"

	"github.com/ledyba/easel/easel-server/impl"
	"github.com/ledyba/easel/proto"
	"github.com/ledyba/easel/util"
)

func TestLanczos(t *testing.T) {
	// Create server
	em := impl.NewEaselMaker()
	var wg sync.WaitGroup
	wg.Add(2)
	go (func() {
		defer wg.Done()
		util.StartupTest(t)
		defer util.ShutdownTest()
		em.Start()
	})()

	srv := impl.NewServer(em)
	go (func() {
		defer wg.Done()
		srv.Start(":8192")
	})()

	// Create client
	conn, err := grpc.Dial(":8192", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := proto.NewEaselServiceClient(conn)
	ctx := context.Background()

	resp1, err := cli.NewEasel(ctx, &proto.NewEaselRequest{
		EaselId: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	easelID := resp1.EaselId

	resp2, err := cli.NewPalette(ctx, &proto.NewPaletteRequest{
		EaselId: easelID,
	})
	if err != nil {
		log.Fatal(err)
	}
	paletteID := resp2.PaletteId

	UpdateLanczos(cli, easelID, paletteID, 10)
	data, img, err := util.LoadImage("github.com/ledyba/easel/test-images/momiji.jpg")
	if err != nil {
		log.Fatal(err)
	}
	dx, dy := img.Bounds().Dx(), img.Bounds().Dy()
	out, err := os.Create("github.com/ledyba/easel/image-filters/lanczos.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	RenderLanczos(cli, easelID, paletteID, data, img, dx*2, dy*2, 95, "image/png", out)

	em.Stop()
	srv.Stop()
	wg.Wait()
}
