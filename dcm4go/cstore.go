package dcm4go

// NewCStoreResponse constructs a C-Store response message based on the C-Store request message
func NewCStoreResponse(assoc *Assoc, request *Object) (*Object, error) {
	return newCStoreResponse(assoc, request)
}

// NewCStoreRequest constructs a C-Echo request message
func NewCStoreRequest(assoc *Assoc, sopClassUID string, sopInstanceUID string, transferSyntaxUID string) (*PresContext, *Object, error) {

	// find the presentation context that was negotiated for this sop class
	pc, err := assoc.findAcceptedPresContextByAbstractSyntax(sopClassUID)
	if err != nil {
		return nil, nil, err
	}

	// construct a request
	request, err := newCStoreRequest(assoc, sopClassUID, sopInstanceUID)
	if err != nil {
		return nil, nil, err
	}

	// construct and return a message
	return pc, request, nil
}
