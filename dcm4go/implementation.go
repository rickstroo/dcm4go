package dcm4go

const (
	// ProtocolVersion is the version of the DICOM protocol
	ProtocolVersion = 0x01

	// ImplementationClassUID is the UID for this implementation, we have our own now
	ImplementationClassUID = "1.3.6.1.4.1.55242"

	// ImplementationVersionName is the version name for this implementation
	ImplementationVersionName = "dcm4go-1.0"

	// DefaultMaxPDULen is the default maximum length of a PDU
	DefaultMaxPDULen = 16 * 2014

	// DefaultMaxNumOps is the default maximum number of ops
	DefaultMaxNumOps = 0x00

	// DefaultMaxOpsPerformed is the default maximum number of ops performed
	DefaultMaxOpsPerformed = 0x00
)
