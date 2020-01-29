package dcm4go

import (
	"fmt"
	"io"
	"io/ioutil"
)

const (
	// LowPriority represents a low priority send
	LowPriority = 0x0002

	// MedPriority represents a medium priority send
	MedPriority = 0x0000

	// HighPriority represents a high priority send
	HighPriority = 0x0001
)

// A BasicCStoreService is the default service provided by the library for handling DICOM C-Store requests
type BasicCStoreService struct {
	BasicService
}

func (service *BasicCStoreService) onClose(assoc *AcceptorAssoc) error {
	return service.BasicService.onClose(assoc)
}

func (service *BasicCStoreService) onCommand(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	reader *PDataReader,
) error {

	// discard the data
	_, err := io.Copy(ioutil.Discard, reader)
	if err != nil {
		return err
	}
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

func (service *BasicCStoreService) onRequest(
	assoc *AcceptorAssoc,
	presContext *PresContext,
	command *Object,
	data *Object,
) error {
	return fmt.Errorf("BasicCStoreService.onRequest not impemented")
}

// NewCStoreResponse constructs a C-Store response message based on the C-Store request message
func NewCStoreResponse(assoc *AcceptorAssoc, request *Message) (*Message, error) {

	// use the same pc id as the request
	pcID := request.pcID

	// use the message id from the request as the message id responded to
	messageID, err := request.Command().asShort(MessageIDTag, 0)
	if err != nil {
		return nil, err
	}

	// use the affected sop class uid from the request in the response
	affectedSOPClassUID, err := request.Command().asString(AffectedSOPClassUIDTag, 0)
	if err != nil {
		return nil, err
	}

	// construct a command
	command := newObject()
	command.addUID(AffectedSOPClassUIDTag, affectedSOPClassUID)
	command.addShort(CommandFieldTag, "US", CStoreRSP)
	command.addShort(MessageIDBeingRespondedToTag, "US", messageID)
	command.addShort(CommandDataSetTypeTag, "US", 0x0101) // no data
	command.addShort(StatusTag, "US", 0x00)               // success

	// construct and return a message
	return &Message{pcID, command, nil}, nil
}

// NewCStoreRequest constructs a C-Echo request message
func NewCStoreRequest(assoc *RequestorAssoc, sopClassUID string, sopInstanceUID string, transferSyntaxUID string) (*Message, error) {

	// find the presentation context that was negotiated for this sop class
	acPresContext, err := assoc.findAcceptedPresContextByAbstractSyntax(sopClassUID)
	if err != nil {
		return nil, err
	}

	// grab the presentation context id
	pcID := acPresContext.id

	// generate a message id
	messageID := nextMessageID()

	// construct a command
	command := newObject()
	command.addUID(AffectedSOPClassUIDTag, sopClassUID)
	command.addShort(CommandFieldTag, "US", CStoreRQ)
	command.addShort(MessageIDTag, "US", messageID)
	command.addShort(PriorityTag, "US", MedPriority)
	command.addShort(CommandDataSetTypeTag, "US", 0x0000) // data will follow
	command.addUID(AffectedSOPInstanceUIDTag, sopInstanceUID)

	// construct and return a message
	return &Message{pcID, command, nil}, nil
}
