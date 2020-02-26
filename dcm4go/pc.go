// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// associate negotiation results for presentation contexts
const (
	pcAcceptance                   = 0x00
	pcUserRejection                = 0x01
	pcNoReason                     = 0x02
	pcAbstractSyntaxNotSupported   = 0x03
	pcTransferSyntaxesNotSupported = 0x04
)

// pc represents a presentation context
type pc struct {
	id               byte
	abstractSyntax   string
	transferSyntaxes []string
	result           byte
}

// String returns a string representation of a requested presentation context
func (pc *pc) String() string {
	return fmt.Sprintf(
		"{id:%d,abstractSyntax:%q,transferSyntaxes:%q,result:%d}",
		pc.id,
		pc.abstractSyntax,
		pc.transferSyntaxes,
		pc.result,
	)
}

func readPC(reader io.Reader, itemType byte) (*pc, error) {

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
	var result byte

	if itemType == rqPCItemType {
		// read the result
		if result, err = readByte(reader); err != nil {
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
	pc := &pc{
		id:               id,               // the pc id
		abstractSyntax:   abstractSyntax,   // the abstract syntax
		transferSyntaxes: transferSyntaxes, // the transfer syntaxes
		result:           result,           // the result
	}

	// return the presentation context
	return pc, nil
}

func writePCs(writer io.Writer, pcs []*pc, itemType byte) error {

	// for each of the  presentation contexts
	for _, pc := range pcs {

		// write it
		if err := writePC(writer, pc, itemType); err != nil {
			return err
		}
	}

	// all is well
	return nil
}

func writePC(writer io.Writer, pc *pc, itemType byte) error {

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
	if err := writeByte(byteWriter, pc.id); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(byteWriter, 0x00); err != nil {
		return err
	}

	// write the result if accepted presentation context, otherwise write a zero
	if itemType == acPCItemType {
		if err := writeByte(byteWriter, pc.result); err != nil {
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
		if err := writeAbstractSyntax(byteWriter, pc.abstractSyntax); err != nil {
			return err
		}
	}

	// write the transfer syntaxes, works for both types of presentation contexts
	// requested presentation contexts can have multiple transfer syntaxes
	// accepted presentation contexts should only have one
	for _, transferSyntax := range pc.transferSyntaxes {
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
