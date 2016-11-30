package main

import "github.com/ledyba/easel"

// EaselMaker ...
type EaselMaker struct {
	newChan chan (chan<- *easel.Easel)
	delChan chan *easel.Easel
}

// NewEaselMaker ...
func NewEaselMaker() *EaselMaker {
	return &EaselMaker{
		newChan: make(chan (chan<- *easel.Easel)),
		delChan: make(chan *easel.Easel),
	}
}

// RequestNewEasel ...
func (em *EaselMaker) RequestNewEasel() <-chan *easel.Easel {
	ch := make(chan *easel.Easel)
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
		}
	}
}
