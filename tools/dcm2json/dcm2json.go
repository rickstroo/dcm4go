package main

import (
	"flag"
	"fmt"

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

	// get the path to the file
	path := flag.String("path", "-", "path to file to convert from DICOM to JSON")

	// threshold for bulk data
	bulkDataThreshold := flag.Uint("bulk", 1024, "threshold for binary bulk data")

	// parse the flags
	flag.Parse()

	// read the file
	groupTwo, otherGroups, err := dcm4go.ReadFile(*path, uint32(*bulkDataThreshold))
	check(err)

	// print the group two object
	fmt.Printf("%s", groupTwo.String())

	// print the other groups
	fmt.Printf("%s", otherGroups.String())
}
