package dcm4go

import "fmt"

// AE represents an application entity
type AE struct {
	aeTitle string
}

// String returns a string representation of an ae
func (ae *AE) String() string {
	return fmt.Sprintf(
		"{aeTitle:%q}",
		ae.aeTitle,
	)
}

// AETitle returns the AE title of the AE
func (ae *AE) AETitle() string {
	return ae.aeTitle
}
