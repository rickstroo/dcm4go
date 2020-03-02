package dcm4go

// type releaseRQPDU struct{}
//
// // readReleaseRQPDU reads an ReleaseRQPDU from a reader
// func readReleaseRQPDU(reader io.Reader) (*releaseRQPDU, error) {
//
// 	pdu := &releaseRQPDU{}
//
// 	// read and ignore the long
// 	if _, err := readLong(reader, binary.BigEndian); err != nil {
// 		return nil, err
// 	}
//
// 	return pdu, nil
// }
//
// // writeReleaseRQPDU writes a release request PDU to a writer
// func writeReleaseRQPDU(writer io.Writer) error {
//
// 	// construct a pdu
// 	pdu := &pdu{typ: aReleaseRQPDU}
//
// 	// write a long zero
// 	if err := writeLong(pdu, 0x00, binary.BigEndian); err != nil {
// 		return err
// 	}
//
// 	// write the pdu header
// 	if err := writePDU(writer, pdu); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
