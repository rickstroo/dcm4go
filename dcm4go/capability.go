package dcm4go

import "fmt"

// A Capability describes the ability of an AE to send or receive data
// It consists of an abstract syntax (i.e. the type of object to send or receive)
// and a set of transfer sytnaxes (i.e. the possible encodings for that object)
type Capability struct {
	AbstractSyntax   string
	TransferSyntaxes []string
}

func (capability *Capability) String() string {
	return fmt.Sprintf(
		"{abstractSyntax:%q,transferSyntaxes:%q}",
		capability.AbstractSyntax,
		capability.TransferSyntaxes,
	)
}

// Equals returns true if two capabilities have the same abstract syntax
// and all the transfer syntaxes are the same.
func (capability *Capability) Equals(other *Capability) bool {

	// if the abstract syntaxes are not the same, the capabilities are not the same
	if capability.AbstractSyntax != other.AbstractSyntax {
		return false
	}

	// a quick comparison of the length of the transfer syntaxes can determine
	// that capabilities are not the same
	if len(capability.TransferSyntaxes) != len(other.TransferSyntaxes) {
		return false
	}

	// now, let's see if all the transfer syntaxes from the other are contained
	// in the transfer syntaxes for this capability
	for _, transferSyntax := range other.TransferSyntaxes {
		if !contains(capability.TransferSyntaxes, transferSyntax) {
			return false
		}
	}

	// must be the same
	return true
}

// Contained returns true if this capability is contained in a set of other capabilities
func (capability *Capability) Contained(others []*Capability) bool {

	// see if this capability is equal to any of the others
	for _, other := range others {

		// if it is equal, this capability is contained in the others
		if capability.Equals(other) {
			return true
		}
	}

	// did not find it, so musts be false
	return false
}
