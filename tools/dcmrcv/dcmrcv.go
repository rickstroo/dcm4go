package main

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		fmt.Printf("panic: %v\n", err)
	}
}

// the main function
func main() {

	// listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:4104")
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
	ae := dcm4go.NewAE("DCMRCV")
	ae.AddSupportedPresentationContext(dcm4go.VerificationUID, defaultTransferSyntaxes, &CEchoRequestHandler{})
	ae.AddSupportedPresentationContext(dcm4go.EnhancedXAImageStorageUID, defaultTransferSyntaxes, &CStoreRequestHandler{})
	ae.AddSupportedPresentationContext(dcm4go.GeneralECGWaveformStorageUID, defaultTransferSyntaxes, &CStoreRequestHandler{})
	fmt.Printf("ae:%v\n", ae)

	// listen for connections
	for {

		// wait for connection
		fmt.Printf("waiting for connection on %v\n", listener.Addr())
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
		fmt.Printf("request is %v\n", request)

		//		response, err := handleRequest(assoc, request)
		response, err := assoc.HandleRequest(request)
		check(err)
		fmt.Printf("response is %v\n", response)

		err = assoc.WriteResponse(conn, response)
		check(err)
	}

	// close the connection
	err = conn.Close()
	check(err)
	fmt.Printf("closed connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())
}

// func handleRequest(assoc *dcm4go.Assoc, request *dcm4go.Message) (*dcm4go.Message, error) {
//
// 	commandField, err := request.Command().AsShort(dcm4go.CommandFieldTag, 0)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	switch commandField {
// 	case dcm4go.CEchoRQ:
// 		return handleVerificationRequest(assoc, request)
// 	case dcm4go.CStoreRQ:
// 		return handleStoreRequest(assoc, request)
// 	}
//
// 	return nil, fmt.Errorf("command field not recognized, 0x%02X", commandField)
// }

func handleVerificationRequest(assoc *dcm4go.Assoc, request *dcm4go.Message) (*dcm4go.Message, error) {
	return dcm4go.NewCEchoResponse(assoc, request)
}

func handleStoreRequest(assoc *dcm4go.Assoc, request *dcm4go.Message) (*dcm4go.Message, error) {

	// for now, just read the data set

	return dcm4go.NewCStoreResponse(assoc, request)
}

// CEchoRequestHandler is a handler for C-Echo requests
type CEchoRequestHandler struct {
}

// HandleRequest handles a C-Echo request
func (handler *CEchoRequestHandler) HandleRequest(assoc *dcm4go.Assoc, request *dcm4go.Message) (*dcm4go.Message, error) {
	fmt.Printf("CEchoRequestHandler\n")
	return dcm4go.NewCEchoResponse(assoc, request)
}

// CStoreRequestHandler is a handler for C-Store requests
type CStoreRequestHandler struct {
}

// HandleRequest handles a C-Echo request
func (handler *CStoreRequestHandler) HandleRequest(assoc *dcm4go.Assoc, request *dcm4go.Message) (*dcm4go.Message, error) {
	fmt.Printf("CStoreRequestHandler\n")
	return dcm4go.NewCStoreResponse(assoc, request)
}
