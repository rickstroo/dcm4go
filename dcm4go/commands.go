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

// newCEchoRequest constructs a C-Echo request message
func newCEchoRequest(assoc *Assoc) (*Object, error) {

	// construct a request
	request := newObject()
	request.addUID(AffectedSOPClassUIDTag, VerificationUID)
	request.addShort(CommandFieldTag, "US", CEchoRQ)
	request.addShort(MessageIDTag, "US", nextMessageID())
	request.addShort(CommandDataSetTypeTag, "US", NoDataSetCode)

	// return the request
	return request, nil
}

// newCEchoResponse constructs a C-Echo response message based on the C-Echo request message
func newCEchoResponse(assoc *Assoc, request *Object) (*Object, error) {

	// use the message id from the request as the message id responded to
	messageID, err := request.asShort(MessageIDTag, 0)
	if err != nil {
		return nil, err
	}

	// construct a response
	response := newObject()
	response.addUID(AffectedSOPClassUIDTag, VerificationUID)
	response.addShort(CommandFieldTag, "US", CEchoRSP)
	response.addShort(MessageIDBeingRespondedToTag, "US", messageID)
	response.addShort(CommandDataSetTypeTag, "US", NoDataSetCode)
	response.addShort(StatusTag, "US", SuccessStatusCode)

	// return the response
	return response, nil
}

// newCStoreRequest constructs a C-Echo request message
func newCStoreRequest(assoc *Assoc, sopClassUID string, sopInstanceUID string) (*Object, error) {

	// generate a message id
	messageID := nextMessageID()

	// construct a request
	request := newObject()
	request.addUID(AffectedSOPClassUIDTag, sopClassUID)
	request.addShort(CommandFieldTag, "US", CStoreRQ)
	request.addShort(MessageIDTag, "US", messageID)
	request.addShort(PriorityTag, "US", MediumPriorityCode)
	request.addShort(CommandDataSetTypeTag, "US", DataSetCode)
	request.addUID(AffectedSOPInstanceUIDTag, sopInstanceUID)

	// return the request
	return request, nil
}

// newCStoreResponse constructs a C-Store response message based on the C-Store request message
func newCStoreResponse(assoc *Assoc, request *Object) (*Object, error) {

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

	// use the affected sop instance uid from the request in the response
	affectedSOPInstanceUID, err := request.asString(AffectedSOPInstanceUIDTag, 0)
	if err != nil {
		return nil, err
	}

	// construct a response
	response := newObject()
	response.addUID(AffectedSOPClassUIDTag, affectedSOPClassUID)
	response.addShort(CommandFieldTag, "US", CStoreRSP)
	response.addShort(MessageIDBeingRespondedToTag, "US", messageID)
	response.addShort(CommandDataSetTypeTag, "US", NoDataSetCode)
	response.addShort(StatusTag, "US", SuccessStatusCode)
	response.addUID(AffectedSOPInstanceUIDTag, affectedSOPInstanceUID)

	// return the response
	return response, nil
}
