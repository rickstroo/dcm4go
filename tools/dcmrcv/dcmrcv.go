package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/rickstroo/dcm4go/dcm4go"
)

// simple error management
func check(err error) {
	if err != nil {
		fmt.Printf("panic: %v\n", err)
		//		panic(err)
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

	// listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:4104")
	check(err)

	// ensure the listener gets closed
	defer listener.Close()

	// let us know where we are listening
	fmt.Printf("listening on %v\n", listener.Addr())

	// listen for connections
	for {

		// wait for connection
		conn, err := listener.Accept()
		check(err)

		fmt.Printf("accepted connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())

		// handle the connection
		err = handleConnection(conn)
		check(err)

		// break for now, so that we don't keep the port open
		break
	}
}

func handleConnection(conn net.Conn) error {

	// accept the association
	assoc, err := dcm4go.AcceptAssoc(conn)
	check(err)

	// close the association
	err = assoc.Close()
	check(err)

	// close the connection
	err = conn.Close()
	check(err)

	fmt.Printf("closed connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())

	return nil
}
