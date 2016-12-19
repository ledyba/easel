package main

import (
	"crypto/tls"
	"fmt"
	"image"
	"sync/atomic"

	log "github.com/Sirupsen/logrus"
	filters "github.com/ledyba/easel/image-filters"
	"github.com/ledyba/easel/util"

	"github.com/ledyba/easel/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Worker ...
type Worker struct {
	name      int
	conn      *grpc.ClientConn
	server    proto.EaselServiceClient
	easelID   string
	paletteID string
}

// ResampleRequest ...
type ResampleRequest struct {
	id         int
	src        string
	dst        string
	dstWidth   int
	dstHeight  int
	dstQuality float32
	status     int
	err        error
}

var workerCount int32

func newWorker() *Worker {
	name := int(atomic.AddInt32(&workerCount, 1))
	log.Infof("Worker Created: %d", name)
	return &Worker{
		name: name,
	}
}

func (w *Worker) connect() error {
	var err error
	var dialOpt grpc.DialOption
	if len(*cert) > 0 && len(*certKey) > 0 {
		var cred tls.Certificate
		cred, err = tls.LoadX509KeyPair(*cert, *certKey)
		if err != nil {
			return err
		}
		log.Info("Auth with x509:")
		log.Infof("    cert: %s", *cert)
		log.Infof("     key: %s", *certKey)
		dialOpt = grpc.WithTransportCredentials(credentials.NewServerTLSFromCert(&cred))
	} else {
		log.Warn("No keypair provided. Insecure.")
		dialOpt = grpc.WithInsecure()
	}
	w.conn, err = grpc.Dial(*server, dialOpt)
	if err != nil {
		return err
	}

	/**** Create Easel ****/
	w.server = proto.NewEaselServiceClient(w.conn)
	eresp, err := w.server.NewEasel(context.Background(), &proto.NewEaselRequest{
		EaselId: "",
	})
	if err != nil {
		return err
	}
	log.Infof("[%d] Easel Created: %s", w.name, eresp.EaselId)
	w.easelID = eresp.EaselId

	/**** Create Palette ****/
	presp, err := w.server.NewPalette(context.Background(), &proto.NewPaletteRequest{
		EaselId: eresp.EaselId,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("[%d] Palette Created: (%s > %s)", w.name, w.easelID, presp.PaletteId)
	w.paletteID = presp.PaletteId
	return nil
}

func (w *Worker) init() error {
	var err error
	/**** Update Palette ****/
	switch *filter {
	case filters.LanczosFilter:
		err = filters.UpdateLanczos(w.server, w.easelID, w.paletteID, *lobes)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown filter: %s", *filter)
	}
	return nil
}

func (w *Worker) render(req *ResampleRequest) ([]byte, error) {
	var err error
	/**** Let's Render ****/
	var output []byte
	var input []byte
	var src image.Image
	switch *filter {
	case filters.LanczosFilter:
		input, src, err = util.LoadImage(req.src)
		if err != nil {
			return nil, err
		}
		output, err = filters.RenderLanczos(w.server, w.easelID, w.paletteID, input, src, req.dstWidth, req.dstHeight, req.dstQuality)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unknown filter: %s", *filter)
	}
	return output, nil
}

func (w *Worker) destroy() {
	if len(w.paletteID) > 0 {
		w.server.DeletePalette(context.Background(), &proto.DeletePaletteRequest{
			EaselId:   w.easelID,
			PaletteId: w.paletteID,
		})
		log.Infof("[%d] Palette Deleted: (%s > %s)", w.name, w.easelID, w.paletteID)
		w.paletteID = ""
	}
	if len(w.easelID) > 0 {
		w.server.DeleteEasel(context.Background(), &proto.DeleteEaselRequest{
			EaselId: w.easelID,
		})
		log.Infof("[%d] Easel Deleted: %s", w.name, w.easelID)
		w.easelID = ""
	}
	w.conn.Close()
	log.Infof("[%d] Worker Destoyed", w.name)
}
