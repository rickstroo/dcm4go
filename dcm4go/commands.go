package dcm4go

const (
	// CEchoRQ is command field value for C-Echo request
	CEchoRQ = 0x0030
	// CEchoRSP is command field value for C-Recho response
	CEchoRSP = 0x8030
	// CStoreRQ is command field value for C-Store request
	CStoreRQ = 0x0001
	// CStoreRSP is command field value for C-Store response
	CStoreRSP = 0x8001
)

const (
	// SuccessStatusCode is value for success
	SuccessStatusCode = 0x00
)

const (
	// NoDataSetCode is value for no data set included
	NoDataSetCode = 0x0101
	// DataSetCode is value for data set
	DataSetCode = 0x0102
)

const (
	// LowPriorityCode is code for low priority
	LowPriorityCode = 0x0002
	// MediumPriorityCode is code for medium priority
	MediumPriorityCode = 0x0000
	// HighPriorityCode is code for high priority
	HighPriorityCode = 0x0001
)

// newRequest constructs a request message
func newRequest(
	affectedSOPClassUID string,
	commandField uint16,
	commandDataSetType uint16,
) *Message {
	request := newObject()
	request.addUID(AffectedSOPClassUIDTag, affectedSOPClassUID)
	request.addShort(CommandFieldTag, "US", commandField)
	request.addShort(MessageIDTag, "US", nextMessageID())
	request.addShort(CommandDataSetTypeTag, "US", commandDataSetType)
	message := &Message{
		command: request,
	}
	return message
}

// newResponse constructs a response message
func newResponse(
	request *Message,
	commandDataSetType uint16,
	statusCode uint16,
) (*Message, error) {

	// use the affected sop class uid from the request
	affectedSOPClassUID, err := request.Command().asString(AffectedSOPClassUIDTag, 0)
	if err != nil {
		return nil, err
	}

	// use the command field tag from the request
	commandField, err := request.Command().asShort(CommandFieldTag, 0)
	if err != nil {
		return nil, err
	}

	// turn it into a response command
	commandField |= 0x8000

	// use the message id from the request as the message id responded to
	messageIDBeingRespondedTo, err := request.Command().asShort(MessageIDTag, 0)
	if err != nil {
		return nil, err
	}

	// construct a response
	response := newObject()
	response.addUID(AffectedSOPClassUIDTag, affectedSOPClassUID)
	response.addShort(CommandFieldTag, "US", commandField)
	response.addShort(MessageIDBeingRespondedToTag, "US", messageIDBeingRespondedTo)
	response.addShort(CommandDataSetTypeTag, "US", commandDataSetType)
	response.addShort(StatusTag, "US", statusCode)

	// return the response
	message := &Message{
		command: response,
	}
	return message, nil
}

// NewCEchoRequest constructs a C-Echo request message
func NewCEchoRequest() *Message {
	return newRequest(VerificationUID, CEchoRQ, NoDataSetCode)
}

// NewCEchoResponse constructs a C-Echo response message based on the C-Echo request message
func NewCEchoResponse(request *Message) (*Message, error) {
	return newResponse(request, NoDataSetCode, SuccessStatusCode)
}

// NewCStoreRequest constructs a C-Store request message
func NewCStoreRequest(sopClassUID string, sopInstanceUID string) *Message {

	// construct a default request
	request := newRequest(sopClassUID, CStoreRQ, DataSetCode)

	// add the C-Store specifics
	request.Command().addShort(PriorityTag, "US", MediumPriorityCode)
	request.Command().addUID(AffectedSOPInstanceUIDTag, sopInstanceUID)

	// return the request
	return request
}

// NewCStoreResponse constructs a C-Store response message based on the C-Store request message
func NewCStoreResponse(request *Message) (*Message, error) {

	// construct a default response
	response, err := newResponse(request, NoDataSetCode, SuccessStatusCode)
	if err != nil {
		return nil, err
	}

	// use the affected sop instance uid from the request in the response
	affectedSOPInstanceUID, err := request.Command().asString(AffectedSOPInstanceUIDTag, 0)
	if err != nil {
		return nil, err
	}
	response.Command().addUID(AffectedSOPInstanceUIDTag, affectedSOPInstanceUID)

	// return the response
	return response, nil
}
