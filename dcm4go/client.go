// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"time"
)

// A Client is a DICOM client.  In DICOM parlance, it is often referrned
// to erroneusly as an SCU or more accurately as a Requestor.
//
// A Client's zero value is a usable client.
type Client struct {

	// ConnectTimeout is the maximum time allowed for connection to a server
	// If zero, there is no connect timeout.
	ConnectTimeout time.Duration

	// ReadTimeout is the maximum time allowed for reading a response
	// If zero, there is no read timeout.
	ReadTimeout time.Duration

	// WriteTimeout is the maximum time allows for writing a request
	// If zero, there is no write timeout.
	WriteTimeout time.Duration
}

// Verify sends  DICOM verifiction request from a client to a server
func (client *Client) Verify(addr string) error {
	return fmt.Errorf("Client.Send(): not implemented")
}

// Verify creates a DICOM client and sends a verification request
func Verify(addr string) error {
	client := &Client{}
	return client.Verify(addr)
}

// Send sends a DICOM store request from a client to a server.
func (client *Client) Send(addr string, request *Request) error {
	return fmt.Errorf("Client.Send(): not implemented")
}

// Send creates a DICOM client and sends a store request
func Send(addr string, request *Request) error {
	client := &Client{}
	return client.Send(addr, request)
}
