// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	// DefaultClientAETitle defines the default AE title for the client
	DefaultClientAETitle = "DCMSND"
)

// A Client is a DICOM client.  In DICOM parlance, it is often referrned
// to erroneusly as an SCU or more accurately as a Requestor.
//
// A Client's zero value is a usable client.
type Client struct {

	// AETitle is the AE Title of the client.
	// If zero, a value of "DCMSND" will be used when establishing associations
	AETitle string

	// ConnectTimeout is the maximum time allowed for connection to a server
	// If zero, there is no connect timeout.
	ConnectTimeout time.Duration

	// ReadTimeout is the maximum time allowed for reading a response
	// If zero, there is no read timeout.
	ReadTimeout time.Duration

	// WriteTimeout is the maximum time allows for writing a request
	// If zero, there is no write timeout.
	WriteTimeout time.Duration
}

// Verify sends  DICOM verifiction request from a client to a server
func (client *Client) Verify(addr string) error {
	return client.verify(addr)
}

// verify implements the verification request
func (client *Client) verify(addr string) error {

	// define the required capabilities
	capabilities := []*Capability{
		&Capability{VerificationUID, []string{ImplicitVRLittleEndianUID}},
	}

	// setup the connection and association
	assoc, err := client.connect(addr, capabilities)
	if err != nil {
		return err
	}

	// make sure the association gets closed
	defer assoc.Close()

	// send a verification request
	if err := assoc.Verify(); err != nil {
		return err
	}

	// all is well
	return nil
}

func (client *Client) connect(addr string, capabilities []*Capability) (*RequestorAssoc, error) {

	// parse the address
	serverAETitle, serverHostPort, err := parseAddr(addr)
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
	fmt.Printf("connected to %v from %v\n", conn.RemoteAddr(), conn.LocalAddr())

	// define an application entity for managing dicom connections
	clientAETitle := getClientAETitle(client)
	local := NewAE(clientAETitle)
	fmt.Printf("local ae:%v\n", local)

	// define the the remote ae
	remote := NewAE(serverAETitle)
	fmt.Printf("remote ae:%v\n", remote)

	// request an association
	assoc, err := RequestAssoc(conn, local, remote, capabilities)
	if err != nil {
		return nil, err
	}
	fmt.Printf("negotiated association from %s to %s\n", local.AETitle(), remote.AETitle())

	// return the association
	return assoc, nil
}

func getClientAETitle(client *Client) string {
	if client.AETitle == "" {
		return DefaultClientAETitle
	}
	return client.AETitle
}

func parseAddr(addr string) (string, string, error) {
	s := strings.Split(addr, "@")
	if len(s) != 2 {
		return "", "", fmt.Errorf("expected address of form 'ae@host:port', found '%v'", addr)
	}
	return s[0], s[1], nil
}

// Send sends a DICOM store request from a client to a server.
func (client *Client) Send(addr string, paths []string) error {
	return client.send(addr, paths)
}

// send implements the send request
func (client *Client) send(addr string, paths []string) error {

	// gather the required capabilities for sending the files
	capabilities, err := readCapabilities(paths)
	if err != nil {
		return err
	}
	fmt.Printf("capabilities is %v\n", capabilities)

	// setup the connection and association
	assoc, err := client.connect(addr, capabilities)
	if err != nil {
		return err
	}

	// make sure the association gets closed
	defer assoc.Close()

	// send the files
	if err := sendFiles(paths, assoc); err != nil {
		return err
	}

	// all is well
	return nil
}

func readCapabilities(paths []string) ([]*Capability, error) {
	capabilities := make([]*Capability, 0, 5)
	for _, path := range paths {
		capability, err := readCapability(path)
		if err != nil {
			return nil, err
		}
		fmt.Printf("capability is %v\n", capability)

		// add the capability if not already present
		if !capability.Contained(capabilities) {
			capabilities = append(capabilities, capability)
		}
	}
	return capabilities, nil
}

func readCapability(path string) (*Capability, error) {

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
	fmt.Printf("sop class uid is %q\n", sopClassUID)

	// get the transfer syntax used to store the file
	transferSyntaxUID, err := groupTwo.AsString(TransferSyntaxUIDTag, 0)
	if err != nil {
		return nil, err
	}
	fmt.Printf("transfer syntax uid is %q\n", transferSyntaxUID)

	// all is well, return the sop class uid and the transfer syntax uid
	capability := &Capability{
		AbstractSyntax:   sopClassUID,
		TransferSyntaxes: []string{transferSyntaxUID},
	}
	return capability, nil
}

func sendFiles(paths []string, assoc *RequestorAssoc) error {
	for _, path := range paths {
		if err := sendFile(path, assoc); err != nil {
			return err
		}
		fmt.Printf("sent file %q\n", path)
	}
	return nil
}

func sendFile(path string, assoc *RequestorAssoc) error {

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
