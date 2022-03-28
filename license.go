package licensecheck

import (
	"errors"

	"github.com/vulsio/licensecheck/core/github"
	"github.com/vulsio/licensecheck/core/golicense"
	"github.com/vulsio/licensecheck/core/java"
	"github.com/vulsio/licensecheck/core/nodejs"
	"github.com/vulsio/licensecheck/core/php"
	"github.com/vulsio/licensecheck/core/python"
	"github.com/vulsio/licensecheck/core/ruby"
	"github.com/vulsio/licensecheck/core/rust"
	"github.com/vulsio/licensecheck/shared"
)

const (
	Java = iota
	PHP
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
	case PHP:
		sc = &php.Scanner{Crawler: s.Crawler}
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
