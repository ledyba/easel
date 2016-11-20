package main

import (
	"runtime"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	log "github.com/Sirupsen/logrus"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/ledyba/easel"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	var err error
	err = glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()
	log.Debug("Initialized.")

	printStartupBanner()
	s := easel.NewStudio()
	s.MakeEasel()
	defer s.Destroy()

}
