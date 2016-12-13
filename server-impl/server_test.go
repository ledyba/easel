package server

import (
	"sync"
	"testing"

	"github.com/ledyba/easel/proto"
	context "golang.org/x/net/context"
)

func TestServer(t *testing.T) {
	em := NewEaselMaker()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		startup(t)
		defer shutdown()

		em.Start()
	}()

	serv := NewServer(em)
	var easelID string
	var paletteID string
	{
		resp, err := serv.NewEasel(context.Background(), &proto.NewEaselRequest{})
		if err != nil {
			t.Fatalf("Failed to create new easel: %v", err)
		}
		if resp == nil || resp.EaselId == "" {
			t.Fatal("Failed to create new easel")
		}
		easelID = resp.EaselId
	}
	{
		resp, err := serv.NewPalette(context.Background(), &proto.NewPaletteRequest{
			EaselId: easelID,
		})
		if err != nil || resp == nil {
			t.Fatal("Failed to create new palette", err, resp)
		}
		if resp.EaselId != easelID {
			t.Fatalf("Failed to create new palette. Wrong EaselID: %v vs %v", easelID, resp.EaselId)
		}
		paletteID = resp.PaletteId
	}
	{
		resp, err := serv.DeletePalette(context.Background(), &proto.DeletePaletteRequest{
			EaselId:   easelID,
			PaletteId: paletteID,
		})
		if err != nil || resp == nil {
			t.Fatal("Failed to create delete palette", err, resp)
		}
	}
	{
		resp, err := serv.DeleteEasel(context.Background(), &proto.DeleteEaselRequest{
			EaselId: easelID,
		})
		if err != nil || resp == nil {
			t.Fatal("Failed to create delete easel", err, resp)
		}
	}

	em.Stop()
	wg.Wait()
}
