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
	Local          string        // a zero value means "DCMSND"
	ConnectTimeOut time.Duration // a zero value means no connect timeout
	WriteTimeOut   time.Duration // a zero value means no write timeout
	ReadTimeOut    time.Duration // a zero value means no read timeout
}

// Echo Echos a DICOM object to an AE.
// The address of the AE is of the format 'aetitle@host:port'.
func (Echoer *Echoer) Echo(remote string) error {
	return Echoer.echo(remote)
}

// Echo Echos a DICOM object to another AE using a default set of options.
// To gain more control over the Echoing, the user should create a Echoer
// with the desired EchoerOpts.
func Echo(remote string) error {
	opts := &EchoerOpts{}
	Echoer := &Echoer{Opts: opts}
	return Echoer.Echo(remote)
}

func (Echoer *Echoer) echo(remote string) error {

	// gather the capabilities
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

	// connect to the remote
	assoc, err := Echoer.connect(Echoer.Opts.Local, remote, capabilities)
	if err != nil {
		return err
	}

	// ensure the association is released and closed
	defer func() {
		assoc.RequestRelease()
		assoc.Close()
	}()

	// Echo the ae
	if err := assoc.Verify(); err != nil {
		return err
	}

	// all is well
	return nil
}

// connect connects the local AE to the remote AE
func (Echoer *Echoer) connect(local string, remote string, capabilities []*Capability) (*RequestorAssoc, error) {

	// parse the address
	serverAETitle, serverHostPort, err := parseAddr(remote)
	if err != nil {
		return nil, err
	}

	// attempt a connection
	// we won't defer a close in this function
	// the connection will be closed when the association is closed
	conn, err := net.Dial("tcp", serverHostPort)
	if err != nil {
		return nil, err
	}
	log.Printf("connected to %v from %v\n", conn.RemoteAddr(), conn.LocalAddr())

	// define an application entity for managing dicom connections
	localAE := &AE{AETitle: local}

	// define the the remote ae
	remoteAE := &AE{AETitle: serverAETitle}

	// request an association
	assoc, err := RequestAssoc(conn, localAE, remoteAE, capabilities)
	if err != nil {
		return nil, err
	}
	log.Printf("negotiated association from %s to %s\n", localAE.AETitle, remoteAE.AETitle)

	// return the association
	return assoc, nil
}
