package dcm4go

import "encoding/binary"

var (
	// ImplicitVRLittleEndianTS is transfer syntax for implicit VR little endian
	ImplicitVRLittleEndianTS = &TransferSyntax{ImplicitVRLittleEndianUID, false, binary.LittleEndian, "Implicit VR Little Endian"}
	// ExplicitVRLittleEndianTS is transfer syntax for explicit VR little endian
	ExplicitVRLittleEndianTS = &TransferSyntax{ExplicitVRLittleEndianUID, true, binary.LittleEndian, "Explicit VR Little Endian"}
	// DeflatedExplicitVRLittleEndianTS is transfer syntax for deflated explicit VR little endian
	DeflatedExplicitVRLittleEndianTS = &TransferSyntax{DeflatedExplicitVRLittleEndianUID, true, binary.LittleEndian, "Deflated Explicit VR Little Endian"}
	// ExplicitVRBigEndianTS is transfer syntax for explicit VR big endian
	ExplicitVRBigEndianTS = &TransferSyntax{ExplicitVRBigEndianUID, true, binary.BigEndian, "Explicit VR Big Endian"}
	// JPEG2000ImageCompressionLosslessOnlyTS is transfer syntax for JPEG 2000 image compression lossless only
	JPEG2000ImageCompressionLosslessOnlyTS = &TransferSyntax{JPEG2000ImageCompressionLosslessOnlyUID, true, binary.LittleEndian, "JPEG 2000 Image Compression (Lossless Only)"}
	// JPEG2000ImageCompressionTS is transfer syntax for JPEG 2000 image compression
	JPEG2000ImageCompressionTS = &TransferSyntax{JPEG2000ImageCompressionUID, true, binary.LittleEndian, "JPEG 2000 Image Compression"}
)

var tses = map[string]*TransferSyntax{
	ImplicitVRLittleEndianUID:               ImplicitVRLittleEndianTS,
	ExplicitVRLittleEndianUID:               ExplicitVRLittleEndianTS,
	DeflatedExplicitVRLittleEndianUID:       DeflatedExplicitVRLittleEndianTS,
	ExplicitVRBigEndianUID:                  ExplicitVRBigEndianTS,
	JPEG2000ImageCompressionLosslessOnlyUID: JPEG2000ImageCompressionLosslessOnlyTS,
	JPEG2000ImageCompressionUID:             JPEG2000ImageCompressionTS,
}
