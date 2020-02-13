package dcm4go

import "net"

// An Acceptor is used to negotiate an associate request, to handle the requests
// sent on the resulting association, and finally, to close the association.
type Acceptor struct {
	conn            net.Conn
	requestHandlers []RequestHandler
	assoc           *Assoc
}

// A RequestHandler is used to handle requests.
// Capabilities defines what abstract syntax and transfer syntax combinations
// this handler is capable of handling.
// ReadData returns true if the handler wants the underlying association to
// read the data (if present) and to include it as part of the request.  If
// ReadData returns false, the handler will not read the data, but will include
// a data reader as part of the request.  This allows the user to handler
// large data sets (i.e. images) more efficiently, potentially simply copying
// them from the network system to a file system without having to read them
// entirely into memory.
type RequestHandler interface {
	Capabilities() []*Capability
	ReadData() bool
	HandleRequest(*Acceptor, *Request) error
}

// AcceptAssoc2 attempts to accept an associate request on a connection
// that has already been established.  The handlers includes a set of
// capabilities (e.g. abstract syntaxes and transfer syntaxes) that
// are used to accept or reject the presentation contexts from the associate
// request.  The handlers also include methods that will be called for
// presentation contexts that are accepted.
func AcceptAssoc2(conn net.Conn, requestHandlers []RequestHandler) (*Acceptor, error) {

	acceptor := &Acceptor{
		conn:            conn,
		requestHandlers: requestHandlers,
	}

	return acceptor, nil
}

// HandleRequest reads a request and calls the appropriate handler.
func (acceptor *Acceptor) HandleRequest() error {
	return nil
}

// Abort terminates the association.
func (acceptor *Acceptor) Abort() error {
	return nil
}
