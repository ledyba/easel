package util

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// StartupTest ...
func StartupTest() {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		panic("Failed to init glfw")
	}
}

// ShutdownTest ...
func ShutdownTest() {
	glfw.Terminate()
	runtime.UnlockOSThread()
}
