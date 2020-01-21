package dcm4go

// CommandHandler is an interface for handling commands
type CommandHandler interface {
	HandleCommand(assoc *Assoc, pcID byte, command *Object, pDataReader *PDataReader) (*Object, error)
}

// RequestHandler is an interface for handling requests
type RequestHandler interface {
	HandleRequest(assoc *Assoc, request *Message) (*Message, error)
}
