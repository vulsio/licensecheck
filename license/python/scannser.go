package python

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ScanLicense returns result of fetch https://pypi.org
// version is not required (if version is given, the result will be more rigorous)
func ScanLicense(name, version string) (string, float64, error) {
	ref := "https://pypi.org/pypi"
	if name != "" {
		ref = fmt.Sprintf("%s/%s", ref, name)
	}
	if version != "" {
		ref = fmt.Sprintf("%s/%s", ref, version)
	}
	resp, err := http.Get(fmt.Sprintf("%s/%s", ref, "json"))
	if err != nil {
		return "unknown", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "unknown", 0, nil
	}

	license := struct {
		Info struct {
			License string `json:"license"`
		} `json:"info"`
	}{}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "unknown", 0, err
	}

	json.Unmarshal(b, &license)

	if license.Info.License == "" {
		return "unknown", 0, err
	}
	return license.Info.License, 1, err
}
