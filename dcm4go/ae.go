package dcm4go

import "fmt"

// AE represents an application entity
type AE struct {
	aeTitle         string
	capabilities    []*Capability
	commandHandlers map[string]CommandHandler
}

// NewAE creates a new application entity
func NewAE(aeTitle string) *AE {
	return &AE{aeTitle, make([]*Capability, 0, 5), make(map[string]CommandHandler)}
}

// String returns a string representation of an ae
func (ae *AE) String() string {
	return fmt.Sprintf(
		"{aeTitle:%q,capabilities:%v,commandHandlers:%v}",
		ae.aeTitle,
		ae.capabilities,
		ae.commandHandlers)
}

// AETitle returns the AE title of the AE
func (ae *AE) AETitle() string {
	return ae.aeTitle
}

// AddSupportedCapability adds a capability that is supported by this AE
func (ae *AE) AddSupportedCapability(abstractSyntax string, transferSyntaxes []string, commandHandler CommandHandler) {
	capability := &Capability{abstractSyntax, transferSyntaxes}
	ae.capabilities = append(ae.capabilities, capability)
	ae.commandHandlers[abstractSyntax] = commandHandler
}

// AddRequestedCapability adds a capability that is supported by this AE
func (ae *AE) AddRequestedCapability(abstractSyntax string, transferSyntaxes []string) {
	capability := &Capability{abstractSyntax, transferSyntaxes}
	ae.capabilities = append(ae.capabilities, capability)
}
