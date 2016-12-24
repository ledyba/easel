package impl

import (
	"runtime"
	"sync"
	"testing"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func startup(t *testing.T) {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		t.Fatal("Failed to init glfw", t)
	}
}

func shutdown() {
	glfw.Terminate()
	runtime.UnlockOSThread()
}

func TestEaselMaker(t *testing.T) {
	em := NewEaselMaker()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		startup(t)
		defer shutdown()
		em.Start()
	}()

	e := <-em.RequestNewEasel()
	if e == nil {
		t.Errorf("Request new easel, but got nil")
	}

	em.RequestDelEasel(e)

	em.Stop()
	wg.Wait()
}
