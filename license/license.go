package license

import (
	"errors"
	"strings"

	"github.com/vuls-saas/license-scanner/license/github"
	"github.com/vuls-saas/license-scanner/license/java"
	"github.com/vuls-saas/license-scanner/license/nodejs"
	"github.com/vuls-saas/license-scanner/license/python"
	"github.com/vuls-saas/license-scanner/license/ruby"
	"github.com/vuls-saas/license-scanner/license/rust"
)

const (
	Java = iota
	Ruby
	Python
	Nodejs
	Go
	Rust
	GitHub
)

// ErrUnKnownScanType is error message when unexpected scanType is given
var ErrUnKnownScanType = errors.New("unknown scanType is specified")

// Scan returns result of license check(string) and confidence(float64)
func Scan(name, version string, scanType int) (string, float64, error) {
	switch scanType {
	case Java:
		return new(java.Scanner).ScanLicense(name, version)
	case Ruby:
		return new(ruby.Scanner).ScanLicense(name, version)
	case Python:
		return new(python.Scanner).ScanLicense(name, version)
	case Nodejs:
		return new(nodejs.Scanner).ScanLicense(name, version)
	case Go:
		return new(github.Scanner).ScanLicense(strings.Replace(name, "github.com/", "", 1), "")
	case Rust:
		return new(rust.Scanner).ScanLicense(name, version)
	case GitHub:
		return new(github.Scanner).ScanLicense(name, version)
	default:
		return "unknown", 0, ErrUnKnownScanType
	}
}
