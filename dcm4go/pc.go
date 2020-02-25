// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// PC represents a presentation context
type PC struct {
	ID               byte
	AbstractSyntax   string
	TransferSyntaxes []string
	Result           byte
}

// String returns a string representation of a requested presentation context
func (pc *PC) String() string {
	return fmt.Sprintf(
		"{id:%d,abstractSyntax:%q,transferSyntaxes:%q,result:%d}",
		pc.ID,
		pc.AbstractSyntax,
		pc.TransferSyntaxes,
		pc.Result,
	)
}

// Equals returns true if two capabilities have the same abstract syntax
// and all the transfer syntaxes are the same.
func (pc *PC) Equals(other *PC) bool {

	// if the abstract syntaxes are not the same, the capabilities are not the same
	if pc.AbstractSyntax != other.AbstractSyntax {
		return false
	}

	// a quick comparison of the length of the transfer syntaxes can determine
	// that capabilities are not the same
	if len(pc.TransferSyntaxes) != len(other.TransferSyntaxes) {
		return false
	}

	// now, let's see if all the transfer syntaxes from the other are contained
	// in the transfer syntaxes for this capability
	for _, transferSyntax := range other.TransferSyntaxes {
		if !contains(pc.TransferSyntaxes, transferSyntax) {
			return false
		}
	}

	// must be the same
	return true
}

// contains looks for a string in a set of strings
func contains(ses []string, t string) bool {
	for _, s := range ses {
		if s == t {
			return true
		}
	}
	return false
}

// Contained returns true if this capability is contained in a set of other capabilities
func (pc *PC) Contained(others []*PC) bool {

	// see if this capability is equal to any of the others
	for _, other := range others {

		// if it is equal, this capability is contained in the others
		if pc.Equals(other) {
			return true
		}
	}

	// did not find it, so musts be false
	return false
}

func readPresContext(reader io.Reader, itemType byte) (*PC, error) {

	// read the presentation context id
	id, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// skip a byte, as per the standard
	if err := skipByte(reader); err != nil {
		return nil, err
	}

	// initialize the reason
	var reason byte

	if itemType == rqPCItemType {
		// read the reason
		if reason, err = readByte(reader); err != nil {
			return nil, err
		}
	} else {
		// skip a byte, as per the standard
		if err := skipByte(reader); err != nil {
			return nil, err
		}
	}

	// skip a byte, as per the standard
	if err := skipByte(reader); err != nil {
		return nil, err
	}

	// initialize the abstract syntax name
	var abstractSyntax string

	// initialize the list of transfer syntaxes
	transferSyntaxes := make([]string, 0, 5)

	// read the abstract syntax and transfer syntax items
	for {

		// read a sub item
		subItemType, err := readByte(reader)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		// skip a byte, as per the standard
		if err := skipByte(reader); err != nil {
			return nil, err
		}

		// read the length
		length, err := readShort(reader, binary.BigEndian)
		if err != nil {
			return nil, err
		}

		if subItemType == abstractSyntaxItemType { // abstract syntax

			// read the uid
			uid, err := readUID(reader, uint32(length))
			if err != nil {
				return nil, err
			}

			// assign it to the abstract syntax
			abstractSyntax = uid

		} else if subItemType == transferSyntaxItemType { // transfer syntax

			// read the transfer syntax
			uid, err := readUID(reader, uint32(length))
			if err != nil {
				return nil, err
			}

			// add it to the list of transfer syntaxes
			transferSyntaxes = append(transferSyntaxes, uid)

		} else { // unrecgonized item
			return nil, fmt.Errorf("unrecognized presentation context sub item type: 0x%02X", subItemType)
		}
	}

	// create the presentation context
	pc := &PC{
		id,               // the pc id
		abstractSyntax,   // the abstract syntax
		transferSyntaxes, // the transfer syntaxes
		reason,           // the reason
	}

	// return the presentation context
	return pc, nil
}

func writePresContexts(writer io.Writer, pcs []*PC, itemType byte) error {

	// for each of the  presentation contexts
	for _, pc := range pcs {

		// write it
		if err := writePresContext(writer, pc, itemType); err != nil {
			return err
		}
	}

	// all is well
	return nil
}

func writePresContext(writer io.Writer, pc *PC, itemType byte) error {

	// write item type
	if err := writeByte(writer, itemType); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// create a byte array output stream so we can calculate the length of the presentation context
	byteWriter := new(bytes.Buffer)

	// write the presentation context id
	if err := writeByte(byteWriter, pc.ID); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(byteWriter, 0x00); err != nil {
		return err
	}

	// write the result if accepted presentation context, otherwise write a zero
	if itemType == acPCItemType {
		if err := writeByte(byteWriter, pc.Result); err != nil {
			return err
		}
	} else {
		if err := writeByte(byteWriter, 0x00); err != nil {
			return err
		}
	}

	// write a zero as per the standard
	if err := writeByte(byteWriter, 0x00); err != nil {
		return err
	}

	// write the abstract syntax if requested presentation context
	if itemType == rqPCItemType {
		if err := writeAbstractSyntax(byteWriter, pc.AbstractSyntax); err != nil {
			return err
		}
	}

	// write the transfer syntaxes, works for both types of presentation contexts
	// requested presentation contexts can have multiple transfer syntaxes
	// accepted presentation contexts should only have one
	for _, transferSyntax := range pc.TransferSyntaxes {
		if err := writeTransferSyntax(byteWriter, transferSyntax); err != nil {
			return err
		}
	}

	// write the length to the underlying writer
	if err := writeShort(writer, uint16(byteWriter.Len()), binary.BigEndian); err != nil {
		return err
	}

	// write the byte array to the underlying writer
	if err := writeBytes(writer, byteWriter.Bytes()); err != nil {
		return err

	}

	// all is good
	return nil
}

func writeAbstractSyntax(writer io.Writer, abstractSyntax string) error {
	return writeSyntax(writer, abstractSyntax, abstractSyntaxItemType)
}

func writeTransferSyntax(writer io.Writer, transferSyntax string) error {
	return writeSyntax(writer, transferSyntax, transferSyntaxItemType)
}

func writeSyntax(writer io.Writer, syntax string, itemType byte) error {

	// write sub item type
	if err := writeByte(writer, itemType); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// write the length of the syntax
	if err := writeShort(writer, uint16(len(syntax)), binary.BigEndian); err != nil {
		return err
	}

	// write the abstract syntax
	if err := writeString(writer, syntax); err != nil {
		return err
	}

	// all is well
	return nil
}
