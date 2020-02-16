package main

import (
	"flag"
	"fmt"
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

func parseAddr(addr string) (string, string, error) {
	s := strings.Split(addr, "@")
	if len(s) != 2 {
		return "", "", fmt.Errorf("expected address of form 'ae@host:port', found '%v'", addr)
	}
	return s[0], s[1], nil
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

	echo1(remote)
	echo2(remote, local)
	echo3(remote, local)
}

// this is about the simplest way to ping a remote ae
func echo1(remote string) {
	check(dcm4go.Echo(remote))
}

// if one wants more control, create a echoer with options
func echo2(local string, remote string) {

	opts := &dcm4go.EchoerOpts{
		LocalAETitle:   local,
		ConnectTimeOut: 30 * time.Second,
		WriteTimeOut:   10 * time.Second,
		ReadTimeOut:    10 * time.Second,
	}

	echoer := &dcm4go.Echoer{
		Opts: opts,
	}

	check(echoer.Echo(remote))
}

// and now, implement using the underlying AE and Assoc APIs
func echo3(local string, remote string) {

	// create a local AE
	localAE := &dcm4go.AE{
		AETitle: local,
	}

	// parse the address
	remoteAETitle, remoteHostPort, err := parseAddr(remote)
	check(err)

	// create a remote AE
	remoteAE := &dcm4go.AE{
		AETitle: remoteAETitle,
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
	conn, err := net.Dial("tcp", remoteHostPort)
	check(err)
	log.Printf("opened connection from %v to %v\n", conn.LocalAddr(), conn.RemoteAddr())

	// ensure the connection get closed
	defer func() {
		check(conn.Close())
		log.Printf("closed connection from %v to %v\n", conn.LocalAddr(), conn.RemoteAddr())
	}()

	// create an association
	assoc, err := localAE.RequestAssoc(conn, remoteAE, capabilities, assocOpts)
	check(err)
	log.Printf("created association from %s to %s\n", assoc.CallingAETitle(), assoc.CalledAETitle())

	// ensure the association gets released
	defer func() {
		check(assoc.RequestRelease())
		log.Printf("released association from %s to %s\n", assoc.CallingAETitle(), assoc.CalledAETitle())
	}()

	// send the echo
	check(assoc.Echo())
}
