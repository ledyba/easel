package util

import (
	"runtime"
	"testing"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// StartupTest ...
func StartupTest(t *testing.T) {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		t.Fatal("Failed to init glfw", t)
	}
}

// ShutdownTest ...
func ShutdownTest() {
	glfw.Terminate()
	runtime.UnlockOSThread()
}
