// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"log"
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

	// conn is the connection between the client and server
	conn net.Conn

	// assoc is the association between the client and server
	assoc *RequestorAssoc
}

// Connect connects the client to a server
func (client *Client) Connect(addr string, services ...*Service) error {

	// parse the address
	serverAETitle, serverHostPort, err := parseAddr(addr)
	if err != nil {
		return err
	}

	// attempt a connection
	// we won't defer a close in this function
	// the connection will be closed when the association is closed
	conn, err := net.Dial("tcp", serverHostPort)
	if err != nil {
		return err
	}
	fmt.Printf("connected to %v from %v\n", conn.RemoteAddr(), conn.LocalAddr())

	// remember the connection
	client.conn = conn

	// define an application entity for managing dicom connections
	clientAETitle := getClientAETitle(client)
	local := &AE{AETitle: clientAETitle}
	fmt.Printf("local ae:%v\n", local)

	// define the the remote ae
	remote := &AE{AETitle: serverAETitle}
	fmt.Printf("remote ae:%v\n", remote)

	// gather the capabilities from the scus
	capabilities := make([]*Capability, 0, 5)
	// for _, scu := range scus {
	// 	for _, capability := range scu.Capabilities() {
	// 		if !capability.Contained(capabilities) {
	// 			capabilities = append(capabilities, capability)
	// 		}
	// 	}
	// }

	// // connect the scus to the client
	// for _, scu := range scus {
	// 	scu.client = client
	// }

	// request an association
	assoc, err := RequestAssoc(conn, local, remote, capabilities)
	if err != nil {
		return err
	}
	fmt.Printf("negotiated association from %s to %s\n", local.AETitle, remote.AETitle)

	// remember the association
	client.assoc = assoc

	// all is well
	return nil
}

// Verify sends  DICOM verifiction request from a client to a server
func (client *Client) Verify() error {

	// make sure there is an association
	if client.assoc == nil {
		return ErrNoAssoc
	}

	// send a verification request
	if err := client.assoc.Verify(); err != nil {
		return err
	}

	// all is well
	return nil
}

// Close releases the association.
// It can be called safely multiple times.
func (client *Client) Close() {

	// if the association is still open, attempt to close it
	if client.assoc != nil {
		if err := client.assoc.RequestRelease(); err != nil {
			log.Printf("error occured while attempting to release association, error is %v", err)
		}
	}

	// forget the association
	client.assoc = nil

	// if the connection is still open, attempt to close it
	if client.conn != nil {
		if err := client.conn.Close(); err != nil {
			log.Printf("error occured while attempting to close connect, error is %v", err)
		}
	}

	// forget the connection
	client.conn = nil
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
func (client *Client) Send(paths []string) error {

	// make sure there is an association
	if client.assoc == nil {
		return ErrNoAssoc
	}

	// send the files
	if err := sendFiles(paths, client.assoc); err != nil {
		return err
	}

	// all is well
	return nil
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

// ReadCapabilities reads the group two elements of a set of files to figure
// out what capabilities are required to send the file.
func ReadCapabilities(paths []string) ([]*Capability, error) {
	capabilities := make([]*Capability, 0, 5)
	for _, path := range paths {
		capability, err := ReadCapability(path)
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

// ReadCapability reads the group two elements of a single file to figure
// out what capabilities are required to send the file.
func ReadCapability(path string) (*Capability, error) {

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

// // An SCU is a Service Class User
// type SCU struct {
// 	client *Client
// }
//
// // Capabilities returns the capabilities required for an SCU
// func (scu *SCU) Capabilities() []*Capability {
// 	return nil
// }
//
// // A CEchoSCU is a Service Class User for C-Echo
// type CEchoSCU struct {
// 	client *Client
// }
//
// // Capabilities returns the capabilities required for a C-Echo SCU
// func (scu *CEchoSCU) Capabilities() []*Capability {
// 	return []*Capability{&Capability{VerificationUID, []string{ImplicitVRLittleEndianUID}}}
// }
//
// // A CStoreSCU is a Service Class User for C-Store SCU
// type CStoreSCU struct {
// 	client       *Client
// 	capabilities []*Capability
// }
//
// // Capabilities returns the capabilities required for a C-Echo SCU
// func (scu *CStoreSCU) Capabilities() []*Capability {
// 	return scu.capabilities
// }
//
// // ReadCapabilities reads the capabilities of a a set of files and
// // assigns it to the capabilities of the CStoreSCU
// func (scu *CStoreSCU) ReadCapabilities(paths []string) error {
// 	capabilities, err := ReadCapabilities(paths)
// 	if err != nil {
// 		return err
// 	}
// 	scu.capabilities = capabilities
// 	return nil
// }
