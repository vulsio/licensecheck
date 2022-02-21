package java

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/licenseclassifier"
)

// ScanLicense returns license fetched from repo1.maven.org/maven2, returns license if confidence is over 90%
// example https://repo1.maven.org/maven2/org/apache/logging/log4j/log4j-core/2.17.1/log4j-core-2.17.1.pom
// name is required in a format like org/apache/logging/log4j/log4j-core/2.17.1/log4j-core-2.17.1.pom
// detection logic is depends on github.com/google/licenseclassifier
func ScanLicense(name string) (string, float64, error) {
	c, err := licenseclassifier.New(0.9)
	if err != nil {
		return "unknown", 0, err
	}
	ref := "https://repo1.maven.org/maven2"

	ref = fmt.Sprintf("%s/%s", ref, name)
	resp, err := http.Get(ref)
	if err != nil {
		return "unknown", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "unknown", 0, nil
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "unknown", 0, err
	}

	matches := c.MultipleMatch(string(b), true)
	if len(matches) == 0 {
		return "unknown", 0, nil
	}

	return matches[0].Name, matches[0].Confidence, nil
}
