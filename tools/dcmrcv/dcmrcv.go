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

// the main function
func main() {

	// set variables, will parse from command line later
	addr := "localhost:4104"
	aeTitle := "DCMRCV"
	folder := "tmp/"

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

		// create a c-store handler
		myCStoreHandler := &MyCStoreHandler{
			Folder: folder,
		}

		// create handlers for C-Echo and C-Store
		handlers := []dcm4go.Handler{
			&dcm4go.BasicCEchoHandler{},
			&MyCStoreHandler{Folder: folder},
		}

		// handle the connection, eventually as a goroutine
		handle(conn, aeTitle, handlers, myCStoreHandler)

		// if shut down, exit
		if shutdown {
			break
		}
	}

	check(listener.Close())
	log.Printf("closed listener on %v\n", listener.Addr())

	// // create handlers for C-Echo and C-Store
	// handlers := []dcm4go.Handler{
	// 	&dcm4go.BasicCEchoHandler{},
	// 	&MyCStoreHandler{Folder: folder},
	// }
	//
	// server := &dcm4go.Server{
	// 	Addr:     addr,
	// 	AETitle:  aeTitle,
	// 	Handlers: handlers,
	// }
	//
	// check(server.ListenAndServe())
}

func handle(conn net.Conn, aeTitle string, handlers []dcm4go.Handler, handler dcm4go.Handler) {

	// ensure the connection gets closed
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("while closing connection, caught error %v", err)
		} else {
			log.Printf("closed connection")
		}
	}()

	// gather all the capabilities of all the handlers
	//capabilities := handler.Capabilities()
	//log.Printf("capabilities is %v\n", capabilities)

	// create an ae for this server
	ae := dcm4go.NewAE(aeTitle)

	// define some options for the association
	assocOpts := &dcm4go.AssocOpts{
		WriteTimeOut: 10 * time.Second,
		ReadTimeOut:  10 * time.Second,
		MaxBufLen:    16 * 1024,
	}

	// attempt to accepet an association
	assoc, err := ae.AcceptAssoc(conn, handlers, assocOpts)
	if err != nil {
		log.Printf("error occured while attempting to accept association, %v", err)
		return
	}
	log.Printf("accepted association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())

	// handle the requests
	for {
		if err := assoc.Serve(handler); err != nil {
			if err != io.EOF {
				log.Printf("error occured while handling a request, %v", err)
				return
			}
			break
		}
	}

	log.Printf("closed association to %q from %q\n", assoc.CalledAETitle(), assoc.CallingAETitle())

	// all is well
	return

}

// A MyCStoreHandler handles C-Store requests
type MyCStoreHandler struct {
	Folder string // folder in which to store the DICOM objects
}

// Capabilities returns the capabilities of a my C-Store handler
func (handler *MyCStoreHandler) Capabilities() []*dcm4go.Capability {
	defaultTransferSyntaxes := []string{
		dcm4go.ImplicitVRLittleEndianUID,
		dcm4go.ExplicitVRLittleEndianUID,
		dcm4go.ExplicitVRBigEndianUID,
	}
	return []*dcm4go.Capability{
		&dcm4go.Capability{
			AbstractSyntax:   dcm4go.EnhancedXAImageStorageUID,
			TransferSyntaxes: defaultTransferSyntaxes,
		},
		&dcm4go.Capability{
			AbstractSyntax:   dcm4go.GeneralECGWaveformStorageUID,
			TransferSyntaxes: defaultTransferSyntaxes,
		},
	}
}

// HandleRequest handles a DICOM C-Store request
func (handler *MyCStoreHandler) HandleRequest(
	assoc *dcm4go.Assoc,
	presContext *dcm4go.PresContext,
	request *dcm4go.Object,
	reader io.Reader,
) error {

	commandField, err := request.AsShort(dcm4go.CommandFieldTag, 0)
	if err != nil {
		return err
	}
	if commandField == dcm4go.CEchoRQ {
		basicCEchoHandler := &dcm4go.BasicCEchoHandler{}
		return basicCEchoHandler.HandleRequest(assoc, presContext, request, reader)
	}

	// store the data
	if err := handler.StoreToFile(assoc, presContext.ID(), request, reader); err != nil {
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

// StoreToFile stores the DICOM object to a file
func (handler *MyCStoreHandler) StoreToFile(
	assoc *dcm4go.Assoc,
	pcID byte,
	command *dcm4go.Object,
	reader io.Reader,
) error {

	// construct the file meta information
	fmi, err := dcm4go.CreateFileMetaInfo(assoc, pcID, command)
	if err != nil {
		return err
	}

	// create a unique file name
	path := handler.Folder + uuid.New().String() + ".dcm" + ".tmp"

	// open a new file
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// ensure the file is closed when we are finished
	defer file.Close()

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

	// return success, it's okay to return nil for data if we manage it
	return nil
}
