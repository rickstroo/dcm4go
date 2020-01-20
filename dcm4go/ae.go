package dcm4go

import "fmt"

// AE represents an application entity
type AE struct {
	aeTitle      string
	presContexts []*SPPresContext
	handlers     map[string]RequestHandler
}

// NewAE creates a new application entity
func NewAE(aeTitle string) *AE {
	return &AE{aeTitle, make([]*SPPresContext, 0, 5), make(map[string]RequestHandler)}
}

// String returns a string representation of an ae
func (ae *AE) String() string {
	return fmt.Sprintf(
		"{aeTitle:%q,presContexts:%s,handlers:%v",
		ae.aeTitle,
		ae.presContexts,
		ae.handlers)
}

// AddSupportedPresentationContext adds a presentation context that is supported by this AE
func (ae *AE) AddSupportedPresentationContext(abstractSyntax string, transferSyntaxes []string, handler RequestHandler) {
	presContext := &SPPresContext{abstractSyntax, transferSyntaxes}
	ae.presContexts = append(ae.presContexts, presContext)
	ae.handlers[abstractSyntax] = handler
}
