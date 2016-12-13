package server

import (
	"errors"

	"github.com/ledyba/easel"
)

// EaselMaker ...
type EaselMaker struct {
	stopChan chan interface{}
	newChan  chan (chan<- *easel.Easel)
	delChan  chan *easel.Easel
}

var (
	// ErrEaselMakerAlreadyClosed ...
	ErrEaselMakerAlreadyClosed = errors.New("Easel maker is already stopped.")
)

// NewEaselMaker ...
func NewEaselMaker() *EaselMaker {
	return &EaselMaker{
		stopChan: make(chan interface{}, 1),
		newChan:  make(chan (chan<- *easel.Easel), 10),
		delChan:  make(chan *easel.Easel, 10),
	}
}

// RequestNewEasel ...
func (em *EaselMaker) RequestNewEasel() <-chan *easel.Easel {
	ch := make(chan *easel.Easel, 1)
	em.newChan <- ch
	return ch
}

// RequestDelEasel ...
func (em *EaselMaker) RequestDelEasel(e *easel.Easel) {
	em.delChan <- e
}

// Start ...
func (em *EaselMaker) Start() {
	for {
		select {
		case e := <-em.delChan:
			e.Destroy()
		case c := <-em.newChan:
			c <- easel.NewEasel()
		case <-em.stopChan:
			close(em.newChan)
			close(em.delChan)
			return
		}
	}
}

// Stop ...
func (em *EaselMaker) Stop() {
	select {
	case em.stopChan <- nil:
	}
}
