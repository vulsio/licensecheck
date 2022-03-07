package licensecheck

import (
	"errors"

	"github.com/vuls-saas/licensecheck/core/github"
	"github.com/vuls-saas/licensecheck/core/golicense"
	"github.com/vuls-saas/licensecheck/core/java"
	"github.com/vuls-saas/licensecheck/core/nodejs"
	"github.com/vuls-saas/licensecheck/core/python"
	"github.com/vuls-saas/licensecheck/core/ruby"
	"github.com/vuls-saas/licensecheck/core/rust"
	"github.com/vuls-saas/licensecheck/shared"
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

// Scanner is struct to scan license info
// Crawler is exported to modify or make it easy to test by mock
type Scanner struct {
	Crawler shared.Crawler
}

// ErrUnKnownScanType is error message when unexpected scanType is given
var ErrUnKnownScanType = errors.New("unknown scanType is specified")

// Scan returns result of license check(string) and confidence(float64)
func (s *Scanner) Scan(name, version string, scanType int) (string, float64, error) {
	var sc shared.Scanner
	switch scanType {
	case Java:
		sc = &java.Scanner{Crawler: s.Crawler}
	case Ruby:
		sc = &ruby.Scanner{Crawler: s.Crawler}
	case Python:
		sc = &python.Scanner{Crawler: s.Crawler}
	case Nodejs:
		sc = &nodejs.Scanner{Crawler: s.Crawler}
	case Go:
		sc = &golicense.Scanner{Crawler: s.Crawler}
	case Rust:
		sc = &rust.Scanner{Crawler: s.Crawler}
	case GitHub:
		sc = &github.Scanner{Crawler: s.Crawler}
	default:
		return "unknown", 0, ErrUnKnownScanType
	}
	return sc.ScanLicense(name, version)
}
