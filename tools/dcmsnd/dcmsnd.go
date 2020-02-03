package main

import (
	"flag"
	"fmt"
	"strings"

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

	var local string
	var remote string
	var path string

	flag.StringVar(&local, "local", "DCMSND", "AE title for local client")
	flag.StringVar(&remote, "remote", "DCMRCV@localhost:4104", "AE title, host and port for remote server")
	flag.StringVar(&path, "path", "", "path to command separated list of files to send")

	flag.Parse()

	paths := strings.Split(path, ",")

	fmt.Printf("local is %q\n", local)
	fmt.Printf("remote is %q\n", remote)
	fmt.Printf("paths is %q\n", paths)

	// create a client
	client := &dcm4go.Client{
		AETitle: local,
	}

	// verify
	check(client.Verify(remote))

	// send the files
	check(client.Send(remote, paths))
}
