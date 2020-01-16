package dcm4go

import "fmt"

// AE represents an application entity
type AE struct {
	aeTitle      string
	presContexts []*SPPresContext
}

// NewAE creates a new application entity
func NewAE(aeTitle string) *AE {
	return &AE{aeTitle, make([]*SPPresContext, 0, 5)}
}

// String returns a string representation of an ae
func (ae *AE) String() string {
	return fmt.Sprintf(
		"{aeTitle:%q,presContexts:%s",
		ae.aeTitle,
		ae.presContexts)
}

// AddSupportedPresentationContext adds a presentation context that is supported by this AE
func (ae *AE) AddSupportedPresentationContext(abstractSyntax string, transferSyntaxes []string) {
	presContext := &SPPresContext{abstractSyntax, transferSyntaxes}
	ae.presContexts = append(ae.presContexts, presContext)
}
