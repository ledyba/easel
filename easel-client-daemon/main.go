package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/chai2010/webp"
	_ "golang.org/x/image/tiff"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

const (
	reqStatusEnqueued   = 0
	reqStatusInProgress = 1
	reqStatusDone       = 2
	reqStatusError      = 3
)

/* Server to work with */
var server = flag.String("server", "localhost:3000", "server to connect")
var cert = flag.String("cert", "", "cert file")
var certKey = flag.String("cert_key", "", "private key file")

var dbAddr = flag.String("db", "user:password@tcp(host:port)/dbname", "db address")

var workers = flag.Int("workers", 10, "workers to run")

/* Filter Flags */
var filter = flag.String("filter", "lanczos", "applied filter name.")
var lobes = flag.Int("filter_lobes", 10, "lobes parameter")

/* General */
var help = flag.Bool("help", false, "Print help and exit")

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

	/* messaging queue */
	requestQueue := make(chan *ResampleRequest, 100)
	notifyQueue := make(chan *ResampleRequest, 100)
	/* chan to controll worker counts */
	workerRestartChan := make(chan struct{}, *workers)
	for i := 0; i < *workers; i++ {
		workerRestartChan <- struct{}{}
	}
	fetcherRestartChan := make(chan struct{}, 1)
	fetcherRestartChan <- struct{}{}
	notifierRestartChan := make(chan struct{}, 1)
	notifierRestartChan <- struct{}{}
	retry := func(r *ResampleRequest) {
		r.ttl--
		if r.ttl < 0 {
			log.Errorf("TTL < 0. Give up. reqId=%d\n  src: %s\n  dst:%s", r.id, r.src, r.dst)
			notifyQueue <- r
			return
		}
		requestQueue <- r
	}
	for {
		select {
		case <-workerRestartChan:
			go (func() {
				var err error
				w := newWorker()
				defer w.destroy()
				defer (func() {
					log.Errorf("[%d] Disconnected. Retry in 5 secs...", w.name)
					time.Sleep(5 * time.Second)
					workerRestartChan <- struct{}{}
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
				sanity := 0
				var output []byte
				for {
					select {
					case r, ok := <-requestQueue:
						if !ok {
							log.Errorf("[%d] Buffer closed.", w.name)
							return
						}
						log.Infof("[%d] Start rendering. reqID=%d\n  src: %s\n  dst: %s", w.name, r.id, r.src, r.dst)
						output, err = w.render(r)
						if err != nil {
							log.Errorf("[%d] Rendering failed. reqID=%d\n  src: %s\n  dst: %s\n  err: %v", w.name, r.id, r.src, r.dst, err)
							r.err = err
							retry(r)
							sanity--
						} else {
							err = ioutil.WriteFile(r.dst, output, os.ModePerm)
							if err != nil {
								log.Errorf("[%d] Rendered successfully, but could not write file. reqID=%d\n  src: %s\n  dst: %s\n err: %v", w.name, r.id, r.src, r.dst, err)
								r.err = err
								retry(r)
								sanity--
								if sanity < 0 {
									return
								}
							} else {
								log.Infof("[%d] Well done! reqID=%d\n  src: %s\n  dst: %s", w.name, r.id, r.src, r.dst)
								notifyQueue <- r
							}
						}
					case <-w.pingTicker.C:
						err = w.ping()
						if err != nil {
							log.Errorf("[%d] Ping Failed (%s > %s).\n err: %v", w.name, w.easelID, w.paletteID, err)
							return
						}
						log.Errorf("[%d] Ping Succeeded (%s > %s).", w.name, w.easelID, w.paletteID)
					}
				}
			})()
		case <-fetcherRestartChan:
			go (func() {
				defer (func() {
					log.Error("DB Fetcher disconnected. Retry in 5 secs...")
					time.Sleep(5 * time.Second)
					fetcherRestartChan <- struct{}{}
				})()
				var db *sql.DB
				var err error
				db, err = sql.Open("mysql", *dbAddr)
				if err != nil {
					log.Errorf("Error on connecting DB: %v", err)
					return
				}
				defer db.Close()
				timer := time.NewTicker(time.Second * 60)
				var rows *sql.Rows
				for {
					select {
					case <-timer.C:
						err = (func() error {
							rows, err = db.Query("select `id`,`src`,`dst`,`dst_width`,`dst_height`,`dst_quality`,`dst_mimetype` from `resample_requests` where status = ?", reqStatusEnqueued)
							if err != nil {
								log.Errorf("Error on selecting db: %v", err)
								return err
							}
							defer rows.Close()
							for rows.Next() {
								r := ResampleRequest{}
								err = rows.Scan(&r.id, &r.src, &r.dst, &r.dstWidth, &r.dstHeight, &r.dstQuality, &r.dstMimeType)
								if err != nil {
									log.Errorf("Error on scanning db: %v", err)
									return err
								}
								var q sql.Result
								q, err = db.Exec("update `resample_requests` SET `status`=? where `id`=? and `status`=?", reqStatusInProgress, r.id, reqStatusEnqueued)
								if err != nil {
									log.Errorf("Error on updating db: %v", err)
									return err
								}
								c, _ := q.RowsAffected()
								if c == 1 {
									log.Infof("Request fetched. reqID=%d\n  src: %s\n  dst: %s", r.id, r.src, r.dst)
									r.ttl = 5
									requestQueue <- &r
								} else {
									log.Warnf("Request is stealed by anyone else. reqID=%d\n  src: %s\n  dst: %s", r.id, r.src, r.dst)
								}
							}
							return nil
						})()
						if err != nil {
							return
						}
					}
				}
			})()
		case <-notifierRestartChan:
			go (func() {
				defer (func() {
					log.Errorf("DB Notifier disconnected. Retry in 5 secs...")
					time.Sleep(5 * time.Second)
					notifierRestartChan <- struct{}{}
				})()
				var db *sql.DB
				var err error
				db, err = sql.Open("mysql", *dbAddr)
				if err != nil {
					log.Errorf("Error on connecting DB: %v", err)
					return
				}
				defer db.Close()
				for {
					select {
					case r := <-notifyQueue:
						var q sql.Result
						if r.err == nil {
							q, err = db.Exec("update `resample_requests` SET `status`=2 where `id`=?", r.id)
							if err != nil {
								log.Errorf("Error on updating db: %v", err)
								break
							}
							c, _ := q.RowsAffected()
							if c == 1 {
								log.Infof("Request updated. reqID=%d status=done. \n  src: %s\n  dst: %s", r.id, r.src, r.dst)
							} else {
								log.Warnf("Request already updated by anyone else. reqID=%d\n  src: %s\n  dst: %s", r.id, r.src, r.dst)
							}
						} else {
							q, err = db.Exec("update `resample_requests` SET `status`=3 where `id`=?", r.id)
							if err != nil {
								log.Errorf("Error on updating db: %v", err)
								break
							}
							c, _ := q.RowsAffected()
							if c == 1 {
								log.Infof("Request updated. reqID=%d status=err. \n  src: %s\n  dst: %s", r.id, r.src, r.dst)
							} else {
								log.Warnf("Request already updated by anyone else. reqID=%d\n  src: %s\n  dst: %s", r.id, r.src, r.dst)
							}
						}
					}
				}
			})()
		}
	}
}
