package dcm4go

import "fmt"

// Capability describes the ability of an AE to send or receive data
// It consists of an abstract syntax (i.e. the type of object to send or receive)
// and a set of transfer sytnaxes (i.e. the possible encodings for that object)
// Capabilities are negotiated between senders and receivers.
type Capability struct {
	abstractSyntax   string
	transferSyntaxes []string
}

func (capability *Capability) String() string {
	return fmt.Sprintf(
		"{abstractSyntax:%q,transferSyntaxes:%q}",
		capability.abstractSyntax,
		capability.transferSyntaxes,
	)
}
