package license

import (
	"errors"
	"strings"

	"github.com/vuls-saas/license-scanner/license/github"
	"github.com/vuls-saas/license-scanner/license/java"
	"github.com/vuls-saas/license-scanner/license/nodejs"
	"github.com/vuls-saas/license-scanner/license/python"
	"github.com/vuls-saas/license-scanner/license/ruby"
)

const (
	Java = iota
	Ruby
	Python
	Nodejs
	Go
	GitHub
)

// Scan returns result of license check(string) and confidence(float64)
func Scan(name, version string, scanType int) (string, float64, error) {
	switch scanType {
	case Java:
		return java.ScanLicense(name)
	case Ruby:
		return ruby.ScanLicense(name, version)
	case Python:
		return python.ScanLicense(name, version)
	case Nodejs:
		return nodejs.ScanLicense(name, version)
	case Go:
		return github.ScanLicense(strings.Replace(name, "github.com/", "", 1))
	case GitHub:
		return github.ScanLicense(name)
	default:
		return "unknown", 0, errors.New("unknown scanType is specified")
	}
}