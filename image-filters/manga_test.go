package filters

import (
	"log"
	"os"
	"testing"

	"github.com/ledyba/easel/proto"
	"github.com/ledyba/easel/util"
)

func TestManga(t *testing.T) {
	testFilter(t, func(t *testing.T, cli proto.EaselServiceClient, easelID, paletteID string) {
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
	})
}
