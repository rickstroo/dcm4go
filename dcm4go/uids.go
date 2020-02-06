package dcm4go

const (

	// UIDs for Transfer Syntaxes

	// ImplicitVRLittleEndianUID is transfer syntax for implicit VR little endian
	ImplicitVRLittleEndianUID = "1.2.840.10008.1.2"

	// ExplicitVRLittleEndianUID is transfer syntax for explicit VR little endian
	ExplicitVRLittleEndianUID = "1.2.840.10008.1.2.1"

	// DeflatedExplicitVRLittleEndianUID is transfer syntax for deflated explicit VR little endian
	DeflatedExplicitVRLittleEndianUID = "1.2.840.10008.1.2.1.99"

	// ExplicitVRBigEndianUID is transfer syntax for explicit VR big endian
	ExplicitVRBigEndianUID = "1.2.840.10008.1.2.2"

	// JPEGBaselineLossy8BitImageCompressionUID is transfer syntax uid for JPEG lossy 8 bit image compression
	JPEGBaselineLossy8BitImageCompressionUID = "1.2.840.10008.1.2.4.50"

	// JPEGBaselineLossy12BitImageCompressionUID is transfer syntax uid for JPEG lossy 12 bit image compression
	JPEGBaselineLossy12BitImageCompressionUID = "1.2.840.10008.1.2.4.51"

	// JPEGLosslessNonHierarchicalImageCompressionUID is transfer syntax uid for JPEG lossless non-hierarchical image compression
	JPEGLosslessNonHierarchicalImageCompressionUID = "1.2.840.10008.1.2.4.70"

	// JPEGLSLosslessImageCompressionUID is transfer syntax uid for JPEG-LS lossless image compression
	JPEGLSLosslessImageCompressionUID = "1.2.840.10008.1.2.4.80"

	// JPEGLSLossyImageCompressionUID is transfer syntax uid for JPEG-LS lossy (near lossless) image compression
	JPEGLSLossyImageCompressionUID = "1.2.840.10008.1.2.4.81"

	// JPEG2000ImageCompressionLosslessOnlyUID is transfer syntax uid for JPEG 2000 image compression (lossless only)
	JPEG2000ImageCompressionLosslessOnlyUID = "1.2.840.10008.1.2.4.90"

	// JPEG2000ImageCompressionUID is transfer syntax uid for JPEG 2000 image compression
	JPEG2000ImageCompressionUID = "1.2.840.10008.1.2.4.91"

	// JPEG2000Part2MulticomponentImageCompressionLosslessOnlyUID is transfer syntax uid for JPEG 2000 part 2 multi-component image compression lossless only
	JPEG2000Part2MulticomponentImageCompressionLosslessOnlyUID = "1.2.840.10008.1.2.4.92"

	// JPEG2000Part2MulticomponentImageCompressionUID is transfer syntax uid for JPEG 2000 image compression
	JPEG2000Part2MulticomponentImageCompressionUID = "1.2.840.10008.1.2.4.93"

	// JPIPReferencedUID is transfer syntax uid for JPIP referenced
	JPIPReferencedUID = "1.2.840.10008.1.2.4.94"

	// JPIPReferencedDeflateUID is transfer syntax uid for JPIP referenced deflate
	JPIPReferencedDeflateUID = "1.2.840.10008.1.2.4.95"

	// RLELosslessUID is transfer syntax uid for RLE lossless
	RLELosslessUID = "1.2.840.10008.1.2.5"

	// RFC2557MimeEncapsulationUID is transfer syntax uid for RFC 2557 mime encapsulation
	RFC2557MimeEncapsulationUID = "1.2.840.10008.1.2.6.1"

	// UIDs for Services

	// VerificationUID is the SOP class UID for verification
	VerificationUID = "1.2.840.10008.1.1"

	// EnhancedXAImageStorageUID is the SOP class UID for enhanced XA image storage
	EnhancedXAImageStorageUID = "1.2.840.10008.5.1.4.1.1.12.1.1"

	// GeneralECGWaveformStorageUID is the SOP class UID for general ECG waveform storage
	GeneralECGWaveformStorageUID = "1.2.840.10008.5.1.4.1.1.9.1.2"

	// UIDs for Other Purposes

	// ApplicationContextNameUID is the UID for the Application Context Name
	ApplicationContextNameUID = "1.2.840.10008.3.1.1.1"
)
