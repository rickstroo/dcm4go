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

	// create a client
	client := &dcm4go.Client{
		AETitle: "DCMSND",
	}

	// verify
	check(client.Verify(addr))

	// send the files
	check(client.Send(addr, paths))
}
