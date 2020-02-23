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

	stores(paths, remote, local)
}

// if one wants more control, create a sender with options
func stores(paths []string, remoteAddr string, local string) {

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

	// connect to the remote
	conn, err := net.Dial("tcp", remoteAE.Host()+":"+remoteAE.Port())
	check(err)
	log.Printf("opened connection from %v to %v", conn.LocalAddr(), conn.RemoteAddr())

	// ensure the connection gets closed
	defer func() {
		check(conn.Close())
		log.Printf("closed connection")
	}()

	// create an association
	assoc, err := localAE.RequestAssoc(conn, remoteAE, capabilities, assocOpts)
	check(err)
	log.Printf(
		"created association from %s to %s",
		assoc.CallingAETitle(),
		assoc.CalledAETitle(),
	)

	// ensure the association gets released
	defer func() {
		check(assoc.RequestRelease())
		log.Printf("released association")
	}()

	// send the files
	for _, path := range paths {
		if err := store(assoc, path); err != nil {
			log.Printf("error while sending file, %v", err)
		}
	}
}

// store sends a single file.
// we do this in a function so that we can use the defer
// statement to ensure that the file is closed.
func store(assoc *dcm4go.RequestorAssoc, path string) error {

	// open the file
	file, err := os.Open(path)
	if err != nil {
		log.Printf("error while opening file, %v", err)
		return err
	}

	// ensure the file gets closed
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error while closing file, %v", err)
		}
		log.Printf("closed file")
	}()

	// send the file
	if err := assoc.Store(file); err != nil {
		log.Printf("error while sending file, %v", err)
		return err
	}

	// close the file
	if err := file.Close(); err != nil {
		log.Printf("error while closing file, %v", err)
	}

	return nil
}
