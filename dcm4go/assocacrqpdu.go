package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// item types and sub item types
const (
	appContextItemType      = 0x10
	rqPresContextItemType   = 0x20
	acPresContextItemType   = 0x21
	abstractSyntaxItemType  = 0x30
	transferSyntaxItemType  = 0x40
	userInfoItemType        = 0x50
	maxLengthItemType       = 0x51
	implClassUIDItemType    = 0x52
	maxNumOpsItemType       = 0x53
	implVersionNameItemType = 0x55
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

func writeVariableItems(writer io.Writer, appContextName string, presContexts []*PresContext, itemType byte, userInfo *UserInfo) error {

	if err := writeAppContextName(writer, appContextName); err != nil {
		return err
	}

	if err := writePresContexts(writer, presContexts, itemType); err != nil {
		return err
	}

	if err := writeUserInfo(writer, userInfo); err != nil {
		return err
	}

	return nil
}

func writeAppContextName(writer io.Writer, appContextName string) error {

	// write item type
	if err := writeByte(writer, 0x10); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}
	// write the length of the application context name
	if err := writeShort(writer, uint16(len(appContextName)), binary.BigEndian); err != nil {
		return err
	}

	// write the application context name
	if err := writeString(writer, appContextName); err != nil {
		return err
	}

	// all is well
	return nil
}
