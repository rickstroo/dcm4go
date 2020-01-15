package dcm4go

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// UserInfo represents user information
type UserInfo struct {
	maxLenReceived     uint32
	implClassUID       string
	implVersionName    string
	maxNumOpsInvoked   uint16
	maxNumOpsPerformed uint16
}

// String returns a string representation of a UserInfo
func (userInfo *UserInfo) String() string {
	return fmt.Sprintf(
		"{maxLenReceived:%v,implClassUID:%q,implVersionName:%q,maxNumOpsInvoked:%v,maxNumOpsPerformed:%v}",
		userInfo.maxLenReceived,
		userInfo.implClassUID,
		userInfo.implVersionName,
		userInfo.maxNumOpsInvoked,
		userInfo.maxNumOpsPerformed)
}

func readUserInfo(reader io.Reader) (*UserInfo, error) {

	// initialize a user info object
	userInfo := &UserInfo{}

	// read the user info sub items until eof
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

		if subItemType == 0x51 { // maximum length

			// read the length
			length, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}

			// check it
			if length != 0x04 {
				return nil, fmt.Errorf("expected length to be 0x04, was 0x%04X", length)
			}

			// read the maximum length received
			maxLenReceived, err := readLong(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}
			userInfo.maxLenReceived = maxLenReceived

		} else if subItemType == 0x52 { // implementation class UID

			// read the length
			length, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}

			// read the implementation class UID
			implClassUID, err := readUID(reader, uint32(length))
			if err != nil {
				return nil, err
			}
			userInfo.implClassUID = implClassUID

		} else if subItemType == 0x53 { // maximum number operations

			// read the length
			length, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}

			// check it
			if length != 0x04 {
				return nil, fmt.Errorf("expected length to be 0x04, was 0x%04X", length)
			}

			// read the maximum number of operations invoked
			maxNumOpsInvoked, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}
			userInfo.maxNumOpsInvoked = maxNumOpsInvoked

			// read the maximum number of operations performed
			maxNumOpsPerformed, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}
			userInfo.maxNumOpsPerformed = maxNumOpsPerformed

		} else if subItemType == 0x55 { // implementation version name

			// read the length
			length, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}

			// read the implementation version name
			implVersionName, err := readText(reader, uint32(length))
			if err != nil {
				return nil, err
			}
			userInfo.implVersionName = implVersionName

		} else {

			// unrecognized item
			fmt.Printf("ignoring unrecognized user info sub item type: 0x%02X\n", subItemType)

			// read the length
			length, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}

			// skip the bytes
			if err := skipBytes(reader, uint32(length)); err != nil {
				return nil, err
			}
		}

	}

	// return the user info
	return userInfo, nil
}

func writeUserInfo(writer io.Writer, userInfo *UserInfo) error {
	return fmt.Errorf("writeUserInfo: not implemented")
}
