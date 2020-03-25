// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains the definition and structures and methods used
// to abort associations.

package dcm4go

// import "io"
//
// // AAbort defines a structure and methods used to send and receive
// // abort request and response messages.
// type AAbort struct {
// 	Source byte
// 	Reason byte
// }
//
// // Read reads an AAbort from a reader.
// func (aAbort *AAbort) Read(reader io.Reader) error {
// 	return nil
// }
//
// // ReadAAbort reads an AAbort from a reader.
// func ReadAAbort(reader io.Reader) (*AAbort, error) {
// 	aAbort := &AAbort{}
// 	if err := aAbort.Read(reader); err != nil {
// 		return nil, err
// 	}
// 	return aAbort, nil
// }
//
// // Write write an AAbort to a writer.
// func (aAbort *AAbort) Write(writer io.Writer) error {
// 	return nil
// }
