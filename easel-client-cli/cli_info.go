package main

import (
	"encoding/base64"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

var gitRev string
var buildAt string

// BuildAt ...
func BuildAt() string {
	return buildAt
}

// DecodeGitRev ...
func DecodeGitRev() string {
	data, err := base64.StdEncoding.DecodeString(gitRev)
	if err != nil {
		return fmt.Sprintf("<an error occured while reading git rev: %v>", err)
	}
	if len(data) == 0 {
		return "<not available>"
	}
	return string(data)
}

func printStartupBanner() {
	log.Info("****************************************")
	log.Info("              Easel Client              ")
	log.Info("****************************************")
	log.Infof("Build at: %s", BuildAt())
	log.Infof("Git Revision: \n%s", DecodeGitRev())
	log.Info("****************************************")
}
