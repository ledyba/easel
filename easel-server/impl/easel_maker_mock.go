package impl

import "github.com/ledyba/easel"

type easelMakerMock struct {
	easels []*easel.Easel
}

// NewEaselMakerMock ...
func NewEaselMakerMock(neasels int) EaselMaker {
	easels := make([]*easel.Easel, neasels)
	for i := 0; i < neasels; i++ {
		easels[i] = easel.NewEasel()
	}
	return &easelMakerMock{
		easels: easels,
	}
}

// RequestNewEasel ...
func (em *easelMakerMock) RequestNewEasel() <-chan *easel.Easel {
	ch := make(chan *easel.Easel, 1)
	if len(em.easels) > 0 {
		e := em.easels[0]
		em.easels = em.easels[1:]
		ch <- e
	} else {
		ch <- nil
	}
	return ch
}

// RequestDelEasel ...
func (em *easelMakerMock) RequestDelEasel(e *easel.Easel) {
	e.Destroy()
}

// Start ...
func (em *easelMakerMock) Start() {
}

// Stop ...
func (em *easelMakerMock) Stop() {
}
