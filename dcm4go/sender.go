package dcm4go

import (
	"fmt"
	"io"
	"time"
)

// A Sender can send a DICOM object to an AE.
type Sender struct {
	Opts *SenderOpts
}

// SenderOpts impact the behaviour of a Sender.
type SenderOpts struct {
	LocalAE        string
	ConnectTimeOut time.Duration // a zero value means no connect timeout
	WriteTimeOut   time.Duration // a zero value means no write timeout
	ReadTimeOut    time.Duration // a zero value means no read timeout
}

// Send sends a DICOM object to an AE.
// The address of the AE is of the format 'aetitle@host:port'.
func (sender *Sender) Send(reader io.Reader, remoteAE string) error {
	return fmt.Errorf("Sender.Send(): not implemented")
}

// Send sends a DICOM object to another AE using a default set of options.
// To gain more control over the sending, the user should create a Sender
// with the desired SenderOpts.
func Send(reader io.Reader, remoteAE string) error {
	opts := &SenderOpts{}
	sender := &Sender{Opts: opts}
	return sender.Send(reader, remoteAE)
}
