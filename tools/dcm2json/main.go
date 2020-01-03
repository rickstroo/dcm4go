package main

import (
	"flag"
	"fmt"

	"github.com/rickstroo/dcm4go/file"
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
	object, err := file.ReadFile(*path)
	check(err)

	// print the object
	fmt.Printf("%v\n", object)
}
