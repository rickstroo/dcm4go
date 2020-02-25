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

// RequestAssoc requests an association
func (ae *AE) RequestAssoc(conn net.Conn, remoteAE *AE, capabilities []*PresContext, opts *AssocOpts) (*RequestorAssoc, error) {
	return requestAssoc(conn, ae, remoteAE, capabilities, opts)
}

// AcceptAssoc waits for an association request
func (ae *AE) AcceptAssoc(conn net.Conn, capabilities []*PresContext, opts *AssocOpts) (*AcceptorAssoc, error) {
	return acceptAssoc(conn, ae, capabilities)
}
