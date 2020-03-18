// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"net"
	"strings"
)

// I'm starting to think that having an AE as the starting point for all the
// APIs in the library makes sense.  It is AEs that application developers
// are actually building.  The APIs below allow for AEs to serve as SCUs,
// SCPs, and even SCUs and SCPs concurrently.
//
// Actually, there are a couple of types of AEs, Requestor AEs and
// Acceptor AEs.  A Requestor AE initiates an association while an Acceptor
// AE provides an association.  It seems to make sense to divide up the
// methods for an AE for each.  There may be some methods that are
// common to both AEs.  Perhaps they should be encapsulated as an AE
// method and then the Requestor AE and Acceptor AE can be derived from
// an AE.
//
// In the end, decided to stick with an AE, and have different types of
// associations, one used for acceptors and one used for requestors.  That
// makes more sense because whether one is an acceptor or requestor is a
// function of the association.  An AE can be a requestor or an acceptor
// at different times.

// AE represents an application entity
type AE struct {
	Title string
	Host  string
	Port  string
}

// NewAE parses an address and returns a filled in AE
func NewAE(addr string) *AE {
	title, host, port := parseAddr(addr)
	return &AE{
		Title: title,
		Host:  host,
		Port:  port,
	}
}

// parseAddr parses an address of the form 'ae@host:port' and returns the
// 'ae', 'host' and 'port' parts separately
func parseAddr(addr string) (string, string, string) {
	s := strings.Split(addr, "@")
	if len(s) == 1 {
		return s[0], "", ""
	}
	t := strings.Split(s[1], ":")
	if len(t) == 1 {
		return s[0], t[0], ""
	}
	return s[0], t[0], t[1]
}

// RequestAssoc requests an association from the local AE to the remote AE.
func (ae *AE) RequestAssoc(conn net.Conn, remoteAE *AE, capabilities *Capabilities, opts *AssocOpts) (*Assoc, error) {
	return RequestAssoc(conn, ae.Title, remoteAE.Title, capabilities, opts)
}

// AcceptAssoc waits for an association request
func (ae *AE) AcceptAssoc(conn net.Conn, capabilities *Capabilities, opts *AssocOpts) (*Assoc, error) {
	return AcceptAssoc(conn, ae.Title, capabilities)
}
