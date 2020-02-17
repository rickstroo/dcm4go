// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
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
	aeTitle string
	host    string
	port    string
}

// NewAE parses an address and returns a filled in AE
func NewAE(addr string) *AE {
	aeTitle, host, port := parseAddr(addr)
	return &AE{
		aeTitle: aeTitle,
		host:    host,
		port:    port,
	}
}

// parseAddr parses an address of the form 'ae@host:port' and returns the
// 'ae' and 'host:port' parts separately
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

// AETitle returns an AE's AE Title
func (ae *AE) AETitle() string {
	return ae.aeTitle
}

// Host returns an AE's Host, if present
func (ae *AE) Host() string {
	return ae.host
}

// Port returns an AE's Port, if present
func (ae *AE) Port() string {
	return ae.port
}

// Been thinking about how to make the AE the starting point for all
// applications, yet separate out the Requestor and Acceptor methods.
// Perhaps this is a good compromise.

// RequestAssoc requests an association and returns a Requestor
func (ae *AE) RequestAssoc(remoteAE *AE, capabilities []*Capability, opts *AssocOpts) (*Requestor, error) {
	requestor, err := requestAssoc(ae, remoteAE, capabilities, opts)
	if err != nil {
		return nil, err
	}
	return requestor, nil
}

// AcceptAssoc waits for an association and returns an Acceptor
func (ae *AE) AcceptAssoc(capabilities []*Capability, opts *AssocOpts) (*Acceptor, error) {
	return nil, fmt.Errorf("AE.AcceptAssoc(): not implemented")
}
