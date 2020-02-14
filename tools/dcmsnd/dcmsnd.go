package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		log.Fatalf("error is %v", err)
	}
}

// the main function
func main() {

	var path string
	var local string
	var remote string
	var help bool

	flag.StringVar(&path, "path", "image.dcm", "specify path of comma separated list of files to send")
	flag.StringVar(&local, "local", "DCMSND", "specify ae title of the local AE")
	flag.StringVar(&remote, "remote", "DCMRCV@localhost:4104", "specify ae title, host and port in the form 'aet@host:port' of the remote AE")
	flag.BoolVar(&help, "help", false, "display usage")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	paths := strings.Split(path, ",")

	// this is about the simplest way to send files
	check(dcm4go.Send(paths, remote))

	// if one wants more control, create a sender with options
	opts := &dcm4go.SenderOpts{
		Local:          local,
		ConnectTimeOut: 30 * time.Second,
		WriteTimeOut:   10 * time.Second,
		ReadTimeOut:    10 * time.Second,
	}
	sender := &dcm4go.Sender{
		Opts: opts,
	}
	check(sender.Send(paths, remote))
}
