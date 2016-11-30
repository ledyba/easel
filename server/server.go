package main

import (
	"bufio"
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"runtime"
	"sync"
	"time"

	"github.com/chai2010/webp"

	"github.com/ledyba/easel"
	"github.com/ledyba/easel/proto"

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

func newServer(em *EaselMaker) *Server {
	return &Server{
		easelMaker: em,
		easelMutex: new(sync.Mutex),
		easelMap:   make(map[string]*EaselEntry),
	}
}

func (serv *Server) gc() {
	serv.easelMutex.Lock()
	defer serv.easelMutex.Unlock()
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
	name := RandString(10)
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
	ent := serv.fetchEasel(req.EaselId)
	if ent == nil {
		return nil, ErrEaselNotFound
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	ent.lock()
	defer ent.unlock()
	palette := ent.easel.NewPalette()
	name := RandString(10)
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
	easelEnt.easel.MakeCurrent()
	defer easelEnt.easel.DetachCurrent()
	paletteEnt.palette.Destroy()
	return &proto.DeletePaletteResponse{}, nil
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

	/* program */
	prog, err := e.CompileProgram(req.VertexShader, req.FragmentShader)
	if err != nil {
		return nil, err
	}
	if p.Program() != nil {
		p.Program().Destroy()
	}
	p.AttachProgram(prog, req.TextureName)

	/* ArrayBuffer */
	var vb *easel.VertexBuffer
	for _, buf := range req.Buffers {
		vb, err = p.AttachArrayBuffer(buf.Data)
		if err != nil {
			return nil, err
		}
		paletteEnt.vertexBuffers[buf.Name] = vb
	}

	/* ArrayIndexBuffer */
	indecies := make([]uint16, len(req.Indecies))
	for i, v := range req.Indecies {
		indecies[i] = uint16(v)
	}
	vb, err = p.AttachArrayIndexBuffer(indecies)
	if err != nil {
		return nil, err
	}
	paletteEnt.indecies = vb

	// Binding VertexAttrib
	for _, attrib := range req.VertexArrtibutes {
		vb = paletteEnt.vertexBuffers[attrib.BufferName]
		if vb == nil {
			return nil, ErrVertexBufferNotFound
		}
		err = p.BindArrayAttrib(vb, attrib.ArgumentName, attrib.ElementSize, attrib.Stride, attrib.Offset)
		if err != nil {
			return nil, err
		}
	}

	return &proto.UpdatePaletteResponse{}, nil
}

// Render ...
func (serv *Server) Render(ctx context.Context, req *proto.RenderRequest) (*proto.RenderResponse, error) {
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
	size := image.Rect(0, 0, int(req.OutWidth), int(req.OutHeight))
	tex, err := e.LoadTexture2D(req.Input)
	if err != nil {
		return nil, err
	}
	defer tex.Destroy()
	img, err := p.Render(paletteEnt.indecies, tex, size)
	if err != nil {
		return nil, err
	}
	resp := &proto.RenderResponse{}
	var bytes bytes.Buffer
	writer := bufio.NewWriter(&bytes)
	switch req.OutFormat {
	case "image/png":
		err = png.Encode(writer, img)
		if err != nil {
			return nil, err
		}
	case "image/jpeg":
	case "image/jpg":
		err = jpeg.Encode(writer, img, &jpeg.Options{
			Quality: int(req.OutQuality),
		})
		if err != nil {
			return nil, err
		}
	case "image/webp":
		err = webp.Encode(writer, img, &webp.Options{
			Quality: req.OutQuality,
		})
		if err != nil {
			return nil, err
		}
	}
	err = writer.Flush()
	if err != nil {
		return nil, err
	}
	resp.Output = bytes.Bytes()
	return resp, nil
}
