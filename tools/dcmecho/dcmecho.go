// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains the implementation of a DICOM C-Echo SCU.

package main

import (
	"flag"
	"log"
	"net"
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

	echo(remote, local)
}

// echo implements using the underlying AE and APIs
func echo(remoteAddr string, localAddr string) {

	// create a local AE
	localAE := dcm4go.NewAE(localAddr)

	// define some options for the association
	assocOpts := &dcm4go.AssocOpts{
		WriteTimeOut: 10 * time.Second,
		ReadTimeOut:  10 * time.Second,
		MaxBufLen:    16 * 1024,
	}

	// set the transfer capabilities
	capabilities := dcm4go.NewCapabilities()
	capabilities.Add(
		dcm4go.NewCapability(
			dcm4go.VerificationUID,
			[]string{dcm4go.ImplicitVRLittleEndianUID},
		),
	)

	// create the remote AE
	remoteAE := dcm4go.NewAE(remoteAddr)

	// connect to the remote
	conn, err := net.Dial("tcp", remoteAE.Host+":"+remoteAE.Port)
	check(err)
	log.Printf(
		"opened connection from %v to %v",
		conn.LocalAddr(),
		conn.RemoteAddr(),
	)

	// ensure the connection gets closed
	defer func() {
		check(conn.Close())
		log.Printf(
			"closed connection from %v to %v",
			conn.LocalAddr(),
			conn.RemoteAddr(),
		)
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
		check(assoc.Release())
		log.Printf(
			"released association from %s to %s",
			assoc.CallingAETitle(),
			assoc.CalledAETitle(),
		)
	}()

	// send the echo
	check(assoc.Echo())
}
