// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

// A Capability represents the ability to transfer or service information.
// It consists of an Abstract Syntax (often referred to as the
// SOP Class UID, which represents the type of the object that is being
//  managed) and a set of Transfer Syntaxes (which represents the supported
// encodings of the object that is being managed.)
type Capability struct {
	abstractSyntax   string
	transferSyntaxes []string
}

// NewCapability creates and initializes a new capability
func NewCapability(abstractSyntax string, transferSyntaxes []string) *Capability {
	capability := &Capability{
		abstractSyntax:   abstractSyntax,
		transferSyntaxes: transferSyntaxes,
	}
	return capability
}

// equal returns true if two capabilities are equal.
func (capability *Capability) equal(otherCapability *Capability) bool {
	if capability.abstractSyntax != otherCapability.abstractSyntax {
		return false
	}
	if len(capability.transferSyntaxes) != len(otherCapability.transferSyntaxes) {
		return false
	}
	for _, transferSyntax := range capability.transferSyntaxes {
		if !otherCapability.containsTransferSyntax(transferSyntax) {
			return false
		}
	}
	return true
}

// containsTransferSyntax returns true if the capability includes
// the specified transfer syntax.
func (capability *Capability) containsTransferSyntax(otherTransferSyntax string) bool {
	for _, transferSyntax := range capability.transferSyntaxes {
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
func (capabilities *Capabilities) contains(otherCapability *Capability) bool {
	for _, capability := range capabilities.capabilities {
		if capability.equal(otherCapability) {
			return true
		}
	}
	return false
}

// Add adds a capability to the set of capabilities
func (capabilities *Capabilities) Add(capability *Capability) {
	if !capabilities.contains(capability) {
		capabilities.capabilities = append(capabilities.capabilities, capability)
	}
}
