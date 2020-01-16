package dcm4go

import (
	"fmt"
)

// SPPresContext defines a supported presentation context
type SPPresContext struct {
	abstractSyntax   string
	transferSyntaxes []string
}

// String returns a string representation of a supported presentation context
func (presContext *SPPresContext) String() string {
	return fmt.Sprintf(
		"{abstractSyntax:%q,transferSyntaxes:%q",
		presContext.abstractSyntax,
		presContext.transferSyntaxes)
}

// RQPresContext represents a requested presentation context
type RQPresContext struct {
	id               byte
	abstractSyntax   string
	transferSyntaxes []string
}

// String returns a string representation of a requested presentation context
func (presContext *RQPresContext) String() string {
	return fmt.Sprintf(
		"{id:%d,abstractSyntax:%q,transferSyntaxes:%q",
		presContext.id,
		presContext.abstractSyntax,
		presContext.transferSyntaxes)
}

// ACPresContext represents an accepted presentation context
type ACPresContext struct {
	id             byte
	result         byte
	transferSyntax string
}

// String returns a string representation of an accepted presentation context
func (presContext *ACPresContext) String() string {
	return fmt.Sprintf("{id:%d,result:%d,transferSyntax:%q}",
		presContext.id,
		presContext.result,
		presContext.transferSyntax)
}
