package ruby

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var errNotFound = errors.New("no license info found")

// ScanLicense returns result of fetch https://rubygems.org
func ScanLicense(name, version string) (string, float64, error) {
	b, err := fetchJson(name, version)
	if err != nil {
		return "unknown", 0, err
	}
	result, confidence, err := parseResponce(b)
	if err != nil {
		return "unknown", 0, err
	}
	return result, confidence, nil
}

func fetchJson(name, version string) ([]byte, error) {
	ref := "https://rubygems.org/api/v2/rubygems/%s/versions/%s.json"
	resp, err := http.Get(fmt.Sprintf(ref, name, version))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errNotFound
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func parseResponce(b []byte) (string, float64, error) {
	license := struct {
		Licenses []string `json:"licenses"`
	}{}
	json.Unmarshal(b, &license)

	if license.Licenses == nil {
		return "", 0, errNotFound
	}
	return strings.Join(license.Licenses, ","), 1, nil
}
