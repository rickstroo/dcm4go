package dcm4go

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