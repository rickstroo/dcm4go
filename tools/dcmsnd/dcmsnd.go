package main

import (
	"fmt"
	"net"
	"os"

	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		fmt.Printf("panic: %v\n", err)
		os.Exit(0)
	}
}

// the main function
func main() {

	// attempt a connection
	conn, err := net.Dial("tcp", "localhost:4104")
	check(err)
	fmt.Printf("connected to %v from %v\n", conn.RemoteAddr(), conn.LocalAddr())

	// ensure the connection gets closed
	defer conn.Close()

	// define an application entity for managing dicom connections
	defaultTransferSyntaxes := []string{
		dcm4go.ImplicitVRLittleEndianUID,
		dcm4go.ExplicitVRLittleEndianUID,
		dcm4go.ExplicitVRBigEndianUID,
	}
	ae := dcm4go.NewAE("DCMSND")
	ae.AddRequestedPresentationContext(dcm4go.VerificationUID, defaultTransferSyntaxes)
	fmt.Printf("ae:%v\n", ae)

	// request an association
	assoc, err := dcm4go.RequestAssoc(conn, ae, "DCMRCV")
	check(err)

	// send a verification request
	check(assoc.Verify())

	// release association
	check(assoc.RequestRelease())
}
