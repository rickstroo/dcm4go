package dcm4go

import (
	"fmt"
	"io"
	"io/ioutil"
)

// A Handler handles a DICOM request.
type Handler interface {
	Capabilities() []*Capability
	onCommand(
		assoc *AcceptorAssoc,
		presContext *PresContext,
		command *Object,
		reader *PDataReader,
	) error
	onRequest(
		assoc *AcceptorAssoc,
		presContext *PresContext,
		command *Object,
		data *Object,
	) error
}

// A BasicHandler is the default handler provided by the library
type BasicHandler struct {
}

// Capabilities returns no capabilities.
func (handler *BasicHandler) Capabilities() []*Capability {
	return nil
}

func (handler *BasicHandler) onCommand(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	reader *PDataReader,
) error {
	if reader != nil {
		// read the data
		data, err := readData(reader, &assoc.Assoc, presContext.id)
		if err != nil {
			return err
		}
		// then call the handler to manage the request
		return handler.onRequest(assoc, presContext, command, data)
	}
	return handler.onRequest(assoc, presContext, command, nil)
}

func (handler *BasicHandler) onRequest(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	data *Object,
) error {
	// hmm, need to do something
	return fmt.Errorf("BasicHandler.onRequest() not implemented")
}

// A BasicCEchoHandler provides a handler for C-Echo requests
type BasicCEchoHandler struct {
	BasicHandler
}

// Capabilities returns the capabilities of a C-Echo handler
func (handler *BasicCEchoHandler) Capabilities() []*Capability {
	return []*Capability{&Capability{VerificationUID, []string{ImplicitVRLittleEndianUID}}}
}

// onCommand handles the request
func (handler *BasicCEchoHandler) onCommand(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	reader *PDataReader,
) error {
	return handler.onRequest(assoc, presContext, command, nil)
}

func (handler *BasicCEchoHandler) onRequest(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	data *Object,
) error {

	// create a response
	response, err := NewCEchoResponse(assoc, &Message{presContext.id, command, nil})
	if err != nil {
		return err
	}

	// write the response
	if err := assoc.WriteResponse(response); err != nil {
		return err
	}

	// all is well
	return nil
}

// A BasicCStoreHandler is the default service provided by the library for handling DICOM C-Store requests
type BasicCStoreHandler struct {
	BasicHandler
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

func (handler *BasicCStoreHandler) onCommand(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	reader *PDataReader,
) error {

	// discard the data
	num, err := io.Copy(ioutil.Discard, reader)
	if err != nil {
		return err
	}
	fmt.Printf("discarded %d bytes\n", num)

	// create a response
	response, err := NewCStoreResponse(assoc, &Message{presContext.id, command, nil})
	if err != nil {
		return err
	}
	fmt.Printf("response is %v\n", response)

	// write the response
	err = assoc.WriteResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func (handler *BasicCStoreHandler) onRequest(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	data *Object,
) error {
	return fmt.Errorf("BasicCStoreHandler.OnRequest(): not implemented")
}
