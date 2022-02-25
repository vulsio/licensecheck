package rust

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ScanLicense returns result of fetch https://crates.io
func ScanLicense(name, version string) (string, float64, error) {
	ref := "https://crates.io/api/v1/crates/%v/%v"
	resp, err := http.Get(fmt.Sprintf(ref, name, version))
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "unknown", 0, nil
	}

	license := struct {
		Version struct {
			License string `json:"license"`
		} `json:"version"`
	}{}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "unknown", 0, err
	}

	json.Unmarshal(b, &license)

	if license.Version.License == "" {
		return "unknown", 0, err
	}
	return license.Version.License, 1, err
}
