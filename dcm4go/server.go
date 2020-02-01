// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"errors"
	"fmt"
	"net"
	"time"
)

// Errors used by the DICOM server.
var (
	// ErrUnrecognizedCallingAETitle is returned to client when the calling AE
	// Title is not recognized
	ErrUnrecognizedCallingAETitle = errors.New("dicom: unrecognized calling AE Title")

	// ErrUnrecognizedCalledAETitle is returned to client when the called AE
	// Title is not recognized
	ErrUnrecognizedCalledAETitle = errors.New("dicom: unrecognized called AE Title")

	// ErrServerClosed is returned to indicate that the server has been closed or shutdown
	ErrServerClosed = errors.New("dicom: server shutdown or closed")
)

// A Handler responds to a DICOM request.
//
// ServeDICOM should write responses to the writer and then return.
// Returning signals that the request is finished.
type Handler struct {
	Capabilities []*Capability
}

// A Server is a DICOM server.  In DICOM parlance, it is often referrned
// to erroneusly as an SCP or more accurately as an Acceptor.
//
// A Server's zero value is a usable server.
type Server struct {

	// Addr optionally specifies the TCP address for the server to listen on
	// in the form host:port.  If empty, "localhost:4104" is used.
	Addr string

	// AETitle optionallly specifies one or more AE Titles for the server
	// to recognized as Called AE Titles.  If nil, a single AE Title
	// "DCMSRV" is used.
	AETitles []string

	// Handlers are the handlers used to respond to DICOM requests.
	// If nil, a set of default handlers are loaded that will respond
	// to requests without provideing any meaninful information.
	Handlers []Handler

	// ReadTimeout is the maximum time allowed for reading a request
	// If zero, there is no read timeout.
	ReadTimeout time.Duration

	// WriteTimeout is the maximum time allows for writing a response
	// If zero, there is no write timeout.
	WriteTimeout time.Duration
}

// Close immediately closes all listens and any connections
// For a graceful shutdown, use Shutdown.
// Close returns any errors returned from closing the listeners.
func (server *Server) Close() error {
	return fmt.Errorf("Server.Close(): not implemented")
}

// Shutdown shuts down the listeners and waits for connections to close.
func (server *Server) Shutdown() error {
	return fmt.Errorf("Server.Shutdown(): not implemented")
}

// ListenAndServe listens on the TCP network address and then calls
// Serve to handle requests on incoming connections.
// ListenAndServe always returns a non-nil error.
// After Shutdown or Close, the returned error is ErrServerClosed.
func (server *Server) ListenAndServe() error {

	addr := server.Addr
	if addr == "" {
		addr = "localhost:4104"
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return server.Serve(listener)
}

// Serve accepts incoming connections on the listener, creating a new
// goroutne for each connection.  The goroutines read requests and then
// call handers to reply to them.
// Serve always returns a non-nil error and closes the listener.
func (server *Server) Serve(listener net.Listener) error {

	// listen for connections
	for {

		// wait for connection
		fmt.Printf("waiting for connection on %v\n", listener.Addr())
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error occured while waiting for connection on %v, error is %v\n", listener.Addr(), err)
			continue
		}

		fmt.Printf("accepted connection on %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())

		// handle the connection, eventually as a goroutine
		server.serve(conn)

		// break for now, so that we don't keep the port open
		break
	}

	err := listener.Close()
	if err != nil {
		fmt.Printf("error occured while closing listener on %v, error is %v\n", listener.Addr(), err)
	}

	// let the caller know that the server is now closed
	return ErrServerClosed
}

// serve will attempt to negotiate an association based on the capabilities
// of the handlers that have been added to this server
// then, serve will wait for requests and will call the appropriate handler
// to handle the request
func (server *Server) serve(conn net.Conn) {

	// gather all the capabilities of all the handlers
	capabilities := collectCapabilities(server.Handlers)
	fmt.Printf("capabilities is %v\n", capabilities)

	fmt.Printf("Server.serve(): not implemented")
}

func collectCapabilities(handlers []Handler) []*Capability {
	capabilities := make([]*Capability, 0, 5)
	for _, handler := range handlers {
		for _, capability := range handler.Capabilities {
			capabilities = append(capabilities, capability)
		}
	}
	return capabilities
}

// ListenAndServe listens on the TCP network address addr and then calls
// Serve with handlers to handle DICOM requests on incoming connections.
func ListenAndServe(addr string, handlers []Handler) error {
	server := &Server{Addr: addr, Handlers: handlers}
	return server.ListenAndServe()
}
