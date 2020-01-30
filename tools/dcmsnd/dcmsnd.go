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
		panic(err)
	}
}

// the main function
func main() {

	// set the path to the file
	// will parse command line args later
	path := "/Users/Rick/data/dicom/ENHXA.dcm"

	// get the sop class uid and transfer syntax uid
	sopClassUID, transferSyntaxUID, err := readGroupTwo(path)
	check(err)

	// attempt a connection
	conn, err := net.Dial("tcp", "localhost:4104")
	check(err)
	fmt.Printf("connected to %v from %v\n", conn.RemoteAddr(), conn.LocalAddr())

	// ensure the connection gets closed
	defer conn.Close()

	// define an application entity for managing dicom connections
	// request support for verification
	// request support for storage for the type of object that we read
	local := dcm4go.NewAE("DCMSND")
	local.AddRequestedCapability(dcm4go.VerificationUID, []string{dcm4go.ImplicitVRLittleEndianUID})
	local.AddRequestedCapability(sopClassUID, []string{transferSyntaxUID})
	fmt.Printf("local ae:%v\n", local)

	// define the the remote ae
	remote := dcm4go.NewAE("DCMRCV")
	fmt.Printf("remote ae:%v\n", remote)

	// request an association
	assoc, err := dcm4go.RequestAssoc(conn, local, remote)
	fmt.Printf("negotated association from %s to %s\n", local.AETitle(), remote.AETitle())
	check(err)

	// send a verification request
	check(assoc.Verify())

	// send object
	check(send(path, assoc))

	// release association
	check(assoc.RequestRelease())
}

func readGroupTwo(path string) (string, string, error) {

	// open the file, which returns a reader, defer a close
	file, err := os.Open(path)
	check(err)

	// make sure we close the file upon exit
	defer file.Close()

	// read the group two attributes
	groupTwo, err := dcm4go.ReadGroupTwo(file, 0)
	if err != nil {
		return "", "", err
	}

	// get the sop class uid of the stored object
	sopClassUID, err := groupTwo.AsString(dcm4go.MediaStorageSOPClassUIDTag, 0)
	if err != nil {
		return "", "", err
	}
	fmt.Printf("sop class uid is %q\n", sopClassUID)

	// get the transfer syntax used to store the file
	transferSyntaxUID, err := groupTwo.AsString(dcm4go.TransferSyntaxUIDTag, 0)
	if err != nil {
		return "", "", err
	}
	fmt.Printf("transfer syntax uid is %q\n", transferSyntaxUID)

	// all is well, return the sop class uid and the transfer syntax uid
	return sopClassUID, transferSyntaxUID, nil
}

func send(path string, assoc *dcm4go.RequestorAssoc) error {

	// open the file, which returns a reader
	file, err := os.Open(path)
	check(err)

	// make sure we close the file upon exit
	defer file.Close()

	// send the object
	if err := assoc.Send(file); err != nil {
		return err
	}

	// all is well
	return nil
}
