package dcm4go

import "net"

// A Requestor is used to negotiate an associate request, to issue requests
// on the resulting association, and eventually, to release the association.
type Requestor struct {
	conn  net.Conn
	assoc *Assoc
}

// RequestAssoc2 is used to send an associate request on a connection that
// has already been established.
func RequestAssoc2(conn net.Conn) (*Requestor, error) {

	requestor := &Requestor{
		conn: conn,
	}

	return requestor, nil
}

// SendRequest is used to send a request and receive responses.
func (requestor *Requestor) SendRequest(request *Request) ([]*Response, error) {
	return nil, nil
}

// ReleaseAssoc releases the association
func (requestor *Requestor) ReleaseAssoc() error {
	return nil
}

// Abort terminates the association.
func (requestor *Requestor) Abort() error {
	return nil
}
