// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"io"
	"io/ioutil"
)

// A Handler handles a DICOM request.
type Handler interface {
	Capabilities() []*Capability
	HandleRequest(*Assoc, *PresContext, *Object, io.Reader) error
}

// A BasicCEchoHandler provides a handler for C-Echo requests
type BasicCEchoHandler struct {
}

// Capabilities returns the capabilities of a C-Echo handler
func (handler *BasicCEchoHandler) Capabilities() []*Capability {
	return []*Capability{&Capability{VerificationUID, []string{ImplicitVRLittleEndianUID}}}
}

// HandleRequest handles a DICOM C-Echo request
func (handler *BasicCEchoHandler) HandleRequest(
	assoc *Assoc,
	pc *PresContext,
	request *Object,
	reader io.Reader,
) error {

	// create a response
	response, err := NewCEchoResponse(request)
	if err != nil {
		return err
	}

	// write the response
	if err := assoc.writeMessage(pc, response, nil, nil); err != nil {
		return err
	}

	// all is well
	return nil
}

// A BasicCStoreHandler is the default service provided by the library for handling DICOM C-Store requests
type BasicCStoreHandler struct {
}

// Capabilities returns the capabilities of a C-Echo handler
func (handler *BasicCStoreHandler) Capabilities() []*Capability {
	defaultTransferSyntaxes := []string{
		ImplicitVRLittleEndianUID,
		ExplicitVRLittleEndianUID,
	}
	return []*Capability{
		&Capability{EnhancedXAImageStorageUID, defaultTransferSyntaxes},
		&Capability{GeneralECGWaveformStorageUID, defaultTransferSyntaxes},
	}
}

// HandleRequest handles a DICOM C-Store request
func (handler *BasicCStoreHandler) HandleRequest(
	assoc *Assoc,
	pc *PresContext,
	request *Object,
	reader io.Reader,
) error {

	// discard the data
	num, err := io.Copy(ioutil.Discard, reader)
	if err != nil {
		return err
	}
	fmt.Printf("discarded %d bytes\n", num)

	// create a response
	response, err := NewCStoreResponse(request)
	if err != nil {
		return err
	}
	fmt.Printf("response is %v\n", response)

	// write the response
	err = assoc.writeMessage(pc, response, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
