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

	store1(paths, remote)
	store2(paths, remote, local)
	store3(paths, remote, local)
}

// this is about the simplest way to send files
func store1(paths []string, remoteAddr string) {
	check(dcm4go.Send(paths, remoteAddr))
}

// if one wants more control, create a sender with options
func store2(paths []string, remoteAddr string, local string) {

	opts := &dcm4go.SenderOpts{
		LocalAETitle:   local,
		ConnectTimeOut: 30 * time.Second,
		WriteTimeOut:   10 * time.Second,
		ReadTimeOut:    10 * time.Second,
	}
	sender := &dcm4go.Sender{
		Opts: opts,
	}
	check(sender.Send(paths, remoteAddr))
}

// if one wants more control, create a sender with options
func store3(paths []string, remoteAddr string, local string) {

	// create a local AE
	localAE := dcm4go.NewAE(local)

	// define some options for the association
	assocOpts := &dcm4go.AssocOpts{
		WriteTimeOut: 10 * time.Second,
		ReadTimeOut:  10 * time.Second,
		MaxBufLen:    16 * 1024,
	}

	// read the transfer capabilities from all the files
	capabilities, err := dcm4go.ReadCapabilities(paths)
	check(err)

	// create the remote AE
	remoteAE := dcm4go.NewAE(remoteAddr)

	// create an association
	requestor, err := localAE.RequestAssoc(remoteAE, capabilities, assocOpts)
	check(err)
	log.Printf(
		"created association from %s to %s",
		requestor.Assoc().CallingAETitle(),
		requestor.Assoc().CalledAETitle(),
	)

	// ensure the association gets released
	defer func() {
		check(requestor.ReleaseAssoc())
		log.Printf("released association")
	}()

	// send the files
	for _, path := range paths {

		// open the file
		file, err := os.Open(path)
		if err != nil {
			log.Printf("error while opening file, %v", err)
			continue
		}

		// send the file
		if err := requestor.Store(file); err != nil {
			log.Printf("error while sending file, %v", err)
		}

		// close the file
		if err := file.Close(); err != nil {
			log.Printf("error while closing file, %v", err)
		}
	}
}
