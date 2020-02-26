// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

// A Capability represents the ability to transfer or service information.
// It consists of an Abstract Syntax (often referred to as the
// SOP Class UID, which represents the type of the object that is being
//  managed) and a set of Transfer Syntaxes (which represents the supported
// encodings of the object that is being managed.)
type Capability struct {
	AbstractSyntax   string
	TransferSyntaxes []string
}

// NewCapability creates and initializes a new capability
func NewCapability(abstractSyntax string, transferSyntaxes []string) *Capability {
	capability := &Capability{
		AbstractSyntax:   abstractSyntax,
		TransferSyntaxes: transferSyntaxes,
	}
	return capability
}

// Equal returns true if two capabilities are equal.
func (capability *Capability) Equal(otherCapability *Capability) bool {
	if capability.AbstractSyntax != otherCapability.AbstractSyntax {
		return false
	}
	if len(capability.TransferSyntaxes) != len(otherCapability.TransferSyntaxes) {
		return false
	}
	for _, transferSyntax := range capability.TransferSyntaxes {
		if !otherCapability.ContainsTransferSyntax(transferSyntax) {
			return false
		}
	}
	return true
}

// ContainsTransferSyntax returns true if the capability includes
// the specified transfer syntax.
func (capability *Capability) ContainsTransferSyntax(otherTransferSyntax string) bool {
	for _, transferSyntax := range capability.TransferSyntaxes {
		if transferSyntax != otherTransferSyntax {
			return false
		}
	}
	return true
}

// Capabilities represents a set of capabilities.  As a set, Capabilities
// ensures that at most one instance of a capability is stored.  A capability
// is considered unique if the combination of its Abstract Syntax and
// Transfer Syntaxes (regardless of ordering) is unique.
type Capabilities struct {
	capabilities []*Capability
}

// NewCapabilities creates and initializes a new set of capabilities
func NewCapabilities() *Capabilities {
	capabilities := &Capabilities{
		capabilities: make([]*Capability, 0, 5),
	}
	return capabilities
}

// Contains returns true if the capabilities include the specified capability
func (capabilities *Capabilities) Contains(otherCapability *Capability) bool {
	for _, capability := range capabilities.capabilities {
		if capability.Equal(otherCapability) {
			return true
		}
	}
	return false
}

// Add adds a capability to the set of capabilities
func (capabilities *Capabilities) Add(capability *Capability) {
	if !capabilities.Contains(capability) {
		capabilities.capabilities = append(capabilities.capabilities, capability)
	}
}
