// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"log"
	"time"
)

// A Echoer can Echo a DICOM object to an AE.
type Echoer struct {
	Opts *EchoerOpts
}

// EchoerOpts impact the behaviour of a Echoer.
type EchoerOpts struct {
	LocalAETitle   string        // a zero value means "DCMSND"
	ConnectTimeOut time.Duration // a zero value means no connect timeout
	WriteTimeOut   time.Duration // a zero value means no write timeout
	ReadTimeOut    time.Duration // a zero value means no read timeout
}

// Echo echos a DICOM object to an AE.
// The address of the AE is of the format 'aetitle@host:port'.
func (echoer *Echoer) Echo(remoteAddr string) error {
	return echoer.echo(remoteAddr)
}

// Echo echos a DICOM object to another AE using a default set of options.
// To gain more control over the Echoing, the user should create a Echoer
// with the desired EchoerOpts.
func Echo(remoteAddr string) error {
	opts := &EchoerOpts{}
	Echoer := &Echoer{Opts: opts}
	return Echoer.Echo(remoteAddr)
}

func (echoer *Echoer) echo(remoteAddr string) error {

	// implement the default value for the LocalAETitle opt
	localAETitle := echoer.Opts.LocalAETitle
	if localAETitle == "" {
		localAETitle = "DCMSND"
	}

	// create an AE for the local
	localAE := NewAE(localAETitle)

	// define some options for the association
	assocOpts := &AssocOpts{
		WriteTimeOut: echoer.Opts.WriteTimeOut,
		ReadTimeOut:  echoer.Opts.ReadTimeOut,
		MaxBufLen:    16 * 1024,
	}

	// set the transfer capabilities
	capabilities := []*Capability{
		&Capability{
			AbstractSyntax:   VerificationUID,
			TransferSyntaxes: []string{ImplicitVRLittleEndianUID},
		},
	}

	// create the remote AE
	remoteAE := NewAE(remoteAddr)

	// create an association
	assoc, err := localAE.RequestAssoc(remoteAE, capabilities, assocOpts)
	if err != nil {
		return err
	}
	log.Printf(
		"created association from %s to %s",
		assoc.CallingAETitle(),
		assoc.CalledAETitle(),
	)

	// ensure the association gets released
	defer func() {
		if err := assoc.ReleaseAssoc(); err != nil {
			log.Printf("while releasing association, caught error %v", err)
		} else {
			log.Printf("released association")
		}
	}()

	// send the echo
	if err := assoc.Echo(); err != nil {
		return err
	}

	// return success
	return nil
}
