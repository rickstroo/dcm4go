package main

import (
	"flag"
	"fmt"
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
func echo2(remote string, local string) {

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
func echo3(remote string, local string) {

	// create a local AE
	localAE := &dcm4go.AE{
		AETitle: local,
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
			AbstractSyntax:   dcm4go.VerificationUID,
			TransferSyntaxes: []string{dcm4go.ImplicitVRLittleEndianUID},
		},
	}

	// create a requestor
	requestor := dcm4go.NewRequestor(localAE)

	// create an association
	check(requestor.RequestAssoc(remote, capabilities, assocOpts))
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

	// send the echo
	check(requestor.Echo())
}
