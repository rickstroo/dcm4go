// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
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

func parseAddr(addr string) (string, string, error) {
	s := strings.Split(addr, "@")
	if len(s) != 2 {
		return "", "", fmt.Errorf("expected address of form 'ae@host:port', found '%v'", addr)
	}
	return s[0], s[1], nil
}

// Verify sends  DICOM verifiction request from a client to a server
func (client *Client) Verify(addr string) error {
	return client.verify(addr)
}

// Verify creates a DICOM client and sends a verification request
// For greater control, create a client and use Client.Verify().
func Verify(addr string) error {
	client := &Client{}
	return client.verify(addr)
}

// verify implements the verification request
func (client *Client) verify(addr string) error {

	// parse the address
	serverAETitle, serverHostPort, err := parseAddr(addr)
	if err != nil {
		return err
	}

	// attempt a connection
	conn, err := net.Dial("tcp", serverHostPort)
	if err != nil {
		return err
	}
	fmt.Printf("connected to %v from %v\n", conn.RemoteAddr(), conn.LocalAddr())

	// ensure the connection gets closed
	defer conn.Close()

	// define an application entity for managing dicom connections
	clientAETitle := client.AETitle
	if clientAETitle == "" {
		clientAETitle = "DCMSND"
	}
	local := NewAE(clientAETitle)

	// request support for verification
	local.AddRequestedCapability(VerificationUID, []string{ImplicitVRLittleEndianUID})
	fmt.Printf("local ae:%v\n", local)

	// define the the remote ae
	remote := NewAE(serverAETitle)
	fmt.Printf("remote ae:%v\n", remote)

	// request an association
	assoc, err := RequestAssoc(conn, local, remote)
	if err != nil {
		return err
	}
	fmt.Printf("negotiated association from %s to %s\n", local.AETitle(), remote.AETitle())

	// send a verification request
	if err := assoc.Verify(); err != nil {
		return err
	}

	// all is well
	return nil
}

// Send sends a DICOM store request from a client to a server.
func (client *Client) Send(addr string, paths []string) error {
	return client.send(addr, paths)
}

// Send creates a DICOM client and sends a store request
// For greater control, create a client and use Client.Send().
func Send(addr string, paths []string) error {
	client := &Client{}
	return client.send(addr, paths)
}

// send implements the send request
func (client *Client) send(addr string, paths []string) error {

	// gather all the abstract syntax and transfer syntaxes required
	capabilities := make([]*Capability, 0, 5)
	for _, path := range paths {
		capability, err := readGroupTwo(path)
		if err != nil {
			return err
		}
		fmt.Printf("capability is %v\n", capability)
		capabilities = append(capabilities, capability)
	}
	fmt.Printf("capabilities is %v\n", capabilities)

	// parse the address
	serverAETitle, serverHostPort, err := parseAddr(addr)
	if err != nil {
		return err
	}

	// attempt a connection
	conn, err := net.Dial("tcp", serverHostPort)
	if err != nil {
		return err
	}
	fmt.Printf("connected to %v from %v\n", conn.RemoteAddr(), conn.LocalAddr())

	// ensure the connection gets closed
	defer conn.Close()

	// define an application entity for managing dicom connections
	clientAETitle := client.AETitle
	if clientAETitle == "" {
		clientAETitle = "DCMSND"
	}
	local := NewAE(clientAETitle)

	// add the capabilities
	for _, capability := range capabilities {
		local.AddRequestedCapability(
			capability.abstractSyntax,
			capability.transferSyntaxes,
		)
	}
	fmt.Printf("local ae:%v\n", local)

	// define the the remote ae
	remote := NewAE(serverAETitle)
	fmt.Printf("remote ae:%v\n", remote)

	// request an association
	assoc, err := RequestAssoc(conn, local, remote)
	if err != nil {
		return err
	}
	fmt.Printf("negotiated association from %s to %s\n", local.AETitle(), remote.AETitle())

	// send each file
	for _, path := range paths {
		if err := sendFile(path, assoc); err != nil {
			return err
		}
		fmt.Printf("sent file %q\n", path)
	}

	// all is well
	return nil
}

func readGroupTwo(path string) (*Capability, error) {

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
		abstractSyntax:   sopClassUID,
		transferSyntaxes: []string{transferSyntaxUID},
	}
	return capability, nil
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
