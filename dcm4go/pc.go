// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// PresContext represents a presentation context
type PresContext struct {
	id               byte
	abstractSyntax   string
	transferSyntaxes []string
	result           byte
}

// ID returns the id
func (presContext *PresContext) ID() byte {
	return presContext.id
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

func readPresContext(reader io.Reader, itemType byte) (*PresContext, error) {

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

	if itemType == rqPresContextItemType {
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
	presContext := &PresContext{
		id,               // the pc id
		abstractSyntax,   // the abstract syntax
		transferSyntaxes, // the transfer syntaxes
		reason,           // the reason
	}

	// return the presentation context
	return presContext, nil
}

func writePresContexts(writer io.Writer, presContexts []*PresContext, itemType byte) error {

	// for each of the  presentation contexts
	for _, presContext := range presContexts {

		// write it
		if err := writePresContext(writer, presContext, itemType); err != nil {
			return err
		}
	}

	// all is well
	return nil
}

func writePresContext(writer io.Writer, presContext *PresContext, itemType byte) error {

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
	if err := writeByte(byteWriter, presContext.id); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(byteWriter, 0x00); err != nil {
		return err
	}

	// write the result if accepted presentation context, otherwise write a zero
	if itemType == acPresContextItemType {
		if err := writeByte(byteWriter, presContext.result); err != nil {
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
	if itemType == rqPresContextItemType {
		if err := writeAbstractSyntax(byteWriter, presContext.abstractSyntax); err != nil {
			return err
		}
	}

	// write the transfer syntaxes, works for both types of presentation contexts
	// requested presentation contexts can have multiple transfer syntaxes
	// accepted presentation contexts should only have one
	for _, transferSyntax := range presContext.transferSyntaxes {
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
