package dcm4go

import (
	"log"
	"net"
	"os"
	"time"
)

// A Sender can send a DICOM object to an AE.
type Sender struct {
	Opts *SenderOpts
}

// SenderOpts impact the behaviour of a Sender.
type SenderOpts struct {
	Local          string        // a zero value means "DCMSND"
	ConnectTimeOut time.Duration // a zero value means no connect timeout
	WriteTimeOut   time.Duration // a zero value means no write timeout
	ReadTimeOut    time.Duration // a zero value means no read timeout
}

// Send sends a DICOM object to an AE.
// The address of the AE is of the format 'aetitle@host:port'.
func (sender *Sender) Send(paths []string, remote string) error {
	return sender.send(paths, remote)
}

// Send sends a DICOM object to another AE using a default set of options.
// To gain more control over the sending, the user should create a Sender
// with the desired SenderOpts.
func Send(paths []string, remote string) error {
	opts := &SenderOpts{}
	sender := &Sender{Opts: opts}
	return sender.Send(paths, remote)
}

func (sender *Sender) send(paths []string, remote string) error {

	// gather the capabilities
	capabilities, err := sender.readCapabilities(paths)
	if err != nil {
		return err
	}

	// connect to the remote
	assoc, err := sender.connect(sender.Opts.Local, remote, capabilities)
	if err != nil {
		return err
	}

	// ensure the association is released and closed
	defer func() {
		assoc.RequestRelease()
		assoc.Close()
	}()

	// send the files
	if err := sendFiles(paths, assoc); err != nil {
		return err
	}

	// all is well
	return nil
}

// connect connects the local AE to the remote AE
func (sender *Sender) connect(local string, remote string, capabilities []*Capability) (*RequestorAssoc, error) {

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

func (sender *Sender) sendFiles(paths []string, assoc *RequestorAssoc) error {
	for _, path := range paths {
		if err := sender.sendFile(path, assoc); err != nil {
			return err
		}
		log.Printf("sent file %q\n", path)
	}
	return nil
}

func (sender *Sender) sendFile(path string, assoc *RequestorAssoc) error {

	// open the file, which returns a reader
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// make sure we close the file upon exit
	defer file.Close()

	// send the object
	if err := assoc.Send(file); err != nil {
		return err
	}

	// all is well
	return nil
}

// readCapabilities reads the group two elements of a set of files to figure
// out what capabilities are required to send the file.
func (sender *Sender) readCapabilities(paths []string) ([]*Capability, error) {

	capabilities := make([]*Capability, 0, 5)

	for _, path := range paths {
		capability, err := sender.readCapability(path)
		if err != nil {
			return nil, err
		}
		log.Printf("capability is %v", capability)

		// add the capability if not already present
		if !capability.Contained(capabilities) {
			capabilities = append(capabilities, capability)
		}
	}
	return capabilities, nil
}

// readCapability reads the group two elements of a single file to figure
// out what capabilities are required to send the file.
func (sender *Sender) readCapability(path string) (*Capability, error) {

	// open the file, which returns a reader, defer a close
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// make sure we close the file upon exit
	defer file.Close()

	// read the group two attributes
	groupTwo, err := ReadGroupTwo(file, 0)
	if err != nil {
		return nil, err
	}

	// get the sop class uid of the stored object
	sopClassUID, err := groupTwo.AsString(MediaStorageSOPClassUIDTag, 0)
	if err != nil {
		return nil, err
	}
	log.Printf("sop class uid is %q\n", sopClassUID)

	// get the transfer syntax used to store the file
	transferSyntaxUID, err := groupTwo.AsString(TransferSyntaxUIDTag, 0)
	if err != nil {
		return nil, err
	}
	log.Printf("transfer syntax uid is %q\n", transferSyntaxUID)

	// all is well, return the sop class uid and the transfer syntax uid
	capability := &Capability{
		AbstractSyntax:   sopClassUID,
		TransferSyntaxes: []string{transferSyntaxUID},
	}
	return capability, nil
}
