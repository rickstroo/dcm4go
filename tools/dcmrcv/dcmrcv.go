package main

import (
	"flag"
	"fmt"
	"net"
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
		connection, err := listener.Accept()
		check(err)

		fmt.Printf("accepted connection on %v from %v\n", connection.LocalAddr(), connection.RemoteAddr())

		// handle the connection
		err = handleConnection(connection)
		check(err)
	}
}

func handleConnection(connection net.Conn) error {

	// close the connection
	connection.Close()

	fmt.Printf("closed connection on %v from %v\n", connection.LocalAddr(), connection.RemoteAddr())

	return nil
}
