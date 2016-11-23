package main

import (
	"sync"
	"time"

	"github.com/ledyba/easel"
	"github.com/ledyba/easel/server/proto"

	context "golang.org/x/net/context"
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
	if len(req.Name) > 0 {
		easelEnt := serv.fetchEasel(req.Name)
		if easelEnt != nil {
			return &proto.PrepareEaselResponse{
				Name: req.Name,
			}, nil
		}
	}
	name := RandString(10)
	serv.makeEasel(name)
	resp := &proto.PrepareEaselResponse{}
	resp.Name = name
	return resp, nil
}

// SetupEasel
func (serv *Server) SetupEasel(ctx context.Context, req *proto.SetupEaselRequest) (*proto.SetupEaselResponse, error) {
	resp := &proto.SetupEaselResponse{}
	return resp, nil
}
