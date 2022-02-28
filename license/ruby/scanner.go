package ruby

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vuls-saas/license-scanner/license/shared"
)

const ref = "https://rubygems.org/api/v2/rubygems/%s/versions/%s.json"

// ScanLicense returns result of fetch https://rubygems.org
func ScanLicense(name, version string) (string, float64, error) {
	b, err := shared.Crawl(fmt.Sprintf(ref, name, version))
	if err != nil {
		return "unknown", 0, err
	}
	result, confidence, err := parseResponce(b)
	if err != nil {
		return "unknown", 0, err
	}
	return result, confidence, nil
}

func parseResponce(b []byte) (string, float64, error) {
	license := struct {
		Licenses []string `json:"licenses"`
	}{}
	json.Unmarshal(b, &license)

	if license.Licenses == nil {
		return "", 0, shared.ErrNotFound
	}
	return strings.Join(license.Licenses, ","), 1, nil
}
