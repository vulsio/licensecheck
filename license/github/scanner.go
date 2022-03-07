package github

import (
	"fmt"

	"github.com/google/licenseclassifier"
	"github.com/vuls-saas/licensecheck/license/shared"
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
)

const ref = "https://raw.githubusercontent.com/%s/%s/%s"

type Scanner struct {
	Crawler shared.Crawler
}

// ScanLicense returns result of Scan on github.com blob objects
// fetches LICENSE, README,md, or COPYING of master/main branch, and returns license if confidence is over 90%
// detection logic is depends on github.com/google/licenseclassifier
func (s *Scanner) ScanLicense(name, _ string) (string, float64, error) {
	if s.Crawler == nil {
		s.Crawler = &shared.DefaultCrawler{}
	}
	classifier, err := licenseclassifier.New(0.9)
	if err != nil {
		return "unknown", 0, err
	}
	// NOTE: avoid to use GitHub REST API, we want to use it without any tokens, without worrying about the rate limit.
	for _, branch := range branches {
		for _, content := range contents {
			b, err := s.Crawler.Crawl(fmt.Sprintf(ref, name, branch, content))
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

func parseResponce(b []byte, c *licenseclassifier.License) (string, float64, error) {
	matches := c.MultipleMatch(string(b), true)
	if len(matches) == 0 {
		return "", 0, shared.ErrNotFound
	}
	return matches[0].Name, matches[0].Confidence, nil
}
