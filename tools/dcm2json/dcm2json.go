package main

import (
	"flag"
	"fmt"
	"os"

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

	// open the file, which returns a reader, defer a close
	file, err := os.Open(*path)
	check(err)

	// make sure we close the file upon exit
	defer file.Close()

	// read the file
	object, err := dcm4go.ReadFile(file, uint32(*bulkDataThreshold))
	check(err)

	// print the object
	print(*path, object)
}

// print the object
func print(path string, object *dcm4go.Object) {
	fmt.Printf("%s\n", dcm4go.ObjectToJSON(path, object))
}
