package dcm4go

import "io"

// I'm starting to think that having an AE as the starting point for all the
// APIs in the library makes sense.  It is AEs that application developers
// are actually building.  The APIs below allow for AEs to serve as SCUs,
// SCPs, and even SCUs and SCPs concurrently.

// AE represents an application entity
type AE struct {
	AETitle string
	Host    string
	Port    int
}

// Should RequestAssoc create the underlying network connections as well?
// From a practical point of view, it may be nice to save the user from having
// to manage that themselves.  From a separation of concerns point of view,
// it might be nice to handle that logic at another layer.  However, that
// means adding additional APIs to the library.

// Still need to think about how best to encapsulate the low level information
// about the transfer capabilities that get turned into presentation contexts
// that are negotiated.  It always seems odd to separate the management of
// the UIDs specifying capabilities from the functions themselves.  That's
// why I was thinking of creating some objects that would encapsulate both.
// The might look something like the following.

// A Context describes the way in which information is stored or transferred
// between AEs.  A Context is made of up an abstract syntax (i.e. the type
// of the information) and a set of transfer syntaxes (i.e. the encoding of
// the information).  Why a set of transfer syntaxes?  Obviously, when
// information is stored or transferred, only one transfer syntax is used.
// However, a Context is also used to describe the capabilities of a service
// to store or transmit information, so, a set is provided to describe those
// capabilities.
type Context struct {
	AbstractSyntax   string
	TransferSyntaxes []string
}

// Transferer is the interface that wraps the basic Transfer method.
// It also wraps the Capabilities method that describes the transfer
// capabilities of the transferer.
type Transferer interface {
	Capabilities() []*Capability
	Transfer(*Assoc, *Request) (*Response, error)
}

// An Service defines a service that can be provided or used.
type Service struct {
	Contexts []*Context
}

// A CEchoService defines the C-Echo service
type CEchoService struct {
	Service
}

// Echo sends a C-Echo request to the associated AE
func (service *CEchoService) Echo() error {
	return nil
}

// OnEcho handles a C-Echo request from an associated AE.
// The handler is passed a reference to the association that the request
// was received on.  The handler is also passed a reference to the C-Echo
// request.  It is expected that the handler will compose an appropriate
// C-Echo response and write it back to the association.
func (service *CEchoService) OnEcho(assoc *Assoc, request *Request) error {
	return nil
}

// A CStoreService defines the C-Store service
type CStoreService struct {
	Service
}

// Store sends a C-Store request to the associated ae AE.
func (service *CStoreService) Store(path string) error {
	return nil
}

// OnStore handles a C-Store request from an associated AE
func (service *CStoreService) OnStore(assoc *Assoc, request *Request, reader io.Reader) error {
	return nil
}

// RequestAssoc sends an associate request with the other AE.
// A set of services that this AE expects to user on this association
// are provided, so that the appropriate transfer capabilities can be
// be negotiated.
func (ae *AE) RequestAssoc(other *AE, services ...*Service) (*Assoc, error) {
	return nil, nil
}

// Similar question for AcceptAssoc.  Should it bind to an address and only
// return when an association is accepted?  One might want to have a separate
// thread created immediately when the connection is accepted.  How might
// that work?  Would there be some state in the AE itself to manage this?
// I guess it should bind to the address associated with the ae.

// AcceptAssoc receives and negotiates an associate request
func (ae *AE) AcceptAssoc(services ...*Service) (*Assoc, error) {
	return nil, nil
}

// Release releases an association
func (assoc *Assoc) Release() error {
	return nil
}

// Abort aborts an association
func (assoc *Assoc) Abort() error {
	return nil
}

// Hmm, I think I'm starting to like this much simpler API.
