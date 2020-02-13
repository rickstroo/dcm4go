package dcm4go

import (
	"io"
)

// A Parser can parse a DICOM object from a reader
type Parser struct {
	opts *ParserOpts
}

// ParserOpts impact the behaviour of a Parser
type ParserOpts struct {
	BulkDataThreshold uint32
}

// Parse parses a DICOM object from a reader and returns the object
func (parser *Parser) Parse(reader io.Reader) (*Object, error) {
	object, err := ReadFile(reader, parser.opts.BulkDataThreshold)
	if err != nil {
		return nil, err
	}
	return object, nil
}

// ParseWithOpts parses a DICOM object from a reader with user provided options
func ParseWithOpts(reader io.Reader, opts *ParserOpts) (*Object, error) {
	parser := &Parser{opts: opts}
	return parser.Parse(reader)
}

// Parse parses a DICOM object from a reader using a default set of options
func Parse(reader io.Reader) (*Object, error) {
	opts := &ParserOpts{BulkDataThreshold: 1024}
	return ParseWithOpts(reader, opts)
}
