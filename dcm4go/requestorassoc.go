package dcm4go

import (
	"fmt"
	"io"
	"net"
)

// RequestorAssoc is an association used by requestors of associations
type RequestorAssoc struct {
	Assoc
}

// RequestAssoc requests an association
func RequestAssoc(conn net.Conn, local *AE, remote *AE, capabilities []*Capability) (*RequestorAssoc, error) {

	// put together an association request pdu
	assocRQPDU := newAssocRQPDU(remote.AETitle, local.AETitle, capabilities)
	fmt.Printf("assocRQPDU is %v", assocRQPDU)

	// write the pdu
	if err := writeAssocRQPDU(conn, assocRQPDU); err != nil {
		return nil, err
	}

	// read the response
	pdu, err := readPDU(conn)
	if err != nil {
		return nil, err
	}
	fmt.Printf("pdu is %v\n", pdu)

	// is this an association accept?
	if pdu.pduType == aAssociateACPDU {
		assocACPDU, err := readAssocACPDU(pdu)
		if err != nil {
			return nil, err
		}
		fmt.Printf("assocACPDU is %v\n", assocACPDU)

		// create an association from the response
		assoc := &RequestorAssoc{
			Assoc{
				conn:       conn,
				ae:         local,
				assocRQPDU: assocRQPDU,
				assocACPDU: assocACPDU,
			},
		}

		// return assoc
		return assoc, nil
	}

	if pdu.pduType == aAssociateRJPDU {
		fmt.Printf("received a rejection\n")
		assocRJPDU, err := readAssocRJPDU(pdu)
		if err != nil {
			return nil, err
		}
		fmt.Printf("assocRJPDU is %v\n", assocRJPDU)

		return nil, fmt.Errorf("associate request rejected, %s", assocRJPDU)
	}

	// is this an abort request?  if so, just return EOF
	if pdu.pduType == aAbortPDU {
		fmt.Printf("received an abort\n")
		return nil, io.EOF
	}

	return nil, fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
}

// RequestRelease requests release from an association
func (assoc *RequestorAssoc) RequestRelease() error {

	// write a request release pdu
	if err := writeReleaseRQPDU(assoc.conn); err != nil {
		return err
	}
	fmt.Printf("wrote a release request\n")

	// read the response
	pdu, err := readPDU(assoc.conn)
	if err != nil {
		return err
	}
	fmt.Printf("pdu is %v\n", pdu)

	if pdu.pduType == aReleaseRPPDU {
		fmt.Printf("received a release response\n")
		if err := readReleaseRPPDU(pdu); err != nil {
			return err
		}

		// all is well
		return nil
	}

	// is this an abort request?  if so, just return EOF
	if pdu.pduType == aAbortPDU {
		fmt.Printf("received an abort\n")
		// all is well
		return nil
	}

	return fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
}

// Close closes the connection
func (assoc *RequestorAssoc) Close() error {
	if assoc.conn != nil {
		if err := assoc.conn.Close(); err != nil {
			return err
		}
		assoc.conn = nil
	}
	return nil
}

// Verify sends a verification request
func (assoc *RequestorAssoc) Verify() error {

	// create a verification request
	pc, request, err := NewCEchoRequest(&assoc.Assoc)
	if err != nil {
		return err
	}
	fmt.Printf("c-echo request is %v\n", request)

	// write the verification request
	if err := assoc.writeCommand(pc, request); err != nil {
		return err
	}

	// read the response
	response, err := assoc.ReadResponse()
	if err != nil {
		return err
	}
	fmt.Printf("c-echo response is %v\n", response)

	// get the status
	status, err := response.command.asShort(StatusTag, 0)
	if err != nil {
		return err
	}

	if status != 0 {
		return fmt.Errorf("status was %d, not success", status)
	}

	return nil
}

// Send sends a store request
func (assoc *RequestorAssoc) Send(reader io.Reader) error {

	// read the group two attributes
	groupTwo, err := ReadGroupTwo(reader, 0)
	if err != nil {
		return err
	}

	// get the sop class uid of the stored object
	sopClassUID, err := groupTwo.AsString(MediaStorageSOPClassUIDTag, 0)
	if err != nil {
		return err
	}
	fmt.Printf("sop class uid is %q\n", sopClassUID)

	// get the sop class instance UID of the stored object
	sopInstanceUID, err := groupTwo.AsString(MediaStorageSOPInstanceUIDTag, 0)
	if err != nil {
		return err
	}
	fmt.Printf("sop class uid is %q\n", sopClassUID)

	// get the transfer syntax used to store the file
	transferSyntaxUID, err := groupTwo.AsString(TransferSyntaxUIDTag, 0)
	if err != nil {
		return err
	}
	fmt.Printf("transfer syntax uid is %q\n", transferSyntaxUID)

	// create a group zero object
	pc, request, err := NewCStoreRequest(&assoc.Assoc, sopClassUID, sopInstanceUID, transferSyntaxUID)
	if err != nil {
		return err
	}
	fmt.Printf("request is %v\n", request)

	// write the request, but no data
	if err := assoc.writeCommand(pc, request); err != nil {
		return err
	}

	// create a pdatawriter to copy the data to
	// it knows how to create pdus and pdvs as required
	// since it implements a writer, we can then simply copy the data
	pDataWriter := newPDataWriter(
		assoc.conn,                               // the writer is the association connection
		pc.id,                                    // write using the same presentation context id as in the request
		false,                                    // false means we are writing data
		assoc.assocRQPDU.userInfo.maxLenReceived, // the max length of each PDU written
	)

	// copy the data
	num, err := io.Copy(pDataWriter, reader)
	if err != nil {
		return err
	}
	fmt.Printf("copied %d bytes\n", num)

	// flush the underlying writer
	// passing true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	// read the response
	response, err := assoc.ReadResponse()
	if err != nil {
		return err
	}
	fmt.Printf("cstore response is %v\n", response)

	// get the status
	status, err := response.command.asShort(StatusTag, 0)
	if err != nil {
		return err
	}

	if status != 0 {
		return fmt.Errorf("status was %d, not success", status)
	}

	// otherwise, all is well
	return nil
}

// ReadResponse reads a response from the association
func (assoc *RequestorAssoc) ReadResponse() (*Message, error) {
	return assoc.ReadRequestOrResponse()
}

// WriteRequest writes a request to the association
func (assoc *RequestorAssoc) WriteRequest(message *Message) error {
	return assoc.WriteRequestOrResponse(message)
}
