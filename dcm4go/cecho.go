package dcm4go

// NewCEchoResponse constructs a C-Echo response message based on the C-Echo request message
func NewCEchoResponse(assoc *Assoc, request *Object) (*Object, error) {
	return newCEchoResponse(assoc, request)
}

// NewCEchoRequest constructs a C-Echo request message
func NewCEchoRequest(assoc *Assoc) (*PresContext, *Object, error) {

	// find the presentation context id that was negotiated for verification
	pc, err := assoc.findAcceptedPresContextByAbstractSyntax(VerificationUID)
	if err != nil {
		return nil, nil, err
	}

	// construct a command
	request, err := newCEchoRequest(assoc)
	if err != nil {
		return nil, nil, err
	}

	// construct and return a message
	return pc, request, nil
}
