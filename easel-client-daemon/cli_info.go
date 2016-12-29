package main

import log "github.com/Sirupsen/logrus"

//go:generate bash ../geninfo.sh

func printStartupBanner() {
	log.Info("****************************************")
	log.Info("          Easel Client Daemon           ")
	log.Info("****************************************")
	log.Infof("Build at: %s", BuildAt())
	log.Infof("Git Revision: \n%s", DecodeGitRev())
	log.Info("****************************************")
}
