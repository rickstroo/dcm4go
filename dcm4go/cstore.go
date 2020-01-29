package dcm4go

const (
	// LowPriority represents a low priority send
	LowPriority = 0x0002

	// MedPriority represents a medium priority send
	MedPriority = 0x0000

	// HighPriority represents a high priority send
	HighPriority = 0x0001
)

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

// NewCStoreResponse constructs a C-Echo response message based on the C-Echo request message
func NewCStoreResponse(assoc *Assoc, request *Message) (*Message, error) {

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
