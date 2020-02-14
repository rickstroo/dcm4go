package dcm4go

import (
	"io"
)

// A Parser can parse a DICOM object from a reader
type Parser struct {
	Opts *ParserOpts
}

// ParserOpts impact the behaviour of a Parser
type ParserOpts struct {
	BulkDataThreshold uint32 // a zero value means no bulk data threshold
}

// Parse parses a DICOM object from a reader and returns the object
func (parser *Parser) Parse(reader io.Reader) (*Object, error) {
	object, err := ReadFile(reader, parser.Opts.BulkDataThreshold)
	if err != nil {
		return nil, err
	}
	return object, nil
}

// Parse parses a DICOM object from a reader using a default set of options.
// To gain more control over the parsing, the user should create a Parser
// with the desired ParserOpts.
func Parse(reader io.Reader) (*Object, error) {
	opts := &ParserOpts{}
	parser := &Parser{Opts: opts}
	return parser.Parse(reader)
}
