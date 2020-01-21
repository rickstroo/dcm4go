package dcm4go

import "fmt"

// AE represents an application entity
type AE struct {
	aeTitle         string
	presContexts    []*SPPresContext
	commandHandlers map[string]CommandHandler
	requestHandlers map[string]RequestHandler
}

// NewAE creates a new application entity
func NewAE(aeTitle string) *AE {
	return &AE{aeTitle, make([]*SPPresContext, 0, 5), make(map[string]CommandHandler), make(map[string]RequestHandler)}
}

// String returns a string representation of an ae
func (ae *AE) String() string {
	return fmt.Sprintf(
		"{aeTitle:%q,presContexts:%s,commandHandlers:%v,requestHandlers:%v}",
		ae.aeTitle,
		ae.presContexts,
		ae.commandHandlers,
		ae.requestHandlers)
}

// AddSupportedPresentationContext adds a presentation context that is supported by this AE
func (ae *AE) AddSupportedPresentationContext(abstractSyntax string, transferSyntaxes []string, commandHandler CommandHandler, requestHandler RequestHandler) {
	presContext := &SPPresContext{abstractSyntax, transferSyntaxes}
	ae.presContexts = append(ae.presContexts, presContext)
	ae.commandHandlers[abstractSyntax] = commandHandler
	ae.requestHandlers[abstractSyntax] = requestHandler
}
