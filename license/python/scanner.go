package python

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var errNotFound = errors.New("no license info found")

// ScanLicense returns result of fetch https://pypi.org
// version is not required (if version is given, the result will be more rigorous)
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
	ref := "https://pypi.org/pypi"
	if name != "" {
		ref = fmt.Sprintf("%s/%s", ref, name)
	}
	if version != "" {
		ref = fmt.Sprintf("%s/%s", ref, version)
	}
	resp, err := http.Get(fmt.Sprintf("%s/%s", ref, "json"))
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
		Info struct {
			License string `json:"license"`
		} `json:"info"`
	}{}
	json.Unmarshal(b, &license)
	if license.Info.License == "" {
		return "", 0, errNotFound
	}
	return license.Info.License, 1, nil
}
