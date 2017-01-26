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

var em impl.EaselMaker

func TestMain(m *testing.M) {
	util.StartupTest()
	defer util.ShutdownTest()
	em = impl.NewEaselMakerMock(10)
	os.Exit(m.Run())
}

func TestLanczos(t *testing.T) {
	// Create server
	var wg sync.WaitGroup
	wg.Add(1)

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
	data, img, err := util.LoadImage("../test-images/momiji.png")
	if err != nil {
		log.Fatal(err)
	}
	dx, dy := img.Bounds().Dx(), img.Bounds().Dy()
	out, err := os.Create("../test-images/momiji.lanczos.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	RenderLanczos(cli, easelID, paletteID, data, img, dx*2, dy*2, 95, "image/png", out)

	em.Stop()
	srv.Stop()
	wg.Wait()
}
