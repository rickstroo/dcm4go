package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		log.Printf("error is %v", err)
		//		panic(err)
		os.Exit(1)
	}
}

// the main function
func main() {

	var local string
	var remote string
	var path string
	var help bool

	flag.StringVar(&local, "local", "DCMSND", "AE title for local client")
	flag.StringVar(&remote, "remote", "DCMRCV@localhost:4104", "AE title, host and port for remote server of form 'ae@host:port'")
	flag.StringVar(&path, "path", "", "comma separated list of files to send")
	flag.BoolVar(&help, "help", false, "display usage")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	paths := strings.Split(path, ",")

	fmt.Printf("local is %q\n", local)
	fmt.Printf("remote is %q\n", remote)
	fmt.Printf("paths is %q\n", paths)

	// create a client
	client := &dcm4go.Client{
		AETitle: local,
	}

	// we're going to use a C-Echo service
	cEchoService := &dcm4go.CEchoService{}

	// we're also going to use a C-Store service
	cStoreService := &dcm4go.CStoreService{}

	// discover the capabilities of the files that we are going to send
	capabilities, err := dcm4go.ReadCapabilities(paths)
	check(err)
	cStoreService.Capabilities = capabilities

	// connect with the SCUs that we are going to use
	check(client.Connect(remote, cEchoService, cStoreService))

	// defer close
	defer func() {
		client.Close()
	}()

	// verify
	check(client.Verify())

	// send the files
	check(client.Send(paths))

}
