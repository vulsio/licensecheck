package github

import (
	"errors"
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
	ref         = "https://raw.githubusercontent.com/%s/%s/%s"
	errNotFound = errors.New("no license info found")
)

// ScanLicense returns result of Scan on github.com blob objects
// fetches LICENSE, README,md, or COPYING of master/main branch, and returns license if confidence is over 90%
// detection logic is depends on github.com/google/licenseclassifier
func ScanLicense(name string) (string, float64, error) {
	classifier, err := licenseclassifier.New(0.9)
	if err != nil {
		return "unknown", 0, err
	}
	for _, branch := range branches {
		for _, content := range contents {
			b, err := fetchBlob(name, branch, content)
			if err != nil {
				continue
			}
			result, confidence, err := parseResponce(b, classifier)
			if err != nil {
				continue
			}
			return result, confidence, nil
		}
	}
	return "unknown", 0, nil
}

func fetchBlob(name, branch, content string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(ref, name, branch, content))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errNotFound
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func parseResponce(b []byte, c *licenseclassifier.License) (string, float64, error) {
	matches := c.MultipleMatch(string(b), true)
	if len(matches) == 0 {
		return "", 0, errNotFound
	}
	return matches[0].Name, matches[0].Confidence, nil
}
