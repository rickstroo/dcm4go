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

	// and for even more control, one can create an AE directly
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

	// set the transfer capabilities
	capabilities := []*dcm4go.Capability{
		&dcm4go.Capability{
			AbstractSyntax: dcm4go.VerificationUID,
			TransferSyntaxes: []string{
				dcm4go.ImplicitVRLittleEndianUID,
				dcm4go.ExplicitVRBigEndianUID,
				dcm4go.ExplicitVRLittleEndianUID,
			},
		},
	}

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

	// send the echo
	check(assoc.Echo())
}
