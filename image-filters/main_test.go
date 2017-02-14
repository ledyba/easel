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

func testFilter(t *testing.T, f func(t *testing.T, cli proto.EaselServiceClient, easelID string, paletteID string)) {
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

	defer em.Stop()
	defer srv.Stop()
	defer wg.Wait()

	f(t, cli, easelID, paletteID)
}
