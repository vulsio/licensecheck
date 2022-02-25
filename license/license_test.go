package license

import (
	"errors"
	"testing"
)

func TestScanInvalidArguments(t *testing.T) {
	result, confidence, err := Scan("", "", 999)
	if !errors.Is(err, ErrUnKnownScanType) {
		t.Error(err)
	}
	if result != "unknown" {
		t.Errorf("want: unknown, got: %s", result)
	}
	if confidence != 0 {
		t.Errorf("want: 0, got: %f", confidence)
	}

}
