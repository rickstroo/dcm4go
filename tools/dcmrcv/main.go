package main

import (
	"flag"
	"fmt"
)

// simple error management
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// the main function
func main() {

	// get the binding
	bind := flag.String("bind", "DCMRCV@localhost:4104", "AE Title, Host and Port to bind to")

	// parse the flags
	flag.Parse()

	// print the bind flat
	fmt.Printf("%s\n", *bind)
}
