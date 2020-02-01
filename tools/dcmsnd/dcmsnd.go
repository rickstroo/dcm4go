package main

import (
	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// the main function
func main() {

	// set address of server and paths to file
	// will parse command line args later
	addr := "DCMRCV@localhost:4104"
	paths := []string{
		"/Users/Rick/data/dicom/ENHXA.dcm",
		"/Users/Rick/data/dicom/GENECG.dcm",
	}

	// try a verify, then send the files
	check(dcm4go.Verify(addr))
	check(dcm4go.SendN(addr, paths))
}
