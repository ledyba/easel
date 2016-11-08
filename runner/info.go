package main

import (
	"encoding/base64"
	"fmt"

	log "github.com/Sirupsen/logrus"
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
