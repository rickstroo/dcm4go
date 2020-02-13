package dcm4go

// A Response represents a response to a DICOM request.
type Response struct {
	presContext *PresContext
	command     *Object
	data        *Object
}
