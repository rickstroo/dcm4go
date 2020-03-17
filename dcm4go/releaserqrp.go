// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains the definition and structures and methods used
// to release associations.

package dcm4go

import "io"

// AReleaseRQRP defines a structure and methods used to send and receive
// release request and response messages.
type AReleaseRQRP struct {
}

// Read reads an AReleaseRQ from a reader.
func (aReleaseRQRP *AReleaseRQRP) Read(reader io.Reader) error {
	return nil
}

// Write write an AReleaseRQ to a writer.
func (aReleaseRQRP *AReleaseRQRP) Write(writer io.Writer) error {
	return nil
}

// AReleaseRQ defines a structure and methods used to send and receive
// associate release request messages.  It inherits all of its structures and
// methods from AReleaseRQRP.
type AReleaseRQ struct {
	AReleaseRQRP
}

// ReadAReleaseRQ reads an AReleaseRQ.
func ReadAReleaseRQ(reader io.Reader) (*AReleaseRQ, error) {
	aReleaseRQ := &AReleaseRQ{}
	if err := aReleaseRQ.Read(reader); err != nil {
		return nil, err
	}
	return aReleaseRQ, nil
}

// AReleaseRP defines a structure and methods used to send and receive
// associate release response messages.  It inherits all of its structures and
// methods from AReleaseRQRP.
type AReleaseRP struct {
	AReleaseRQRP
}

// ReadAReleaseRP reads an AReleaseRP.
func ReadAReleaseRP(reader io.Reader) (*AReleaseRP, error) {
	aReleaseRP := &AReleaseRP{}
	if err := aReleaseRP.Read(reader); err != nil {
		return nil, err
	}
	return aReleaseRP, nil
}
