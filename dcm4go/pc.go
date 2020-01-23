package dcm4go

import (
	"fmt"
)

// PresContext represents a presentation context
type PresContext struct {
	id               byte
	abstractSyntax   string
	transferSyntaxes []string
	result           byte
}

// String returns a string representation of a requested presentation context
func (presContext *PresContext) String() string {
	return fmt.Sprintf(
		"{id:%d,abstractSyntax:%q,transferSyntaxes:%q,result:%d}",
		presContext.id,
		presContext.abstractSyntax,
		presContext.transferSyntaxes,
		presContext.result)
}
