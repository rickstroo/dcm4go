package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		fmt.Printf("panic: %v\n", err)
		//		panic(err)
	}
}

// parse the bind flag
func parseBind(bind string) (string, string, error) {
	atIndex := strings.Index(bind, "@")
	if atIndex == -1 {
		return "", "", fmt.Errorf("did not find @ separator")
	}
	aeTitle := bind[:atIndex]
	address := bind[atIndex+1:]
	return aeTitle, address, nil
}

// the main function
func main() {

	// get the binding
	bind := flag.String("bind", "DCMRCV@localhost:4104", "AE Title, Host and Port to bind to")
	flag.Parse()

	aeTitle, address, err := parseBind(*bind)
	check(err)
	fmt.Printf("aeTitle is %q, address is %q\n", aeTitle, address)

	// listen for incoming connections
	listener, err := net.Listen("tcp", address)
	check(err)
	fmt.Printf("listening on %v\n", listener.Addr())

	// ensure the listener gets closed
	defer listener.Close()

	// define an application entity for managing dicom connections
	defaultTransferSyntaxes := []string{
		dcm4go.ImplicitVRLittleEndianUID,
		dcm4go.ExplicitVRLittleEndianUID,
		dcm4go.ExplicitVRBigEndianUID,
	}
	ae := dcm4go.NewAE(aeTitle)
	ae.AddSupportedPresentationContext(dcm4go.VerificationUID, defaultTransferSyntaxes)
	ae.AddSupportedPresentationContext(dcm4go.EnhancedXAImageStorageUID, defaultTransferSyntaxes)
	fmt.Printf("ae:%v\n", ae)

	// listen for connections
	for {

		// wait for connection
		conn, err := listener.Accept()
		check(err)
		fmt.Printf("accepted connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())

		// handle the connection, eventully as a goroutine
		handleConnection(conn, ae)

		// break for now, so that we don't keep the port open
		break
	}
}

// handleConnection handles the connection
// in production, this will be called as a goroutine
// that's why it does not return an error
func handleConnection(conn net.Conn, ae *dcm4go.AE) {

	// accept the association
	assoc, err := dcm4go.AcceptAssoc(conn, ae)
	check(err)
	fmt.Printf("accepted association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())

	// handle the requests
	for {
		request, err := assoc.ReadRequest(conn)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			check(err)
			break
		}

		response, err := handleRequest(request)
		check(err)

		err = assoc.WriteResponse(conn, response)
		check(err)
	}

	// close the association
	err = assoc.Close()
	check(err)
	fmt.Printf("closed association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())

	// close the connection
	err = conn.Close()
	check(err)
	fmt.Printf("closed connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())
}

func handleRequest(request *dcm4go.Message) (*dcm4go.Message, error) {

	//	command := request.Command()

	return nil, fmt.Errorf("handleRequest: not implemented")
}
