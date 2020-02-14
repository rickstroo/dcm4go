package main

import (
	"flag"
	"log"
	"os"
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

	var local string
	var remote string
	var help bool

	flag.StringVar(&local, "local", "DCMSND", "specify ae title of the local AE")
	flag.StringVar(&remote, "remote", "DCMRCV@localhost:4104", "specify ae title, host and port in the form 'aet@host:port' of the remote AE")
	flag.BoolVar(&help, "help", false, "display usage")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// this is about the simplest way to ping a remote ae
	check(dcm4go.Echo(remote))

	// if one wants more control, create a echoer with options
	opts := &dcm4go.EchoerOpts{
		Local:          local,
		ConnectTimeOut: 30 * time.Second,
		WriteTimeOut:   10 * time.Second,
		ReadTimeOut:    10 * time.Second,
	}
	echoer := &dcm4go.Echoer{
		Opts: opts,
	}
	check(echoer.Echo(remote))
}
