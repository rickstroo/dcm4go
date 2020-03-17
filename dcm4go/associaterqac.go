// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains the definition and structures and methods used
// to request and accept associations.

package dcm4go

import "io"

// AAssociateRQAC defines a structure and methods used to send and receive
// associate request and associate accept messages.
type AAssociateRQAC struct {
	ProtocolVersion byte
	CallingAETitle  string
	CalledAETitle   string
	PresContexts    *[]PresContext
	UserInfo        *UserInfo
}

// A PresContext defines a presentation context.
type PresContext struct {
}

// A UserInfo defines user information.
type UserInfo struct {
}

func (aAssociateRQAC *AAssociateRQAC) Read(reader io.Reader) error {
	return nil
}

func (aAssociateRQAC *AAssociateRQAC) Write(writer io.Writer) error {
	return nil
}

// AAssociateRQ defines a structure and methods used to send and receive
// associate request messages.  It inherits all of its structures and methods
// from AAssociateRQAC.
type AAssociateRQ struct {
	AAssociateRQAC
}

// ReadAAssociateRQ reads an AAssociateAC.
func ReadAAssociateRQ(reader io.Reader) (*AAssociateRQ, error) {
	aAssociateRQ := &AAssociateRQ{}
	if err := aAssociateRQ.Read(reader); err != nil {
		return nil, err
	}
	return aAssociateRQ, nil
}

// AAssociateAC defines a structure and methods used to send and receive
// associate accept messages.  It inherits all of its structures and methods
// from AAssociateRQAC.
type AAssociateAC struct {
	AAssociateRQAC
}

// ReadAAssociateAC reads an AAssociateAC.
func ReadAAssociateAC(reader io.Reader) (*AAssociateAC, error) {
	aAssociateAC := &AAssociateAC{}
	if err := aAssociateAC.Read(reader); err != nil {
		return nil, err
	}
	return aAssociateAC, nil
}

// AAssociateRJ defines a structure and methods used to send and receive
// associate reject messages.
type AAssociateRJ struct {
}

// ReadAAssociateRJ reads an AssociateRJ.
func (aAssociateRJ *AAssociateRJ) Read(reader io.Reader) error {
	return nil
}

// ReadAAssociateRJ reads an AssociateRJ.
func ReadAAssociateRJ(reader io.Reader) (*AAssociateRJ, error) {
	aAssociateRJ := &AAssociateRJ{}
	if err := aAssociateRJ.Read(reader); err != nil {
		return nil, err
	}
	return aAssociateRJ, nil
}

func (aAssociateRJ *AAssociateRJ) Write(writer io.Writer) error {
	return nil
}
