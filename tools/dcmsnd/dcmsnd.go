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

	// set args, will parse command line later
	local := "DCMSND"                 // address of local client
	remote := "DCMRCV@localhost:4104" // address of remote server
	paths := []string{                // paths to files to send
		"/Users/Rick/data/dicom/ENHXA.dcm",
		"/Users/Rick/data/dicom/GENECG.dcm",
	}

	// create a client
	client := &dcm4go.Client{
		AETitle: local,
	}

	// verify
	check(client.Verify(remote))

	// send the files
	check(client.Send(remote, paths))
}
