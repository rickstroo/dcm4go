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

	// JPEGBaselineLossy8BitImageCompressionTS is transfer syntax uid for JPEG lossy 8 bit image compression
	JPEGBaselineLossy8BitImageCompressionTS = &TransferSyntax{JPEGBaselineLossy8BitImageCompressionUID, true, binary.LittleEndian, "JPEG Baseline Lossy 8 Bit Image Compression"}

	// JPEGBaselineLossy12BitImageCompressionTS is transfer syntax uid for JPEG lossy 12 bit image compression
	JPEGBaselineLossy12BitImageCompressionTS = &TransferSyntax{JPEGBaselineLossy12BitImageCompressionUID, true, binary.LittleEndian, "JPEG Baseline Lossy 12 Bit Image Compression"}

	// JPEGLosslessNonHierarchicalImageCompressionTS is transfer syntax uid for JPEG lossless non-hierarchical image compression
	JPEGLosslessNonHierarchicalImageCompressionTS = &TransferSyntax{JPEGLosslessNonHierarchicalImageCompressionUID, true, binary.LittleEndian, "JPEG Lossless Non-Hierarchical Image Compression"}

	// JPEGLSLosslessImageCompressionTS is transfer syntax uid for JPEG-LS lossless image compression
	JPEGLSLosslessImageCompressionTS = &TransferSyntax{JPEGLSLosslessImageCompressionUID, true, binary.LittleEndian, "JPEG-LS Lossless Image Compression"}

	// JPEGLSLossyImageCompressionTS is transfer syntax uid for JPEG-LS lossy (near lossless) image compression
	JPEGLSLossyImageCompressionTS = &TransferSyntax{JPEGLSLossyImageCompressionUID, true, binary.LittleEndian, "JPEG-LS Lossy Image Compression"}

	// JPEG2000ImageCompressionLosslessOnlyTS is transfer syntax for JPEG 2000 image compression lossless only
	JPEG2000ImageCompressionLosslessOnlyTS = &TransferSyntax{JPEG2000ImageCompressionLosslessOnlyUID, true, binary.LittleEndian, "JPEG 2000 Image Compression (Lossless Only)"}

	// JPEG2000ImageCompressionTS is transfer syntax for JPEG 2000 image compression
	JPEG2000ImageCompressionTS = &TransferSyntax{JPEG2000ImageCompressionUID, true, binary.LittleEndian, "JPEG 2000 Image Compression"}

	// JPEG2000Part2MulticomponentImageCompressionLosslessOnlyTS is transfer syntax uid for JPEG 2000 part 2 multi-component image compression lossless only
	JPEG2000Part2MulticomponentImageCompressionLosslessOnlyTS = &TransferSyntax{JPEG2000Part2MulticomponentImageCompressionLosslessOnlyUID, true, binary.LittleEndian, "JPEG 2000 Part2 Multi-component Image Compression (Lossless Only)"}

	// JPEG2000Part2MulticomponentImageCompressionTS is transfer syntax uid for JPEG 2000 image compression
	JPEG2000Part2MulticomponentImageCompressionTS = &TransferSyntax{JPEG2000Part2MulticomponentImageCompressionUID, true, binary.LittleEndian, "JPEG 2000 Part2 Multi-component Image Compression"}

	// JPIPReferencedTS is transfer syntax uid for JPIP referenced
	JPIPReferencedTS = &TransferSyntax{JPIPReferencedUID, true, binary.LittleEndian, "JPIP Referenced"}

	// JPIPReferencedDeflateTS is transfer syntax uid for JPIP referenced deflate
	JPIPReferencedDeflateTS = &TransferSyntax{JPIPReferencedDeflateUID, true, binary.LittleEndian, "JPIP Referenced Deflate"}

	// RLELosslessTS is transfer syntax uid for RLE lossless
	RLELosslessTS = &TransferSyntax{RLELosslessUID, true, binary.LittleEndian, "RLE Loessless"}

	// RFC2557MimeEncapsulationTS is transfer syntax uid for RFC 2557 mime encapsulation
	RFC2557MimeEncapsulationTS = &TransferSyntax{RFC2557MimeEncapsulationUID, true, binary.LittleEndian, "RFC 2557 Mime Encapsulation"}
)

var tses = map[string]*TransferSyntax{
	ImplicitVRLittleEndianUID:                                  ImplicitVRLittleEndianTS,
	ExplicitVRLittleEndianUID:                                  ExplicitVRLittleEndianTS,
	DeflatedExplicitVRLittleEndianUID:                          DeflatedExplicitVRLittleEndianTS,
	ExplicitVRBigEndianUID:                                     ExplicitVRBigEndianTS,
	JPEGBaselineLossy8BitImageCompressionUID:                   JPEGBaselineLossy8BitImageCompressionTS,
	JPEGBaselineLossy12BitImageCompressionUID:                  JPEGBaselineLossy12BitImageCompressionTS,
	JPEGLosslessNonHierarchicalImageCompressionUID:             JPEGLosslessNonHierarchicalImageCompressionTS,
	JPEGLSLosslessImageCompressionUID:                          JPEGLSLosslessImageCompressionTS,
	JPEGLSLossyImageCompressionUID:                             JPEGLSLossyImageCompressionTS,
	JPEG2000ImageCompressionLosslessOnlyUID:                    JPEG2000ImageCompressionLosslessOnlyTS,
	JPEG2000ImageCompressionUID:                                JPEG2000ImageCompressionTS,
	JPEG2000Part2MulticomponentImageCompressionLosslessOnlyUID: JPEG2000Part2MulticomponentImageCompressionLosslessOnlyTS,
	JPEG2000Part2MulticomponentImageCompressionUID:             JPEG2000Part2MulticomponentImageCompressionTS,
	JPIPReferencedUID:                                          JPIPReferencedTS,
	JPIPReferencedDeflateUID:                                   JPIPReferencedDeflateTS,
	RLELosslessUID:                                             RLELosslessTS,
	RFC2557MimeEncapsulationUID:                                RFC2557MimeEncapsulationTS,
}
