package dcm4go

import (
	"bytes"
	"fmt"
)

// NewCEchoResponse constructs a C-Echo response message based on the C-Echo request message
func NewCEchoResponse(assoc *Assoc, request *Message) (*Message, error) {

	// use the same pc id as the request
	pcID := request.pcID

	// use the message id from the request as the message id responded to
	messageID, err := request.Command().asShort(MessageIDTag, 0)
	if err != nil {
		return nil, err
	}

	// create a temporary object for what we know
	temp := newObject()
	temp.addUID(AffectedSOPClassUIDTag, VerificationUID)
	temp.addShort(CommandFieldTag, "US", CEchoRSP)
	temp.addShort(MessageIDBeingRespondedToTag, "US", messageID)
	temp.addShort(CommandDataSetTypeTag, "US", 0x0101)
	temp.addShort(StatusTag, "US", 0x00)

	// create a buffer to write the temporary object to
	buf := new(bytes.Buffer)

	// create an encoder for writing objects
	encoder := newEncoder()

	// find the transfer syntax for commands, always implicit VR little endian
	transferSyntax := ImplicitVRLittleEndianTS()
	fmt.Printf("transfer syntax is %v\n", transferSyntax)

	// write the temporary to the buffer
	if err := encoder.writeObject(buf, temp, transferSyntax.explicitVR, transferSyntax.byteOrder); err != nil {
		return nil, err
	}

	// now create the final command object
	command := newObject()

	// initialize it with the command group length attribute
	command.addLong(CommandGroupLengthTag, "UL", uint32(buf.Len()))

	// add the rest of the attributes from the temporary object
	command.addAll(temp)

	// construct and return a message
	return &Message{pcID, command, nil}, nil
}
