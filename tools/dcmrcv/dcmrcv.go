package main

import (
	"fmt"

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

	addr := "localhost:4104"
	aeTitle := "DCMRCV"
	handlers := []dcm4go.Handler{&dcm4go.BasicCEchoHandler{}}

	server := &dcm4go.Server{
		Addr:     addr,
		AETitle:  aeTitle,
		Handlers: handlers,
	}

	check(server.ListenAndServe())

	// // listen for incoming connections
	// listener, err := net.Listen("tcp", "localhost:4104")
	// check(err)
	// fmt.Printf("listening on %v\n", listener.Addr())
	//
	// // ensure the listener gets closed
	// defer listener.Close()
	//
	// // define an application entity for managing dicom connections
	// // this includes adding handlers for the request, and for c-store requests,
	// // also for the command, as we want to manage the data ourselves
	// defaultTransferSyntaxes := []string{
	// 	dcm4go.ImplicitVRLittleEndianUID,
	// 	dcm4go.ExplicitVRLittleEndianUID,
	// 	dcm4go.ExplicitVRBigEndianUID,
	// }
	// ae := dcm4go.NewAE("DCMRCV")
	// ae.AddSupportedPresentationContext(dcm4go.VerificationUID, defaultTransferSyntaxes, nil)
	// ae.AddSupportedPresentationContext(dcm4go.EnhancedXAImageStorageUID, defaultTransferSyntaxes, &CStoreCommandHandler{"tmp/"})
	// ae.AddSupportedPresentationContext(dcm4go.GeneralECGWaveformStorageUID, defaultTransferSyntaxes, &CStoreCommandHandler{"tmp/"})
	// fmt.Printf("ae:%v\n", ae)
	//
	// // listen for connections
	// for {
	//
	// 	// wait for connection
	// 	fmt.Printf("waiting for connection on %v\n", listener.Addr())
	// 	conn, err := listener.Accept()
	// 	check(err)
	// 	fmt.Printf("accepted connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())
	//
	// 	// handle the connection, eventully as a goroutine
	// 	handleConnection(conn, ae)
	//
	// 	// break for now, so that we don't keep the port open
	// 	break
	// }
}

// // handleConnection handles the connection
// // in production, this will be called as a goroutine
// // that's why it does not return an error
// func handleConnection(conn net.Conn, ae *dcm4go.AE) {
//
// 	// accept the association
// 	assoc, err := dcm4go.AcceptAssoc(conn, ae)
// 	check(err)
// 	fmt.Printf("accepted association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())
//
// 	// handle the requests
// 	for {
//
// 		// read a request and call handlers as appropriate
// 		request, err := assoc.ReadRequest()
// 		if err != nil {
// 			if errors.Is(err, io.EOF) {
// 				break
// 			}
// 			check(err)
// 			break
// 		}
// 		fmt.Printf("request is %v\n", request)
//
// 		response, err := handleRequest(assoc, request)
// 		check(err)
// 		fmt.Printf("response is %v\n", response)
//
// 		err = assoc.WriteResponse(response)
// 		check(err)
// 	}
//
// 	// close the connection
// 	err = conn.Close()
// 	check(err)
// 	fmt.Printf("closed connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())
// }
//
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
//
// func handleVerificationRequest(assoc *dcm4go.Assoc, request *dcm4go.Message) (*dcm4go.Message, error) {
// 	return dcm4go.NewCEchoResponse(assoc, request)
// }
//
// func handleStoreRequest(assoc *dcm4go.Assoc, request *dcm4go.Message) (*dcm4go.Message, error) {
// 	return dcm4go.NewCStoreResponse(assoc, request)
// }
//
// // CStoreCommandHandler is a handler for C-Store commands
// type CStoreCommandHandler struct {
// 	folder string
// }
//
// // HandleCommand handles a C-Echo request
// func (handler *CStoreCommandHandler) HandleCommand(assoc *dcm4go.Assoc, pcID byte, command *dcm4go.Object, pDataReader *dcm4go.PDataReader) (*dcm4go.Object, error) {
//
// 	// construct the file meta information
// 	fmi, err := dcm4go.CreateFileMetaInfo(assoc, pcID, command)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// create a unique file name
// 	path := handler.folder + uuid.New().String() + ".dcm" + ".tmp"
//
// 	// open a new file
// 	file, err := os.Create(path)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// ensure the file is closed when we are finished
// 	defer file.Close()
//
// 	// write the file meta information
// 	if err := dcm4go.WriteFile(file, fmi, pDataReader); err != nil {
// 		return nil, err
// 	}
//
// 	// close file
// 	if err := file.Close(); err != nil {
// 		return nil, err
// 	}
//
// 	// rename the file
// 	if err := os.Rename(path, strings.TrimSuffix(path, ".tmp")); err != nil {
// 		return nil, err
// 	}
//
// 	// return success, it's okay to return nil for data if we manage it
// 	return nil, nil
// }
