package main

import (
	"errors"
	"runtime"
	"sync"
	"time"

	"github.com/ledyba/easel"
	"github.com/ledyba/easel/proto"

	context "golang.org/x/net/context"
)

var (
	// ErrEaselNotFound ...
	ErrEaselNotFound = errors.New("Easel not found")
	// ErrPaletteNotFound ...
	ErrPaletteNotFound = errors.New("Palette not found")
)

// Server ...
type Server struct {
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

func newServer() *Server {
	return &Server{
		easelMutex: new(sync.Mutex),
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
		e.easel.Destroy()
		return true
	}

	return false
}

func (serv *Server) makeEasel(name string) *EaselEntry {
	e := easel.NewEasel()
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

	/* program */
	prog, err := e.CompileProgram(req.VertexShader, req.FragmentShader)
	if err != nil {
		return nil, err
	}
	if p.Program() != nil {
		p.Program().Destroy()
	}
	p.AttachProgram(prog)

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

	return &proto.UpdatePaletteResponse{}, nil
}
