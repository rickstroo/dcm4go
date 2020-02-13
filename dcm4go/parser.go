package dcm4go

import (
	"fmt"
	"io"
)

// A Parser can parse a DICOM object from a reader
type Parser struct {
	opts *ParserOpts
}

// ParserOpts impact the behaviour of a Parser
type ParserOpts struct {
}

// Parse parses a DICOM object from a reader and returns the object
func (parser *Parser) Parse(reader io.Reader) (*Object, error) {
	return nil, fmt.Errorf("Parser.Parse(), not implemented")
}

// Parse parses a DICOM object from a read using a default set of options
func Parse(reader io.Reader) (*Object, error) {
	parserOpts := &ParserOpts{}
	parser := &Parser{opts: parserOpts}
	return parser.Parse(reader)
}
