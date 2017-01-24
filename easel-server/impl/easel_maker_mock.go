package impl

import "github.com/ledyba/easel"

type easelMakerMock struct {
}

// NewEaselMakerMock ...
func NewEaselMakerMock() EaselMaker {
	return &easelMakerMock{}
}

// RequestNewEasel ...
func (em *easelMakerMock) RequestNewEasel() <-chan *easel.Easel {
	ch := make(chan *easel.Easel, 1)
	ch <- easel.NewEasel()
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
