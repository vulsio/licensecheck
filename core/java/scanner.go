package java

import (
	"encoding/xml"
	"fmt"

	"github.com/google/licenseclassifier"
	"github.com/vulsio/licensecheck/shared"
)

const ref = "https://repo1.maven.org/maven2"

// Project is struct to unmarshal pom.xml of java project
type Project struct {
	Licenses struct {
		License struct {
			Name string `xml:"name"`
		} `xml:"license"`
	} `xml:"licenses"`
}

// Scanner is struct to scan license info
// Crawler is exported to modify or make it easy to test by mock
type Scanner struct {
	Crawler shared.Crawler
}

// ScanLicense returns license fetched from repo1.maven.org/maven2, returns license if confidence is over 90%
// example https://repo1.maven.org/maven2/org/apache/logging/log4j/log4j-core/2.17.1/log4j-core-2.17.1.pom
// name is required in a format like org/apache/logging/log4j/log4j-core/2.17.1/log4j-core-2.17.1.pom
// If structure of pom is not expected, Detection logic is depends on github.com/google/licenseclassifier
func (s *Scanner) ScanLicense(name, _ string) (string, float64, error) {
	if s.Crawler == nil {
		s.Crawler = &shared.DefaultCrawler{}
	}
	b, err := s.Crawler.Crawl(fmt.Sprintf("%s/%s", ref, name))
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
	var p Project
	if err := xml.Unmarshal(b, &p); err != nil {
		return "", 0, err
	}
	if p.Licenses.License.Name != "" {
		return p.Licenses.License.Name, 1, nil
	}

	c, err := licenseclassifier.New(0.7)
	if err != nil {
		return "", 0, err
	}
	matches := c.MultipleMatch(string(b), true)
	if len(matches) == 0 {
		return "", 0, shared.ErrNotFound
	}
	return matches[0].Name, matches[0].Confidence, nil
}
