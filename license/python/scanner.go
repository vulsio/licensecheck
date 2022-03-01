package python

import (
	"encoding/json"
	"fmt"

	"github.com/vuls-saas/license-scanner/license/shared"
)

const ref = "https://pypi.org/pypi"

type Scanner struct {
	Crawler shared.Crawler
}

// ScanLicense returns result of fetch https://pypi.org
// version is not required (if version is given, the result will be more rigorous)
func (s *Scanner) ScanLicense(name, version string) (string, float64, error) {
	if s.Crawler == nil {
		s.Crawler = &shared.DefaultCrawler{}
	}
	b, err := s.Crawler.Crawl(fmt.Sprintf("%s/%s/%s", ref, name, version))
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
		Info struct {
			License string `json:"license"`
		} `json:"info"`
	}{}
	json.Unmarshal(b, &license)
	if license.Info.License == "" {
		return "", 0, shared.ErrNotFound
	}
	return license.Info.License, 1, nil
}
