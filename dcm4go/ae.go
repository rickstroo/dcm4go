package dcm4go

import "fmt"

// AE represents an application entity
type AE struct {
	aeTitle         string
	presContexts    []*PresContext
	commandHandlers map[string]CommandHandler
}

// NewAE creates a new application entity
func NewAE(aeTitle string) *AE {
	return &AE{aeTitle, make([]*PresContext, 0, 5), make(map[string]CommandHandler)}
}

// String returns a string representation of an ae
func (ae *AE) String() string {
	return fmt.Sprintf(
		"{aeTitle:%q,presContexts:%s,commandHandlers:%v}",
		ae.aeTitle,
		ae.presContexts,
		ae.commandHandlers)
}

// AddSupportedPresentationContext adds a presentation context that is supported by this AE
func (ae *AE) AddSupportedPresentationContext(abstractSyntax string, transferSyntaxes []string, commandHandler CommandHandler) {
	presContext := &PresContext{0, abstractSyntax, transferSyntaxes, 0}
	ae.presContexts = append(ae.presContexts, presContext)
	ae.commandHandlers[abstractSyntax] = commandHandler
}

// AddRequestedPresentationContext adds a presentation context that is requested by this AE
func (ae *AE) AddRequestedPresentationContext(abstractSyntax string, transferSyntaxes []string) {
	presContext := &PresContext{0, abstractSyntax, transferSyntaxes, 0}
	ae.presContexts = append(ae.presContexts, presContext)
}
