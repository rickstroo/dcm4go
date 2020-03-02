// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

// const (
// 	aAssociateRQPDUType = 0x01
// 	aAssociateACPDUType = 0x02
// 	aAssociateRJPDUType = 0x03
// 	pDataTFPDUType      = 0x04
// 	aReleaseRQPDUType   = 0x05
// 	aReleaseRPPDUType   = 0x06
// 	aAbortPDUType       = 0x07
// )
//
// type pdu0 struct {
// 	typ byte
// 	len uint32
// }
//
// func newPDU(typ byte, len uint32) *pdu0 {
// 	return &pdu0{typ: typ, len: len}
// }
//
// func readPDU0(reader io.Reader) (*pdu0, error) {
//
// 	var buf [6]byte
// 	if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 		return nil, err
// 	}
//
// 	pdu := &pdu0{}
//
// 	pdu.typ = buf[0]
// 	pdu.len = binary.BigEndian.Uint32(buf[2:6])
//
// 	return pdu, nil
// }
//
// func (pdu *pdu0) write(writer io.Writer) error {
//
// 	var buf [6]byte
// 	buf[0] = pdu.typ
// 	buf[1] = 0x00
// 	binary.BigEndian.PutUint32(buf[2:6], pdu.len)
//
// 	if _, err := writer.Write(buf[:]); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// type abortPDU0 struct {
// 	source byte // the initiator of the abort
// 	reason byte // the reason for the abort
// }
//
// func readAbortPDU0(reader io.Reader) (*abortPDU0, error) {
//
// 	var buf [4]byte
// 	if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 		return nil, err
// 	}
//
// 	pdu := &abortPDU0{}
//
// 	pdu.reason = buf[2]
// 	pdu.source = buf[3]
//
// 	return pdu, nil
// }
//
// func (pdu *abortPDU0) write(writer io.Writer) error {
//
// 	var buf [4]byte
// 	buf[0] = 0x00
// 	buf[1] = 0x00
// 	buf[2] = pdu.reason
// 	buf[3] = pdu.source
//
// 	if err := newPDU(aAbortPDUType, uint32(len(buf))).write(writer); err != nil {
// 		return err
// 	}
//
// 	if _, err := writer.Write(buf[:]); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// type releaseRQPDU0 struct {
// 	res [4]byte
// }
//
// func readReleaseRQPDU0(reader io.Reader) (*releaseRQPDU0, error) {
//
// 	var buf [4]byte
// 	if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 		return nil, err
// 	}
//
// 	pdu := &releaseRQPDU0{}
// 	pdu.res = buf
//
// 	return pdu, nil
// }
//
// func (pdu *releaseRQPDU0) write(writer io.Writer) error {
//
// 	buf := pdu.res
//
// 	if err := newPDU(aReleaseRQPDUType, uint32(len(buf))).write(writer); err != nil {
// 		return err
// 	}
//
// 	if _, err := writer.Write(buf[:]); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// type releaseRPPDU0 struct {
// 	res [4]byte
// }
//
// func readReleaseRPPDU0(reader io.Reader) (*releaseRPPDU0, error) {
//
// 	var buf [4]byte
// 	if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 		return nil, err
// 	}
//
// 	pdu := &releaseRPPDU0{}
// 	pdu.res = buf
//
// 	return pdu, nil
// }
//
// func (pdu *releaseRPPDU0) write(writer io.Writer) error {
//
// 	buf := pdu.res
//
// 	if err := newPDU(aReleaseRQPDUType, uint32(len(buf))).write(writer); err != nil {
// 		return err
// 	}
//
// 	if _, err := writer.Write(buf[:]); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// // // item types and sub item types
// // const (
// // 	appContextItemType      = 0x10
// // 	rqPCItemType            = 0x20
// // 	acPCItemType            = 0x21
// // 	abstractSyntaxItemType  = 0x30
// // 	transferSyntaxItemType  = 0x40
// // 	userInfoItemType        = 0x50
// // 	maxLengthItemType       = 0x51
// // 	implClassUIDItemType    = 0x52
// // 	maxNumOpsItemType       = 0x53
// // 	implVersionNameItemType = 0x55
// // )
//
// // assocACRQPDU presents an Association Accept and an Association Request PDU
// type assocACRQPDU0 struct {
// 	protocol       uint16
// 	calledAETitle  string
// 	callingAETitle string
// 	appContextName string
// 	pcs            []*pc
// 	userInfo       *userInfo
// }
//
// func (pdu *assocACRQPDU0) String() string {
// 	return fmt.Sprintf(
// 		"{protocol:%v,calledAET:%q,callingAET:%q,appContextName:%q,pcs:%s,userInfo:%s}",
// 		pdu.protocol,
// 		strings.TrimSpace(pdu.calledAETitle),
// 		strings.TrimSpace(pdu.callingAETitle),
// 		pdu.appContextName,
// 		pdu.pcs,
// 		pdu.userInfo,
// 	)
// }
//
// func readAssocACRQPDU0(reader io.Reader, pcItemType byte) (*assocACRQPDU0, error) {
//
// 	var buf [68]byte
// 	if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 		return nil, err
// 	}
//
// 	protocol := binary.BigEndian.Uint16(buf[0:2])
//
// 	calledAETitle := string(buf[4:20])
// 	callingAETitle := string(buf[20:36])
//
// 	pdu := &assocACRQPDU0{}
// 	pdu.protocol = protocol
// 	pdu.calledAETitle = calledAETitle
// 	pdu.callingAETitle = callingAETitle
// 	pdu.pcs = make([]*pc, 0, 5)
//
// 	for {
//
// 		// read the item time and the length
// 		var buf [4]byte
// 		if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 			return nil, err
// 		}
//
// 		// read the item type
// 		itemType := buf[0]
//
// 		// read the length
// 		length := binary.BigEndian.Uint16(buf[2:4])
//
// 		if itemType == appContextItemType { // application context name
//
// 			buf := make([]byte, length)
// 			if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 				return nil, err
// 			}
// 			pdu.appContextName = string(buf[:])
//
// 		} else if itemType == pcItemType { // presentation context item type
//
// 			// create a limited reader for the requested presentation contextx
// 			limitedReader := io.LimitReader(reader, int64(length))
//
// 			// read the presentation context
// 			pc, err := readPC(limitedReader, itemType)
// 			if err != nil {
// 				return nil, err
// 			}
//
// 			// add it to the list of requested presentation contexts
// 			pdu.pcs = append(pdu.pcs, pc)
//
// 		} else if itemType == userInfoItemType { // user info
//
// 			// create a limited reader for the user info
// 			limitedReader := io.LimitReader(reader, int64(length))
//
// 			// read the user info sub items
// 			userInfo, err := readUserInfo(limitedReader)
// 			if err != nil {
// 				return nil, err
// 			}
// 			pdu.userInfo = userInfo
//
// 		} else {
// 			// unrecognized item
// 			return nil, fmt.Errorf("unrecognized item type: 0x%02X", itemType)
// 		}
// 	}
//
// 	// return the pdu
// 	return pdu, nil
// }
//
// func (pdu *assocACRQPDU0) write(writer io.Writer, pduType byte, pcItemType byte) error {
//
// 	// // create the pdu
// 	// pdu := &pdu{typ: pduType}
// 	//
// 	// // create a byte array output stream so we can calculate the length of the rest of the PDU
// 	//
//
// 	var buf [68]byte
//
// 	binary.BigEndian.PutUint16(buf[0:2], pdu.protocol)
// 	buf[2] = 0x00
// 	buf[3] = 0x00
// 	copy(buf[4:20], padAETitle(pdu.calledAETitle))
// 	copy(buf[20:36], padAETitle(pdu.callingAETitle))
// 	for i := 36; i < 68; i++ {
// 		buf[i] = 0x00
// 	}
//
// 	byteWriter := new(bytes.Buffer)
//
// 	// write the variable items
// 	if err := writeVariableItems(byteWriter, pdu, pcItemType); err != nil {
// 		return err
// 	}
//
// 	// write the byte array to the original writer
// 	if err := writeBytes(pdu, byteWriter.Bytes()); err != nil {
// 		return err
// 	}
//
// 	// write the pdu
// 	if err := writePDU(writer, pdu); err != nil {
// 		return err
// 	}
//
// 	// all is good
// 	return nil
// }
//
// func padAETitle0(aeTitle string) string {
// 	return fmt.Sprintf("%-16s", aeTitle)
// }
//
// func writeVariableItems0(writer io.Writer, assocACRQPDU *assocACRQPDU, itemType byte) error {
//
// 	if err := writeAppContextName(writer, assocACRQPDU.appContextName); err != nil {
// 		return err
// 	}
//
// 	if err := writePCs(writer, assocACRQPDU.pcs, itemType); err != nil {
// 		return err
// 	}
//
// 	if err := writeUserInfo(writer, assocACRQPDU.userInfo); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func writeAppContextName0(writer io.Writer, appContextName string) error {
//
// 	var buf [4]byte
//
// 	buf[0] = appContextItemType
// 	buf[1] = 0x00
// 	binary.BigEndian.PutUint16(buf[2:4], uint16(len(appContextName)))
//
// 	if _, err := writer.Write(buf[:]); err != nil {
// 		return err
// 	}
//
// 	if _, err := writer.Write([]byte(appContextName)); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
