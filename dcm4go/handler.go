package dcm4go

// CommandHandler is an interface for handling commands
type CommandHandler interface {
	HandleCommand(assoc *Assoc, pcID byte, command *Object, pDataReader *PDataReader) (*Object, error)
}
