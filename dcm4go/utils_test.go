package dcm4go

import (
	"errors"
	"fmt"
)

func testErrEquals(a, b error) error {
	if errors.Is(a, b) {
		return nil
	}
	return fmt.Errorf("expected %v, found %v", a, b)
}

func testByteEquals(a, b byte) error {
	if a != b {
		return fmt.Errorf("expected 0x%02X, was 0x%02X", a, b)
	}
	return nil
}

func testShortEquals(a, b uint16) error {
	if a != b {
		return fmt.Errorf("expected 0x%04X, was 0x%04X", a, b)
	}
	return nil
}

func testLongEquals(a, b uint32) error {
	if a != b {
		return fmt.Errorf("expected 0x%08X, was 0x%08X", a, b)
	}
	return nil
}

func testStringEquals(a, b string) error {
	if a != b {
		return fmt.Errorf("expected '%s', was '%s'", a, b)
	}
	return nil
}
