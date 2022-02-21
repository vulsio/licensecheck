package ruby

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ScanLicense returns result of fetch https://rubygems.org
func ScanLicense(name, version string) (string, float64, error) {
	ref := "https://rubygems.org/api/v2/rubygems/%s/versions/%s.json"
	resp, err := http.Get(fmt.Sprintf(ref, name, version))
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "unknown", 0, nil
	}

	license := struct {
		Licenses []string `json:"licenses"`
	}{}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	json.Unmarshal(b, &license)

	if license.Licenses == nil {
		return "unknown", 0, err
	}
	return strings.Join(license.Licenses, ","), 1, err
}
