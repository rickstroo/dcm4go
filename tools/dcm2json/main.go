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

	// parse the flags
	flag.Parse()

	// read the file
	groupTwo, otherGroups, err := dcm4go.ReadFile(*path)
	check(err)

	// print the group two object
	fmt.Printf("%s", groupTwo.String())

	// print the other groups
	fmt.Printf("%s", otherGroups.String())
}
