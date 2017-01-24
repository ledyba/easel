package impl

import (
	"errors"

	"github.com/ledyba/easel"
)

// EaselMaker ...
type EaselMaker interface {
	RequestNewEasel() <-chan *easel.Easel
	RequestDelEasel(e *easel.Easel)
	Start()
	Stop()
}

// easelMakerImpl ...
type easelMakerImpl struct {
	stopChan chan struct{}
	newChan  chan (chan<- *easel.Easel)
	delChan  chan *easel.Easel
}

var (
	// ErreaselMakerImplAlreadyClosed ...
	ErreaselMakerImplAlreadyClosed = errors.New("Easel maker is already stopped.")
)

// NewEaselMaker ...
func NewEaselMaker() EaselMaker {
	return &easelMakerImpl{
		stopChan: make(chan struct{}, 1),
		newChan:  make(chan (chan<- *easel.Easel), 10),
		delChan:  make(chan *easel.Easel, 10),
	}
}

// RequestNewEasel ...
func (em *easelMakerImpl) RequestNewEasel() <-chan *easel.Easel {
	ch := make(chan *easel.Easel, 1)
	em.newChan <- ch
	return ch
}

// RequestDelEasel ...
func (em *easelMakerImpl) RequestDelEasel(e *easel.Easel) {
	em.delChan <- e
}

// Start ...
func (em *easelMakerImpl) Start() {
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
func (em *easelMakerImpl) Stop() {
	select {
	case em.stopChan <- struct{}{}:
	}
}
