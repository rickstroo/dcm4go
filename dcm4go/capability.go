package dcm4go

import "fmt"

// A Capability describes the ability of an AE to send or receive data
// It consists of an abstract syntax (i.e. the type of object to send or receive)
// and a set of transfer sytnaxes (i.e. the possible encodings for that object)
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
