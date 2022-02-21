package nodejs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ScanLicense returns result of fetch https://registry.npmjs.org
func ScanLicense(name, version string) (string, float64, error) {
	ref := "https://registry.npmjs.org/%s/%s"
	resp, err := http.Get(fmt.Sprintf(ref, name, version))
	if err != nil {
		return "unknown", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "unknown", 0, nil
	}

	license := struct {
		License string `json:"license"`
	}{}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "unknown", 0, err
	}

	json.Unmarshal(b, &license)

	if license.License == "" {
		return "unknown", 0, err
	}
	return license.License, 1, err
}
