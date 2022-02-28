package java

import (
	"encoding/xml"
	"fmt"

	"github.com/google/licenseclassifier"
	"github.com/vuls-saas/license-scanner/license/shared"
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

// ScanLicense returns license fetched from repo1.maven.org/maven2, returns license if confidence is over 90%
// example https://repo1.maven.org/maven2/org/apache/logging/log4j/log4j-core/2.17.1/log4j-core-2.17.1.pom
// name is required in a format like org/apache/logging/log4j/log4j-core/2.17.1/log4j-core-2.17.1.pom
// If structure of pom is not expected, Detection logic is depends on github.com/google/licenseclassifier
func ScanLicense(name string) (string, float64, error) {
	b, err := shared.Crawl(fmt.Sprintf("%s/%s", ref, name))
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
	xml.Unmarshal(b, &p)
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
