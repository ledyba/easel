package main

import (
	log "github.com/Sirupsen/logrus"
)

func printStartupBanner() {
	log.Info("****************************************")
	log.Info("              Easel Client              ")
	log.Info("****************************************")
	log.Infof("Build at: %s", BuildAt())
	log.Infof("Git Revision: \n%s", DecodeGitRev())
	log.Info("****************************************")
}
