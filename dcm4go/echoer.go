package dcm4go

import (
	"log"
	"net"
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
	localAE := &AE{
		AETitle: localAETitle,
	}

	// parse the address
	remoteAETitle, remoteHostPort, err := parseAddr(remoteAddr)
	if err != nil {
		return err
	}

	// create an AE for the remote
	remoteAE := &AE{
		AETitle: remoteAETitle,
	}

	// define some options for the association
	assocOpts := &AssocOpts{
		WriteTimeOut: echoer.Opts.WriteTimeOut,
		ReadTimeOut:  echoer.Opts.ReadTimeOut,
		MaxBufLen:    16 * 1024,
	}

	// set the transfer capabilities
	capabilities := []*Capability{
		&Capability{
			AbstractSyntax: VerificationUID,
			TransferSyntaxes: []string{
				ImplicitVRLittleEndianUID,
				ExplicitVRBigEndianUID,
				ExplicitVRLittleEndianUID,
			},
		},
	}

	// open a connection
	conn, err := net.Dial("tcp", remoteHostPort)
	if err != nil {
		return err
	}
	log.Printf("opened connection from %v to %v", conn.LocalAddr(), conn.RemoteAddr())

	// ensure the connection get closed
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("error while attempting to close connection from %v to %v, error is %v", conn.LocalAddr(), conn.RemoteAddr(), err)
		} else {
			log.Printf("closed connection from %v to %v", conn.LocalAddr(), conn.RemoteAddr())
		}
	}()

	// create an association
	assoc, err := localAE.RequestAssoc(conn, remoteAE, capabilities, assocOpts)
	if err != nil {
		return err
	}
	log.Printf("created association from %s to %s", assoc.CallingAETitle(), assoc.CalledAETitle())

	// ensure the association gets released
	defer func() {
		if err := assoc.RequestRelease(); err != nil {
			log.Printf("error while attempting to release association from %s to %s, error is %v", assoc.CallingAETitle(), assoc.CalledAETitle(), err)
		} else {
			log.Printf("released association from %s to %s", assoc.CallingAETitle(), assoc.CalledAETitle())
		}
	}()

	// send the echo
	if err := assoc.Echo(); err != nil {
		return err
	}

	// return success
	return nil
}
