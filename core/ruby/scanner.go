package ruby

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vulsio/licensecheck/shared"
)

const ref = "https://rubygems.org/api/v2/rubygems/%s/versions/%s.json"

// Scanner is struct to scan license info
// Crawler is exported to modify or make it easy to test by mock
type Scanner struct {
	Crawler shared.Crawler
}

// ScanLicense returns result of fetch https://rubygems.org
func (s *Scanner) ScanLicense(name, version string) (string, float64, error) {
	if s.Crawler == nil {
		s.Crawler = &shared.DefaultCrawler{}
	}
	b, err := s.Crawler.Crawl(fmt.Sprintf(ref, name, version))
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
	license := struct {
		Licenses []string `json:"licenses"`
	}{}
	if err := json.Unmarshal(b, &license); err != nil {
		return "", 0, shared.ErrNotFound
	}

	if len(license.Licenses) == 0 {
		return "", 0, shared.ErrNotFound
	}
	return strings.Join(license.Licenses, ","), 1, nil
}
