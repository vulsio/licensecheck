package shared

import (
	"io/ioutil"
	"net/http"
)

type Crawler interface {
	Crawl(url string) ([]byte, error)
}

type DefaultCrawler struct{}

// Crawl executes GET request with options like User-Agent
func (d *DefaultCrawler) Crawl(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// It is recommended to add User-Agent
	// https://crates.io/policies
	req.Header.Set("User-Agent", "licensecheck")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrNotFound
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
