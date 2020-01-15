package dcm4go

import (
	"fmt"
	"io"
)

// AssocACPDU presents an Association Accept PDU
type AssocACPDU struct {
	protocol       uint16
	calledAETitle  string
	callingAETitle string
	appContextName string
	presContexts   []*ACPresContext
	userInfo       *UserInfo
}

// ACPresContext represents an accepted presentation context
type ACPresContext struct {
	id             byte
	transferSyntax string
}

// String returns a string representation of a UserInfo
func (presContext *ACPresContext) String() string {
	return fmt.Sprintf("{id:%d,transferSyntax:%q}", presContext.id, presContext.transferSyntax)
}

func writeAssocACPDU(reader io.Reader, assocACPDU *AssocACPDU) error {
	return fmt.Errorf("writeAssocACPDU: not implemented")
}
