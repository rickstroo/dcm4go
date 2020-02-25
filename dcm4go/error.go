package dcm4go

import "errors"

// ErrIllegalPrefix means that a prefix other than "DICM" was encountered
// at the beginning of the file
var ErrIllegalPrefix = errors.New("illegal prefix")

// ErrUnexpectedAttribute means that an attribute was not expected
// when decoding a DICOM strea
var ErrUnexpectedAttribute = errors.New("unexpected attribute")

// ErrUnrecognizedVR means that the VR was not known
var ErrUnrecognizedVR = errors.New("unrecognized VR")

// ErrIndexOutOfBounds means that an index was out of bounds when
// trying to read the value of an attribute
var ErrIndexOutOfBounds = errors.New("index out of bounds")

// ErrWrongType means that the value was not of the type that was
// expected when trying to read the value of an attribute
var ErrWrongType = errors.New("wrong type")

// ErrAttributeNotFound means that the attribute was not found in the
// object that we were searching.  It might be argued that this is not
// an error, and that we should simply return a boolean.
var ErrAttributeNotFound = errors.New("attribute not found")

// ErrUnrecognizedTransferSyntax means that the transfer syntax
// was not unrecognized
var ErrUnrecognizedTransferSyntax = errors.New("transfer syntax not recognized")

// ErrUnexpectedPDU means that the pdu type was unexpected
var ErrUnexpectedPDU = errors.New("unexpected PDU")

// ErrAssociateRequestRejected means that the associate request was rejected
var ErrAssociateRequestRejected = errors.New("associate request rejected")

// ErrNoConn means that there is no connection between the client and server
var ErrNoConn = errors.New("no connection")

// ErrNoAssoc means that there is no association between the client and the server
var ErrNoAssoc = errors.New("no association")

// ErrUnrecognizedCallingAETitle is returned to client when the calling AE
// Title is not recognized
var ErrUnrecognizedCallingAETitle = errors.New("unrecognized calling AE Title")

// ErrUnrecognizedCalledAETitle is returned to client when the called AE
// Title is not recognized
var ErrUnrecognizedCalledAETitle = errors.New("unrecognized called AE Title")

// ErrUnexpectedRequest is returned when the request is not expected
var ErrUnexpectedRequest = errors.New("unexpected request")

// ErrAssociationAborted is returned when the association is aborted
var ErrAssociationAborted = errors.New("association aborted")
