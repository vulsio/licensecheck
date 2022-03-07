package golicense

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/vulsio/licensecheck/shared"
)

const (
	ref            = "https://pkg.go.dev/%s?tab=licenses"
	refWithVersion = "https://pkg.go.dev/%s@%s?tab=licenses"
)

// Scanner is struct to scan license info
// Crawler is exported to modify or make it easy to test by mock
type Scanner struct {
	Crawler shared.Crawler
}

var versionRegexp = regexp.MustCompile(`<div id="#lic-0">.*</div>`)

// ScanLicense returns result of fetch https://registry.npmjs.org
func (s *Scanner) ScanLicense(name, version string) (string, float64, error) {
	if s.Crawler == nil {
		s.Crawler = &shared.DefaultCrawler{}
	}
	url := fmt.Sprintf(ref, name)
	if version != "" {
		url = fmt.Sprintf(refWithVersion, name, version)
	}
	b, err := s.Crawler.Crawl(url)
	if err != nil {
		return "unknown", 0, err
	}
	result, confidence, err := parseResponce(b)
	if err != nil {
		return "unknown", 0, err
	}
	return result, confidence, nil
}

func parseResponce(b []byte) (string, float64, error) {
	v := string(versionRegexp.Find(b))
	v = strings.TrimPrefix(v, `<div id="#lic-0">`)
	v = strings.TrimSuffix(v, `</div>`)
	if v == "" {
		return "", 0, shared.ErrNotFound
	}
	return v, 1, nil
}
