package github

import (
	"io/ioutil"
	"testing"

	"github.com/google/licenseclassifier"
)

func TestParseResponce(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		result     string
		confidence float64
	}{
		{
			name:       "MIT",
			in:         "../../testdata/github/MIT_sample.txt",
			result:     "MIT",
			confidence: 1,
		},
		{
			name:       "Apache-2.0",
			in:         "../../testdata/github/Apache-2.0_sample.txt",
			result:     "Apache-2.0",
			confidence: 1,
		},
		{
			name:       "GPL-3.0",
			in:         "../../testdata/github/GPL-3.0_sample.txt",
			result:     "GPL-3.0",
			confidence: 1,
		},
		{
			name:       "MIT in README",
			in:         "../../testdata/github/README_sample.md",
			result:     "MIT",
			confidence: 1,
		},
	}
	for _, tt := range tests {
		classifier, err := licenseclassifier.New(0.9)
		if err != nil {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			b, err := ioutil.ReadFile(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			result, confidence, err := parseResponce(b, classifier)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.result {
				t.Errorf("want: %s, got: %s", tt.result, result)
			}
			if confidence != tt.confidence {
				t.Errorf("want: %f, got: %f", tt.confidence, confidence)
			}
		})
	}
}
