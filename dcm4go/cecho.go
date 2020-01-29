package dcm4go

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

// NewCEchoResponse constructs a C-Echo response message based on the C-Echo request message
func NewCEchoResponse(assoc *Assoc, request *Message) (*Message, error) {

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
