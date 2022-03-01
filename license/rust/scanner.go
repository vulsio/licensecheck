package rust

import (
	"encoding/json"
	"fmt"

	"github.com/vuls-saas/license-scanner/license/shared"
)

const ref = "https://crates.io/api/v1/crates/%v/%v"

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
	json.Unmarshal(b, &license)

	if license.Version.License == "" {
		return "", 0, shared.ErrNotFound
	}
	return license.Version.License, 1, nil
}
