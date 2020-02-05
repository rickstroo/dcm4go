package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
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

	// set variables, will parse from command line later
	addr := "localhost:4104"
	aeTitle := "DCMRCV"
	folder := "tmp/"

	// create handlers for C-Echo and C-Store
	handlers := []dcm4go.Handler{
		&dcm4go.BasicCEchoHandler{},
		&MyCStoreHandler{Folder: folder},
	}

	server := &dcm4go.Server{
		Addr:     addr,
		AETitle:  aeTitle,
		Handlers: handlers,
	}

	check(server.ListenAndServe())
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
	pc *dcm4go.PresContext,
	request *dcm4go.Object,
	reader *dcm4go.PDataReader,
) error {

	// store the data
	if err := handler.StoreToFile(assoc, pc.ID(), request, reader); err != nil {
		return err
	}

	// create a response
	response, err := dcm4go.NewCStoreResponse(assoc, request)
	if err != nil {
		return err
	}
	fmt.Printf("response is %v\n", response)

	// write the response
	err = assoc.WriteResponse(pc.ID(), response, nil)
	if err != nil {
		return err
	}

	return nil
}

// StoreToFile stores the DICOM object to a file
func (handler *MyCStoreHandler) StoreToFile(assoc *dcm4go.Assoc, pcID byte, command *dcm4go.Object, pDataReader *dcm4go.PDataReader) error {

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
	if err := dcm4go.WriteFile(file, fmi, pDataReader); err != nil {
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
