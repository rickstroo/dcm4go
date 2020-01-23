package dcm4go

import (
	"fmt"
	"strings"
)

// AssocACRQPDU presents an Association Accept and an Association Request PDU
type AssocACRQPDU struct {
	protocol       uint16
	calledAETitle  string
	callingAETitle string
	appContextName string
	presContexts   []*PresContext
	userInfo       *UserInfo
}

// String returns a string representation of a AssocACRQPDU
func (pdu *AssocACRQPDU) String() string {
	return fmt.Sprintf(
		"{protocol:%v,calledAET:%q,callingAET:%q,appContextName:%q,presContexts:%s,userInfo:%s}",
		pdu.protocol,
		strings.TrimSpace(pdu.calledAETitle),
		strings.TrimSpace(pdu.callingAETitle),
		pdu.appContextName,
		pdu.presContexts,
		pdu.userInfo)
}
