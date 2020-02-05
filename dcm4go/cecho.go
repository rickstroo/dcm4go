package dcm4go

// NewCEchoResponse constructs a C-Echo response message based on the C-Echo request message
func NewCEchoResponse(assoc *Assoc, request *Object) (*Object, error) {

	// use the message id from the request as the message id responded to
	messageID, err := request.asShort(MessageIDTag, 0)
	if err != nil {
		return nil, err
	}

	// construct a command
	response := newObject()
	response.addUID(AffectedSOPClassUIDTag, VerificationUID)
	response.addShort(CommandFieldTag, "US", CEchoRSP)
	response.addShort(MessageIDBeingRespondedToTag, "US", messageID)
	response.addShort(CommandDataSetTypeTag, "US", 0x0101)
	response.addShort(StatusTag, "US", 0x00)

	// construct and return a message
	return response, nil
}

// NewCEchoRequest constructs a C-Echo request message
func NewCEchoRequest(assoc *Assoc) (*PresContext, *Object, error) {

	// find the presentation context id that was negotiated for verification
	pc, err := assoc.findAcceptedPresContextByAbstractSyntax(VerificationUID)
	if err != nil {
		return nil, nil, err
	}

	// construct a command
	request := newObject()
	request.addUID(AffectedSOPClassUIDTag, VerificationUID)
	request.addShort(CommandFieldTag, "US", CEchoRQ)
	request.addShort(MessageIDTag, "US", nextMessageID())
	request.addShort(CommandDataSetTypeTag, "US", 0x0101)

	// construct and return a message
	return pc, request, nil
}
