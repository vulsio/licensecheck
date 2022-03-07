package rust

import (
	"encoding/json"
	"fmt"

	"github.com/vulsio/licensecheck/shared"
)

const ref = "https://crates.io/api/v1/crates/%v/%v"

// Scanner is struct to scan license info
// Crawler is exported to modify or make it easy to test by mock
type Scanner struct {
	Crawler shared.Crawler
}

// ScanLicense returns result of fetch https://crates.io
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
		Version struct {
			License string `json:"license"`
		} `json:"version"`
	}{}
	if err := json.Unmarshal(b, &license); err != nil {
		return "", 0, shared.ErrNotFound
	}

	if license.Version.License == "" {
		return "", 0, shared.ErrNotFound
	}
	return license.Version.License, 1, nil
}
