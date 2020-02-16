package main

import (
	"flag"
	"log"
	"net"
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

	// and for even more control, one can create AEs
	// and manage the association completion
	localAE := &dcm4go.AE{
		AETitle: local,
	}
	remoteAE := &dcm4go.AE{
		AETitle: remote,
	}

	// define some options for the association
	assocOpts := &dcm4go.AssocOpts{
		WriteTimeOut: 10 * time.Second,
		ReadTimeOut:  10 * time.Second,
		MaxBufLen:    16 * 1024,
	}

	// read the transfer capabilities from all the files
	capabilities, err := dcm4go.ReadCapabilities(paths)
	check(err)

	// open a connection
	conn, err := net.Dial("tcp", remote)
	check(err)

	// ensure the connection get closed
	defer func() {
		check(conn.Close())
	}()

	// create an association
	assoc, err := localAE.RequestAssoc(conn, remoteAE, capabilities, assocOpts)
	check(err)

	// ensure the association gets released
	defer func() {
		check(assoc.RequestRelease())
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
		if err := assoc.Store(file); err != nil {
			log.Printf("error while sending file, %v", err)
		}

		// close the file
		if err := file.Close(); err != nil {
			log.Printf("error while closing file, %v", err)
		}
	}

}
