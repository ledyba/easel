package main

import (
	"errors"
	"sync"
	"time"

	"github.com/ledyba/easel"
	"github.com/ledyba/easel/proto"

	context "golang.org/x/net/context"
)

var (
	// ErrEaselNotFound ...
	ErrEaselNotFound = errors.New("Easel not found")
)

// Server ...
type Server struct {
	easelMutex *sync.Mutex
	easelMap   map[string]*EaselEntry
}

// EaselEntry ...
type EaselEntry struct {
	easel  *easel.Easel
	usedAt time.Time
	mutex  *sync.Mutex
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

func (serv *Server) makeEasel(name string) *EaselEntry {
	e := easel.NewEasel()
	ent := &EaselEntry{
		easel:  e,
		usedAt: time.Now(),
		mutex:  new(sync.Mutex),
	}
	serv.easelMutex.Lock()
	defer serv.easelMutex.Unlock()
	serv.easelMap[name] = ent
	return ent
}

// PrepareEasel ...
func (serv *Server) PrepareEasel(c context.Context, req *proto.PrepareEaselRequest) (*proto.PrepareEaselResponse, error) {
	if len(req.Id) > 0 {
		easelEnt := serv.fetchEasel(req.Id)
		if easelEnt != nil {
			return &proto.PrepareEaselResponse{
				Id: req.Id,
			}, nil
		}
	}
	name := RandString(10)
	serv.makeEasel(name)
	resp := &proto.PrepareEaselResponse{}
	resp.Id = name
	return resp, nil
}

// SetupEasel ...
func (serv *Server) SetupEasel(ctx context.Context, req *proto.SetupEaselRequest) (*proto.SetupEaselResponse, error) {
	ent := serv.fetchEasel(req.Id)
	if ent == nil {
		return nil, ErrEaselNotFound
	}
	resp := &proto.SetupEaselResponse{}
	return resp, nil
}
