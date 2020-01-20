package dcm4go

// RequestHandler is an interface for handling requests
type RequestHandler interface {
	HandleRequest(assoc *Assoc, request *Message) (*Message, error)
}
