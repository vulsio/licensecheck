package rust

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var errNotFound = errors.New("no license info found")

// ScanLicense returns result of fetch https://crates.io
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
	ref := "https://crates.io/api/v1/crates/%v/%v"
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
		Version struct {
			License string `json:"license"`
		} `json:"version"`
	}{}
	json.Unmarshal(b, &license)

	if license.Version.License == "" {
		return "", 0, errNotFound
	}
	return license.Version.License, 1, nil
}
