package github

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/licenseclassifier"
)

var (
	branches = []string{
		"master",
		"main",
	}
	contents = []string{
		"LICENSE",
		"README.md",
		"COPYING",
	}
	ref = "https://raw.githubusercontent.com/%s/%s/%s"
)

// ScanLicense returns result of Scan on github.com blob objects
// fetches LICENSE, README,md, or COPYING of master/main branch, and returns license if confidence is over 90%
// detection logic is depends on github.com/google/licenseclassifier
func ScanLicense(name string) (string, float64, error) {
	c, err := licenseclassifier.New(0.9)
	if err != nil {
		return "unknown", 0, err
	}
	for _, branch := range branches {
		for _, content := range contents {
			resp, err := http.Get(fmt.Sprintf(ref, name, branch, content))
			if err != nil {
				continue
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				continue
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			matches := c.MultipleMatch(string(b), true)
			if len(matches) == 0 {
				continue
			}
			return matches[0].Name, matches[0].Confidence, nil
		}
	}
	return "unknown", 0, nil
}
