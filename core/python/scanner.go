package python

import (
	"encoding/json"
	"net/url"
	"path"

	"github.com/vulsio/licensecheck/shared"
)

const ref = "https://pypi.org/pypi"

// Scanner is struct to scan license info
// Crawler is exported to modify or make it easy to test by mock
type Scanner struct {
	Crawler shared.Crawler
}

// ScanLicense returns result of fetch https://pypi.org
// version is not required (if version is given, the result will be more rigorous)
func (s *Scanner) ScanLicense(name, version string) (string, float64, error) {
	if s.Crawler == nil {
		s.Crawler = &shared.DefaultCrawler{}
	}
	u, err := url.Parse(ref)
	if err != nil {
		return "unknown", 0, err
	}
	u.Path = path.Join(u.Path, name, version, "json")
	b, err := s.Crawler.Crawl(u.String())
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
	if err := json.Unmarshal(b, &license); err != nil {
		return "", 0, shared.ErrNotFound
	}
	if license.Info.License == "" {
		return "", 0, shared.ErrNotFound
	}
	return license.Info.License, 1, nil
}
