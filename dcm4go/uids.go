package dcm4go

const (
	// ImplicitVRLittleEndianUID is transfer syntax for implicit VR little endian
	ImplicitVRLittleEndianUID = "1.2.840.10008.1.2"
	// ExplicitVRLittleEndianUID is transfer syntax for explicit VR little endian
	ExplicitVRLittleEndianUID = "1.2.840.10008.1.2.1"
	// ExplicitVRBigEndianUID is transfer syntax for explicit VR big endian
	ExplicitVRBigEndianUID = "1.2.840.10008.1.2.2"
	// DeflatedExplicitVRLittleEndianUID is transfer syntax for deflated explicit VR little endian
	DeflatedExplicitVRLittleEndianUID = "1.2.840.10008.1.2.1.99"
	// JPEG2000ImageCompressionLosslessOnlyUID is transfer syntax uid for JPEG 2000 image compression (lossless only)
	JPEG2000ImageCompressionLosslessOnlyUID = "1.2.840.10008.1.2.4.90"
	// JPEG2000ImageCompressUID is transfer syntax uid for JPEG 2000 image compression
	JPEG2000ImageCompressionUID = "1.2.840.10008.1.2.4.91"

	// VerificationUID is the SOP class UID for verification
	VerificationUID = "1.2.840.10008.1.1"

	// EnhancedXAImageStorageUID is the SOP class UID for enhanced XA image storage
	EnhancedXAImageStorageUID = "1.2.840.10008.5.1.4.1.1.12.1.1"
	// GeneralECGWaveformStorageUID is the SOP class UID for general ECG waveform storage
	GeneralECGWaveformStorageUID = "1.2.840.10008.5.1.4.1.1.9.1.2"
)
