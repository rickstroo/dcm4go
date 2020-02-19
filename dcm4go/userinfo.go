package dcm4go

import (
	"bytes"
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

		if subItemType == maxLengthItemType { // maximum length

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

		} else if subItemType == implClassUIDItemType { // implementation class UID

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

		} else if subItemType == maxNumOpsItemType { // maximum number operations

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

		} else if subItemType == implVersionNameItemType { // implementation version name

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

	// write the item type
	if err := writeByte(writer, userInfoItemType); err != nil {
		return err
	}

	// write the padding zero
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// create a byte array output stream so we can calculate the length of the rest of the PDU
	byteWriter := new(bytes.Buffer)

	if err := writeMaxLenReceived(byteWriter, userInfo); err != nil {
		return err
	}

	if err := writeImplClassUID(byteWriter, userInfo); err != nil {
		return err
	}

	if err := writeMaxNumOps(byteWriter, userInfo); err != nil {
		return err
	}

	// write the length to the original writer
	if err := writeShort(writer, uint16(byteWriter.Len()), binary.BigEndian); err != nil {
		return err
	}

	// write the byte array to the original writer
	if err := writeBytes(writer, byteWriter.Bytes()); err != nil {
		return err

	}

	// all is well
	return nil
}

func writeMaxLenReceived(writer io.Writer, userInfo *UserInfo) error {

	// write the sub item type
	if err := writeByte(writer, maxLengthItemType); err != nil {
		return err
	}

	// write the padding zero
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// write the length, must be four
	if err := writeShort(writer, 0x04, binary.BigEndian); err != nil {
		return err
	}

	// write the maximum length received
	if err := writeLong(writer, userInfo.maxLenReceived, binary.BigEndian); err != nil {
		return err
	}

	// all is well
	return nil
}

func writeImplClassUID(writer io.Writer, userInfo *UserInfo) error {

	// write the sub item type
	if err := writeByte(writer, implClassUIDItemType); err != nil {
		return err
	}

	// write the padding zero
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// write the length of the implementation class uid
	if err := writeShort(writer, uint16(len(userInfo.implClassUID)), binary.BigEndian); err != nil {
		return err
	}

	// write the implementation class uid
	if err := writeString(writer, userInfo.implClassUID); err != nil {
		return err
	}

	// all is well
	return nil
}

func writeMaxNumOps(writer io.Writer, userInfo *UserInfo) error {

	// write the sub item type
	if err := writeByte(writer, maxNumOpsItemType); err != nil {
		return err
	}

	// write the padding zero
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// write the length, must be four
	if err := writeShort(writer, 0x04, binary.BigEndian); err != nil {
		return err
	}

	// write the max num ops invoked
	if err := writeShort(writer, userInfo.maxNumOpsInvoked, binary.BigEndian); err != nil {
		return err
	}

	// write the max num ops performance
	if err := writeShort(writer, userInfo.maxNumOpsPerformed, binary.BigEndian); err != nil {
		return err
	}

	// all is well
	return nil
}

func writeImplVersionName(writer io.Writer, userInfo *UserInfo) error {

	// write the sub item type
	if err := writeByte(writer, implVersionNameItemType); err != nil {
		return err
	}

	// write the padding zero
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// write the length of the implementation class uid
	if err := writeShort(writer, uint16(len(userInfo.implVersionName)), binary.BigEndian); err != nil {
		return err
	}

	// write the implementation version name
	if err := writeString(writer, userInfo.implVersionName); err != nil {
		return err
	}

	// all is well
	return nil
}
