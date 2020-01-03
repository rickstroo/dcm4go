package file

import (
	"fmt"

	"github.com/rickstroo/dcm4go/core"
)

// ReadFile reads a DICOM file
func ReadFile(path string) (*core.Object, error) {
	fmt.Printf("about to read file '%s'\n", path)
	return core.NewObject(), nil
}
