// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

// A Request represents a DICOM request.  There are many types of DICOM
// requests that will be defined (e.g. CEchoRequest, CStoreRequest).
// The presContext describes the presentation context that we used to
// transfer the request.
// The command describes the command portion of the request.  It is always
// filled in.
// The data describes the data portion of the request, if any data is present.
// If no data was included as part of the request, it will be nil.
// A handler of requests has the option of allowing the library to read the
// data or to have a reader to the data passed back to it.  If the handler
// chooses not to have the library read the data, the data will be nil, and
// the dataReader will be filled in, allowing the handler to manage the data
// directly.  It is assumed that this would allow for the more efficient
// processing of data for large data sets.
type Request struct {
	presContext *PresContext
	command     *Object
	data        *Object
	dataReader  *PDataReader
}
