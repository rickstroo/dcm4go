package dcm4go

import (
	"fmt"
)

// A BasicCEchoService is the default service provided by the library for handling DICOM C-Echo requests
type BasicCEchoService struct {
	BasicService
}

// NewBasicCEchoService creates and initializes a Basic C-Echo service
func NewBasicCEchoService() Service {
	service := newBasicService()
	service.addCapability(&Capability{VerificationUID, []string{ImplicitVRLittleEndianUID}})
	return service
}

func (service *BasicCEchoService) onCommand(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	reader *PDataReader,
) error {
	return service.BasicService.onCommand(assoc, presContext, command, reader)
}

func (service *BasicCEchoService) onRequest(
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
	fmt.Printf("response is %v\n", response)

	// write the response
	err = assoc.WriteResponse(response)
	if err != nil {
		return err
	}

	return nil
}

// NewCEchoResponse constructs a C-Echo response message based on the C-Echo request message
func NewCEchoResponse(assoc *AcceptorAssoc, request *Message) (*Message, error) {

	// use the same pc id as the request
	pcID := request.pcID

	// use the message id from the request as the message id responded to
	messageID, err := request.Command().asShort(MessageIDTag, 0)
	if err != nil {
		return nil, err
	}

	// construct a command
	command := newObject()
	command.addUID(AffectedSOPClassUIDTag, VerificationUID)
	command.addShort(CommandFieldTag, "US", CEchoRSP)
	command.addShort(MessageIDBeingRespondedToTag, "US", messageID)
	command.addShort(CommandDataSetTypeTag, "US", 0x0101)
	command.addShort(StatusTag, "US", 0x00)

	// construct and return a message
	return &Message{pcID, command, nil}, nil
}

// NewCEchoRequest constructs a C-Echo request message
func NewCEchoRequest(assoc *RequestorAssoc) (*Message, error) {

	// find the presentation context id that was negotiated for verification
	presContext, err := assoc.findAcceptedPresContextByAbstractSyntax(VerificationUID)
	if err != nil {
		return nil, err
	}

	// construct a command
	command := newObject()
	command.addUID(AffectedSOPClassUIDTag, VerificationUID)
	command.addShort(CommandFieldTag, "US", CEchoRQ)
	command.addShort(MessageIDTag, "US", nextMessageID())
	command.addShort(CommandDataSetTypeTag, "US", 0x0101)

	// construct and return a message
	return &Message{presContext.id, command, nil}, nil
}
