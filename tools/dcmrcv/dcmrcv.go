package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
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
	}

	check(listener.Close())
	log.Printf("closed listener on %v\n", listener.Addr())
}

// handleConnection does the actual work, returning an error if it cannot
func handleConnection(conn net.Conn, aeTitle string) error {

	// ensure the connection gets closed
	defer conn.Close()

	// create some capabilities
	defaultTransferSyntaxes := []string{
		dcm4go.ImplicitVRLittleEndianUID,
		dcm4go.ExplicitVRLittleEndianUID,
		dcm4go.ExplicitVRBigEndianUID,
	}
	capabilities := []*dcm4go.Capability{
		&dcm4go.Capability{ // verification
			AbstractSyntax:   dcm4go.VerificationUID,
			TransferSyntaxes: defaultTransferSyntaxes,
		},
		&dcm4go.Capability{ // storage
			AbstractSyntax:   dcm4go.EnhancedXAImageStorageUID,
			TransferSyntaxes: defaultTransferSyntaxes,
		},
		&dcm4go.Capability{ // storage
			AbstractSyntax:   dcm4go.GeneralECGWaveformStorageUID,
			TransferSyntaxes: defaultTransferSyntaxes,
		},
	}

	// create an ae for this server
	ae := dcm4go.NewAE(aeTitle)

	// define some options for the association
	assocOpts := &dcm4go.AssocOpts{
		WriteTimeOut: 10 * time.Second,
		ReadTimeOut:  10 * time.Second,
		MaxBufLen:    16 * 1024,
	}

	// attempt to accepet an association
	assoc, err := ae.AcceptAssoc(conn, capabilities, assocOpts)
	if err != nil {
		return err
	}
	log.Printf("accepted association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())

	// handle the requests
	for {
		presContext, request, err := assoc.ReadRequest()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		log.Printf("request is %v", request)
		if err := handleRequest(assoc, presContext, request); err != nil {
			return err
		}
	}

	log.Printf("released association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())

	// return success
	return nil
}

func handleRequest(
	assoc *dcm4go.AcceptorAssoc,
	presContext *dcm4go.PresContext,
	request *dcm4go.Object,
) error {
	commandField, err := request.AsShort(dcm4go.CommandFieldTag, 0)
	if err != nil {
		return err
	}
	switch commandField {
	case dcm4go.CEchoRQ:
		return handleEchoRequest(assoc, presContext, request)
	case dcm4go.CStoreRQ:
		return handleCStoreRequest(assoc, presContext, request)
	default:
		return dcm4go.ErrUnexpectedRequest
	}
}

func handleEchoRequest(
	assoc *dcm4go.AcceptorAssoc,
	presContext *dcm4go.PresContext,
	request *dcm4go.Object,
) error {

	// create a response
	response, err := dcm4go.NewCEchoResponse(request)
	if err != nil {
		return err
	}

	// write the response
	if err := assoc.WriteResponse(presContext, response, nil); err != nil {
		return err
	}

	// all is well
	return nil
}

func handleCStoreRequest(
	assoc *dcm4go.AcceptorAssoc,
	presContext *dcm4go.PresContext,
	request *dcm4go.Object,
) error {

	// store the data
	if err := storeToFile(assoc, presContext, request); err != nil {
		return err
	}

	// create a response
	response, err := dcm4go.NewCStoreResponse(request)
	if err != nil {
		return err
	}
	fmt.Printf("response is %v\n", response)

	// write the response
	err = assoc.WriteResponse(presContext, response, nil)
	if err != nil {
		return err
	}

	return nil
}

// storeToFile stores the DICOM object to a file
func storeToFile(
	assoc *dcm4go.AcceptorAssoc,
	presContext *dcm4go.PresContext,
	command *dcm4go.Object,
) error {

	// construct the file meta information
	fmi, err := dcm4go.CreateFileMetaInfo(&assoc.Assoc, presContext.ID(), command)
	if err != nil {
		return err
	}

	// create a unique file name
	path := folder + uuid.New().String() + ".dcm" + ".tmp"

	// open a new file
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// ensure the file is closed in case of early termination
	defer file.Close()

	// get a reader for the data set
	reader, err := assoc.DataReader()
	if err != nil {
		return err
	}

	// write the file meta information
	if err := dcm4go.WriteFile(file, fmi, reader); err != nil {
		return err
	}

	// sync the file, because we really really really want
	// to be sure the file is flushed from memory and stored
	// to some storage media
	if err := file.Sync(); err != nil {
		return err
	}

	// close file
	if err := file.Close(); err != nil {
		return err
	}

	// rename the file
	if err := os.Rename(path, strings.TrimSuffix(path, ".tmp")); err != nil {
		return err
	}

	// return success
	return nil
}
