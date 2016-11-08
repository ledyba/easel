package main

import (
	"encoding/base64"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var gitRev string
var buildAt string

func printStartupBanner() {
	log.Info("****************************************")
	log.Info("         ___  ___  ___  ___ (         ")
	log.Info("        |___)|   )|___ |___)|         ")
	log.Info("        |__  |__/| __/ |__  |         ")
	log.Info("****************************************")
	log.Infof("Build at: %s", buildAt)
	log.Infof("Git Revision: \n%s", decodeGitRev())
	log.Info("****************************************")
	log.Info("  <<GLFW>>")
	log.Infof("    Version: %s", glfw.GetVersionString())
	mons := glfw.GetMonitors()
	cmon := glfw.GetPrimaryMonitor()
	log.Infof("    Monitors: %d", len(mons))
	for i, mon := range mons {
		w, h := mon.GetPhysicalSize()
		x, y := mon.GetPos()
		cvm := mon.GetVideoMode()
		vms := mon.GetVideoModes()
		used := ""
		if cmon == mon {
			used = "** Primary **"
		}
		log.Infof("    --------------------------------")
		log.Infof("    [Monitor %d]%s", i, used)
		log.Infof("      Name: %s", mon.GetName())
		log.Infof("      PhysicalSize:  %dx%d", w, h)
		log.Infof("      Pos:          (%d,%d)", x, y)
		for j, vm := range vms {
			used := ""
			if vm == cvm {
				used = "** Current **"
			}
			log.Infof("      [VideoMode %d]%s", j, used)
			log.Infof("        Red/Green/Blue: %d/%d/%d", vm.RedBits, vm.GreenBits, vm.BlueBits)
			log.Infof("        Resolution: %dx%d(%d Hz)", vm.Width, vm.Height, vm.RefreshRate)
		}
	}
	log.Info("****************************************")
}

func decodeGitRev() string {
	data, err := base64.StdEncoding.DecodeString(gitRev)
	if err != nil {
		return fmt.Sprintf("<an error occured while reading git rev: %v>", err)
	}
	if len(data) == 0 {
		return "<not available>"
	}
	return string(data)
}
