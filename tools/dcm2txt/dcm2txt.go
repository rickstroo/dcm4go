package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		log.Fatalf("error is %v", err)
	}
}

// the main function
func main() {

	var path string
	var help bool

	flag.StringVar(&path, "path", "image.dcm", "specify path of file to parse")
	flag.BoolVar(&help, "help", false, "display usage")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	object, err := dcm4go.Parse(path)
	check(err)

	for _, attr := range object.Attributes() {
		fmt.Printf("%v\n", attr.Text())
	}
}
