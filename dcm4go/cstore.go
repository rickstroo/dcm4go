package dcm4go

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
