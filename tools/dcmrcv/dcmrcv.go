package main

import (
	"log"
	"net"

	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		log.Fatalf("error is %v", err)
	}
}

var shutdown = false
var folder = "tmp/"

// the main function
func main() {

	// set variables, will parse from command line later
	addr := "localhost:4104"
	aeTitle := "DCMRCV"

	// bind to a listening port
	listener, err := net.Listen("tcp", addr)
	check(err)
	log.Printf("opened listener on %v\n", listener.Addr())

	// listen for connections
	for {

		// wait for connection
		log.Printf("waiting for connection on %v\n", listener.Addr())
		conn, err := listener.Accept()
		check(err)
		log.Printf("accepted connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())

		// handle the connection, eventually as a goroutine
		// wrap the call so that we can handle errors here
		func() {
			if err := handleConnection(conn, aeTitle); err != nil {
				log.Printf("error occured while handling connection, error is %v", err)
			}
		}()

		// if shut down, exit
		if shutdown {
			break
		}

		// break all the time for now
		break
	}

	check(listener.Close())
	log.Printf("closed listener on %v\n", listener.Addr())
}

// handleConnection does the actual work, returning an error if it cannot
func handleConnection(conn net.Conn, aeTitle string) error {

	// ensure the connection gets closed
	//defer conn.Close()

	// create some capabilities
	defaultTransferSyntaxes := []string{
		dcm4go.ImplicitVRLittleEndianUID,
		dcm4go.ExplicitVRLittleEndianUID,
		dcm4go.ExplicitVRBigEndianUID,
	}

	capabilities := dcm4go.NewCapabilities()
	capabilities.Add(
		dcm4go.NewCapability(
			dcm4go.VerificationUID,
			[]string{dcm4go.ImplicitVRLittleEndianUID},
		),
	)
	capabilities.Add(
		dcm4go.NewCapability(
			dcm4go.EnhancedXAImageStorageUID,
			defaultTransferSyntaxes,
		),
	)
	capabilities.Add(
		dcm4go.NewCapability(
			dcm4go.GeneralECGWaveformStorageUID,
			defaultTransferSyntaxes,
		),
	)

	// start the state machine
	return dcm4go.StartMachineForServiceProvider(conn, aeTitle, capabilities)
}

// 	// create an ae for this server
// 	ae := dcm4go.NewAE(aeTitle)
//
// 	// define some options for the association
// 	assocOpts := &dcm4go.AssocOpts{
// 		WriteTimeOut: 10 * time.Second,
// 		ReadTimeOut:  10 * time.Second,
// 		MaxBufLen:    16 * 1024,
// 	}
//
// 	// attempt to accepet an association
// 	assoc, err := ae.AcceptAssoc(conn, capabilities, assocOpts)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("accepted association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())
//
// 	// handle the requests
// 	for {
// 		request, err := assoc.ReadRequest()
// 		if err != nil {
// 			if err != io.EOF {
// 				return err
// 			}
// 			break
// 		}
// 		log.Printf("request is %v", request)
// 		if err := handleRequest(assoc, request); err != nil {
// 			return err
// 		}
// 	}
//
// 	log.Printf("released association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())
//
// 	// return success
// 	return nil
// }
//
// func handleRequest(assoc *dcm4go.AcceptorAssoc, request *dcm4go.Message) error {
// 	commandField, err := request.Command().AsShort(dcm4go.CommandFieldTag, 0)
// 	if err != nil {
// 		return err
// 	}
// 	switch commandField {
// 	case dcm4go.CEchoRQ:
// 		return handleEchoRequest(assoc, request)
// 	case dcm4go.CStoreRQ:
// 		return handleCStoreRequest(assoc, request)
// 	default:
// 		return dcm4go.ErrUnexpectedRequest
// 	}
// }
//
// func handleEchoRequest(assoc *dcm4go.AcceptorAssoc, request *dcm4go.Message) error {
//
// 	// create a response
// 	response, err := dcm4go.NewCEchoResponse(request)
// 	if err != nil {
// 		return err
// 	}
//
// 	// write the response
// 	if err := assoc.WriteResponse(response); err != nil {
// 		return err
// 	}
//
// 	// all is well
// 	return nil
// }
//
// func handleCStoreRequest(assoc *dcm4go.AcceptorAssoc, request *dcm4go.Message) error {
//
// 	// store the data
// 	if err := storeToFile(assoc, request); err != nil {
// 		return err
// 	}
//
// 	// create a response
// 	response, err := dcm4go.NewCStoreResponse(request)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("response is %v\n", response)
//
// 	// write the response
// 	err = assoc.WriteResponse(response)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// // storeToFile stores the DICOM object to a file
// func storeToFile(assoc *dcm4go.AcceptorAssoc, request *dcm4go.Message) error {
//
// 	// create a unique file name
// 	path := folder + uuid.New().String() + ".dcm" + ".tmp"
//
// 	// store the file
// 	if err := store(assoc, path, request); err != nil {
// 		return err
// 	}
//
// 	// rename the file
// 	if err := os.Rename(path, strings.TrimSuffix(path, ".tmp")); err != nil {
// 		return err
// 	}
//
// 	// return success
// 	return nil
// }
//
// // store stores the DICOM object to a file
// func store(assoc *dcm4go.AcceptorAssoc, path string, request *dcm4go.Message) error {
//
// 	// construct the file meta information
// 	fmi, err := dcm4go.CreateFileMetaInfo(&assoc.Assoc, request.PCID(), request.Command())
// 	if err != nil {
// 		return err
// 	}
//
// 	// open a new file
// 	file, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
//
// 	// ensure the file is closed in case of early termination
// 	defer file.Close()
//
// 	// write the file meta information
// 	if err := dcm4go.WriteFileMetaInfo(file, fmi); err != nil {
// 		return err
// 	}
//
// 	// copy the data
// 	num, err := io.Copy(file, request.DataReader())
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("copied %d bytes", num)
//
// 	// sync the file, because we really really really want
// 	// to be sure the file is flushed from memory and stored
// 	// to some storage media
// 	if err := file.Sync(); err != nil {
// 		return err
// 	}
//
// 	// return success
// 	return nil
// }
