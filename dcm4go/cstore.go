package dcm4go

const (
	// LowPriority represents a low priority send
	LowPriority = 0x0002

	// MedPriority represents a medium priority send
	MedPriority = 0x0000

	// HighPriority represents a high priority send
	HighPriority = 0x0001
)

// NewCStoreResponse constructs a C-Store response message based on the C-Store request message
func NewCStoreResponse(assoc *Assoc, request *Object) (*Object, error) {

	// use the message id from the request as the message id responded to
	messageID, err := request.asShort(MessageIDTag, 0)
	if err != nil {
		return nil, err
	}

	// use the affected sop class uid from the request in the response
	affectedSOPClassUID, err := request.asString(AffectedSOPClassUIDTag, 0)
	if err != nil {
		return nil, err
	}

	// construct a command
	response := newObject()
	response.addUID(AffectedSOPClassUIDTag, affectedSOPClassUID)
	response.addShort(CommandFieldTag, "US", CStoreRSP)
	response.addShort(MessageIDBeingRespondedToTag, "US", messageID)
	response.addShort(CommandDataSetTypeTag, "US", 0x0101) // no data
	response.addShort(StatusTag, "US", 0x00)               // success

	// construct and return a message
	return response, nil
}

// NewCStoreRequest constructs a C-Echo request message
func NewCStoreRequest(assoc *Assoc, sopClassUID string, sopInstanceUID string, transferSyntaxUID string) (*PresContext, *Object, error) {

	// find the presentation context that was negotiated for this sop class
	pc, err := assoc.findAcceptedPresContextByAbstractSyntax(sopClassUID)
	if err != nil {
		return nil, nil, err
	}

	// generate a message id
	messageID := nextMessageID()

	// construct a request
	request := newObject()
	request.addUID(AffectedSOPClassUIDTag, sopClassUID)
	request.addShort(CommandFieldTag, "US", CStoreRQ)
	request.addShort(MessageIDTag, "US", messageID)
	request.addShort(PriorityTag, "US", MedPriority)
	request.addShort(CommandDataSetTypeTag, "US", 0x0000) // data will follow
	request.addUID(AffectedSOPInstanceUIDTag, sopInstanceUID)

	// construct and return a message
	return pc, request, nil
}
