package impl

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"runtime"
	"sync"
	"time"

	"github.com/chai2010/webp"
	"golang.org/x/image/tiff"

	"github.com/ledyba/easel"
	"github.com/ledyba/easel/proto"
	"github.com/ledyba/easel/util"

	log "github.com/Sirupsen/logrus"
	context "golang.org/x/net/context"
)

var (
	// ErrEaselNotFound ...
	ErrEaselNotFound = errors.New("Easel not found")
	// ErrPaletteNotFound ...
	ErrPaletteNotFound = errors.New("Palette not found")
	// ErrVertexBufferNotFound ...
	ErrVertexBufferNotFound = errors.New("VertexBuffer not found")
)

const (
	// ExpiredDuration ...
	ExpiredDuration = time.Minute * 30
)

// Server ...
type Server struct {
	easelMaker *EaselMaker
	easelMutex *sync.Mutex
	easelMap   map[string]*EaselEntry
}

// EaselEntry ...
type EaselEntry struct {
	easel      *easel.Easel
	usedAt     time.Time
	mutex      *sync.Mutex
	paletteMap map[string]*PaletteEntry
}

func (ent *EaselEntry) lock() {
	ent.mutex.Lock()
	ent.usedAt = time.Now()
}
func (ent *EaselEntry) unlock() {
	ent.mutex.Unlock()
}

// PaletteEntry ...
type PaletteEntry struct {
	palette       *easel.Palette
	usedAt        time.Time
	mutex         *sync.Mutex
	vertexBuffers map[string]*easel.VertexBuffer
	indecies      *easel.VertexBuffer
}

func (ent *PaletteEntry) lock() {
	ent.mutex.Lock()
	ent.usedAt = time.Now()
}

func (ent *PaletteEntry) unlock() {
	ent.mutex.Unlock()
}

// NewServer ...
func NewServer(em *EaselMaker) *Server {
	return &Server{
		easelMaker: em,
		easelMutex: new(sync.Mutex),
		easelMap:   make(map[string]*EaselEntry),
	}
}

// StartGC ...
func (serv *Server) StartGC() {
	t := time.NewTicker(ExpiredDuration)
	log.Info("Start Easel GC timer.")
	for {
		select {
		case <-t.C:
			log.Info("Start Easel GC")
			serv.gc()
		}
	}
}

func (serv *Server) gc() {
	serv.easelMutex.Lock()
	defer serv.easelMutex.Unlock()

	now := time.Now()
	cnt := 0
	for key, e := range serv.easelMap {
		e.lock()
		defer e.unlock()
		if now.Sub(e.usedAt) >= ExpiredDuration {
			e.easel.Destroy()
			serv.deleteEasel(key)
			cnt++
			log.Warnf("Easel (%s) garbage collected.", key)
			continue
		}
	}
	log.Info("%d easels garbage collected.", cnt)
}

func (serv *Server) fetchEasel(name string) *EaselEntry {
	serv.easelMutex.Lock()
	defer serv.easelMutex.Unlock()
	entry := serv.easelMap[name]
	return entry
}

func (serv *Server) deleteEasel(name string) bool {
	serv.easelMutex.Lock()
	defer serv.easelMutex.Unlock()
	e, ok := serv.easelMap[name]
	if ok {
		delete(serv.easelMap, name)
		serv.easelMaker.RequestDelEasel(e.easel)
		return true
	}

	return false
}

func (serv *Server) makeEasel(name string) *EaselEntry {
	e := <-serv.easelMaker.RequestNewEasel()
	ent := &EaselEntry{
		easel:      e,
		usedAt:     time.Now(),
		mutex:      new(sync.Mutex),
		paletteMap: make(map[string]*PaletteEntry),
	}
	serv.easelMap[name] = ent
	return ent
}

// NewEasel ...
func (serv *Server) NewEasel(c context.Context, req *proto.NewEaselRequest) (*proto.NewEaselResponse, error) {
	if len(req.EaselId) > 0 {
		easelEnt := serv.fetchEasel(req.EaselId)
		if easelEnt != nil {
			return &proto.NewEaselResponse{
				EaselId: req.EaselId,
			}, nil
		}
	}
	name := util.RandString(10)
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	serv.makeEasel(name)
	resp := &proto.NewEaselResponse{}
	resp.EaselId = name
	return resp, nil
}

// DeleteEasel ...
func (serv *Server) DeleteEasel(c context.Context, req *proto.DeleteEaselRequest) (*proto.DeleteEaselResponse, error) {
	if found := serv.deleteEasel(req.EaselId); found {
		return &proto.DeleteEaselResponse{}, nil
	}
	return &proto.DeleteEaselResponse{}, ErrEaselNotFound
}

// NewPalette ...
func (serv *Server) NewPalette(ctx context.Context, req *proto.NewPaletteRequest) (*proto.NewPaletteResponse, error) {
	var err error
	ent := serv.fetchEasel(req.EaselId)
	if ent == nil {
		return nil, ErrEaselNotFound
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	ent.lock()
	defer ent.unlock()
	ent.easel.MakeCurrent()
	defer ent.easel.DetachCurrent()
	palette, err := ent.easel.NewPalette()
	if err != nil {
		return nil, err
	}
	name := util.RandString(10)
	ent.paletteMap[name] = &PaletteEntry{
		palette:       palette,
		usedAt:        time.Now(),
		mutex:         new(sync.Mutex),
		vertexBuffers: make(map[string]*easel.VertexBuffer),
	}
	resp := &proto.NewPaletteResponse{
		EaselId:   req.EaselId,
		PaletteId: name,
	}
	return resp, nil
}

// DeletePalette ...
func (serv *Server) DeletePalette(ctx context.Context, req *proto.DeletePaletteRequest) (*proto.DeletePaletteResponse, error) {
	easelEnt := serv.fetchEasel(req.EaselId)
	if easelEnt == nil {
		return nil, ErrEaselNotFound
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	easelEnt.lock()
	defer easelEnt.unlock()
	paletteEnt := easelEnt.paletteMap[req.PaletteId]
	if paletteEnt == nil {
		return nil, ErrPaletteNotFound
	}
	paletteEnt.lock()
	defer paletteEnt.unlock()
	easelEnt.easel.MakeCurrent()
	defer easelEnt.easel.DetachCurrent()
	paletteEnt.palette.Destroy()
	return &proto.DeletePaletteResponse{}, nil
}

func (serv *Server) updatePalette(e *easel.Easel, paletteEnt *PaletteEntry, req *proto.PaletteUpdate) error {
	var err error
	p := paletteEnt.palette
	err = p.Bind()
	if err != nil {
		return err
	}
	defer p.Unbind()
	/* program */
	if len(req.VertexShader) > 0 && len(req.FragmentShader) > 0 {
		var prog *easel.Program
		prog, err = e.CompileProgram(req.VertexShader, req.FragmentShader)
		if err != nil {
			return err
		}
		if p.Program() != nil {
			p.Program().Destroy()
		}
		p.AttachProgram(prog)
	}
	err = p.Program().Use()
	if err != nil {
		return err
	}
	defer p.Program().Unuse()

	/* ArrayBuffer */
	if req.Buffers != nil {
		var vb *easel.VertexBuffer
		for _, buf := range req.Buffers {
			if old, ok := paletteEnt.vertexBuffers[buf.Name]; ok {
				old.Destroy()
			}
			if buf.Data == nil || len(buf.Data) <= 0 {
				delete(paletteEnt.vertexBuffers, buf.Name)
				continue
			}
			vb, err = p.MakeArrayBuffer(buf.Data)
			if err != nil {
				return err
			}
			paletteEnt.vertexBuffers[buf.Name] = vb
		}
	}

	/* ArrayIndexBuffer */
	if req.Indecies != nil {
		var vb *easel.VertexBuffer
		indecies := make([]uint16, len(req.Indecies))
		for i, v := range req.Indecies {
			indecies[i] = uint16(v)
		}
		vb, err = p.AttachArrayIndexBuffer(indecies)
		if err != nil {
			return err
		}
		paletteEnt.indecies = vb
	}

	// Binding VertexAttrib
	if req.VertexArrtibutes != nil {
		var vb *easel.VertexBuffer
		for _, attrib := range req.VertexArrtibutes {
			vb = paletteEnt.vertexBuffers[attrib.BufferName]
			if vb == nil {
				return ErrVertexBufferNotFound
			}
			err = p.BindArrayAttrib(vb, paletteEnt.indecies, attrib.ArgumentName, attrib.ElementSize, attrib.Stride, attrib.Offset)
			if err != nil {
				return err
			}
		}
	}

	// Binding Uniforms
	var tex *easel.Texture2D
	var old *easel.Texture2D
	if req.UniformVariables != nil {
		for _, uni := range req.UniformVariables {
			if uni.Texture != nil {
				tex, _, err = e.LoadTexture2D(uni.Texture)
				if err != nil {
					return err
				}
				old, err = p.BindTexture(uni.Name, tex)
				if err != nil {
					return err
				}
				if old != nil {
					old.Destroy()
				}
			} else if uni.FloatValue != nil {
				err = p.BindUniformf(uni.Name, int(uni.FloatValue.ElementSize), uni.FloatValue.Data)
				if err != nil {
					return err
				}
			} else if uni.IntValue != nil {
				err = p.BindUniformi(uni.Name, int(uni.IntValue.ElementSize), uni.IntValue.Data)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// UpdatePalette ...
func (serv *Server) UpdatePalette(ctx context.Context, req *proto.UpdatePaletteRequest) (*proto.UpdatePaletteResponse, error) {
	var err error
	easelEnt := serv.fetchEasel(req.EaselId)
	if easelEnt == nil {
		return nil, ErrEaselNotFound
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	easelEnt.lock()
	defer easelEnt.unlock()
	easelEnt.easel.MakeCurrent()
	defer easelEnt.easel.DetachCurrent()
	paletteEnt := easelEnt.paletteMap[req.PaletteId]
	if paletteEnt == nil {
		return nil, ErrPaletteNotFound
	}
	paletteEnt.lock()
	defer paletteEnt.unlock()
	e := easelEnt.easel
	if req.Updates != nil {
		err = serv.updatePalette(e, paletteEnt, req.Updates)
		if err != nil {
			return nil, err
		}
	}
	return &proto.UpdatePaletteResponse{}, nil
}

// Ping ...
func (serv *Server) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PongResponse, error) {
	easelEnt := serv.fetchEasel(req.EaselId)
	if easelEnt == nil {
		return nil, ErrEaselNotFound
	}
	easelEnt.lock()
	defer easelEnt.unlock()
	paletteEnt := easelEnt.paletteMap[req.PaletteId]
	if paletteEnt == nil {
		return nil, ErrPaletteNotFound
	}
	paletteEnt.lock()
	defer paletteEnt.unlock()
	log.Info("Ping: ")
	return &proto.PongResponse{
		EaselId:   req.EaselId,
		PaletteId: req.PaletteId,
	}, nil
}

// Listup ...
func (serv *Server) Listup(ctx context.Context, req *proto.ListupRequest) (*proto.ListupResponse, error) {
	serv.easelMutex.Lock()
	defer serv.easelMutex.Unlock()

	easels := make([]*proto.EaselInfo, 0)
	for k, v := range serv.easelMap {
		info := &proto.EaselInfo{}
		info.Id = k
		info.UpdatedAt = v.usedAt.String()
		info.Palettes = make([]*proto.PaletteInfo, 0)
		(func() {
			v.lock()
			defer v.unlock()
			for paletteName, palette := range v.paletteMap {
				info.Palettes = append(info.Palettes, &proto.PaletteInfo{
					Id:        paletteName,
					UpdatedAt: palette.usedAt.String(),
				})
			}
		})()
	}
	return &proto.ListupResponse{
		Easels: easels,
	}, nil
}

// Render ...
func (serv *Server) Render(ctx context.Context, req *proto.RenderRequest) (*proto.RenderResponse, error) {
	now := time.Now()
	var err error
	easelEnt := serv.fetchEasel(req.EaselId)
	if easelEnt == nil {
		return nil, ErrEaselNotFound
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	easelEnt.lock()
	defer easelEnt.unlock()
	paletteEnt := easelEnt.paletteMap[req.PaletteId]
	if paletteEnt == nil {
		return nil, ErrPaletteNotFound
	}
	paletteEnt.lock()
	defer paletteEnt.unlock()
	e := easelEnt.easel
	p := paletteEnt.palette
	e.MakeCurrent()
	defer e.DetachCurrent()
	if req.Updates != nil {
		if err = serv.updatePalette(e, paletteEnt, req.Updates); err != nil {
			return nil, err
		}
	}
	size := image.Rect(0, 0, int(req.OutWidth), int(req.OutHeight))
	img, err := p.Render(size)
	if err != nil {
		return nil, err
	}
	var bytes bytes.Buffer
	writer := bufio.NewWriter(&bytes)
	err = saveImage(writer, img, req.OutFormat, req.OutQuality)
	if err != nil {
		return nil, err
	}
	err = writer.Flush()
	if err != nil {
		return nil, err
	}
	resp := &proto.RenderResponse{}
	resp.Output = bytes.Bytes()
	log.Infof("Rendered %dx%d %s (%d bytes) image, %fms elapsed.",
		img.Bounds().Dx(),
		img.Bounds().Dy(),
		req.OutFormat,
		bytes.Len(), (time.Now().Sub(now)).Seconds()*1000)
	return resp, nil
}

func saveImage(writer io.Writer, img image.Image, format string, quality float32) error {
	var err error
	switch format {
	case "image/png":
		err = png.Encode(writer, img)
		if err != nil {
			return err
		}
	case "image/jpeg":
		fallthrough
	case "image/jpg":
		err = jpeg.Encode(writer, img, &jpeg.Options{
			Quality: int(quality),
		})
		if err != nil {
			return err
		}
	case "image/webp":
		err = webp.Encode(writer, img, &webp.Options{
			Quality: quality,
		})
		if err != nil {
			return err
		}
	case "image/tiff":
		fallthrough
	case "image/x-tiff":
		err = tiff.Encode(writer, img, &tiff.Options{
			Compression: tiff.Deflate,
		})
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown mime-type: %s", format)
	}
	return nil
}
