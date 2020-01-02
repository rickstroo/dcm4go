package main

import (
	"flag"
	"fmt"
)

// the main function
func main() {

	// get the path to the file
	path := flag.String("path", "-", "path to file to convert from DICOM to JSON")
	flag.Parse()

	// simple message to start with
	fmt.Printf("about to open file '%s'\n", *path)
}
