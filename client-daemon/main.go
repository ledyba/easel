package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"

	log "github.com/Sirupsen/logrus"
)

var server *string = flag.String("server", "localhost:3000", "server to connect")
var filter *string = flag.String("filter", "lanczos", "Filternames.")
var workers *int = flag.Int("workers", 10, "workers to run")
var lobes *int = flag.Int("lobes", 10, "lobes parameter")
var help *bool = flag.Bool("help", false, "Print help and exit")

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage of %s:
	%s [OPTIONS] IN OUT
Options:
`, os.Args[0], os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	printStartupBanner()
	if *help {
		usage()
		return
	}

	restartChan := make(chan bool, *workers)
	for i := 0; i < *workers; i++ {
		restartChan <- true
	}
	for {
		select {
		case <-restartChan:
			go (func() {
				var err error
				w := newWorker()
				defer w.destroy()
				defer (func() {
					restartChan <- true
				})()
				err = w.connect()
				if err != nil {
					log.Errorf("[%d] Error on connect: %v", w.name, err)
					return
				}
				err = w.init()
				if err != nil {
					log.Errorf("[%d] Error on initialize: %v", w.name, err)
					return
				}
			})()
		}
	}

}
