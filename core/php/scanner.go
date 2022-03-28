package php

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vulsio/licensecheck/shared"
)

const ref = "https://packagist.org/packages/%s.json"

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
	b, err := s.Crawler.Crawl(fmt.Sprintf(ref, name))
	if err != nil {
		return "unknown", 0, err
	}
	result, confidence, err := parseResponce(b, version)
	if err != nil {
		return "unknown", 0, err
	}
	return result, confidence, nil
}

func parseResponce(b []byte, version string) (string, float64, error) {
	license := struct {
		Package struct {
			Versions map[string]struct {
				License []string `json:"license"`
			} `json:"versions"`
		} `json:"package"`
	}{}
	if err := json.Unmarshal(b, &license); err != nil {
		return "", 0, shared.ErrNotFound
	}
	if version == "" {
		if pkg, ok := license.Package.Versions["dev-main"]; ok {
			return joinedResult(pkg.License)
		}
		if pkg, ok := license.Package.Versions["dev-master"]; ok {
			return joinedResult(pkg.License)
		}
	} else {
		if pkg, ok := license.Package.Versions[version]; ok {
			return joinedResult(pkg.License)
		}
	}
	return "", 0, shared.ErrNotFound
}

func joinedResult(licenses []string) (string, float64, error) {
	s := strings.Join(licenses, ",")
	if s == "" {
		return "", 0, shared.ErrNotFound
	}
	return s, 1, nil
}
