package shared

import "errors"

// ErrNotFound is error message when license info is not found
var ErrNotFound = errors.New("no license info found")

type Scanner interface {
	ScanLicense(name, version string) (string, float64, error)
}
