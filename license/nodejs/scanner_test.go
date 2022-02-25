package nodejs

import (
	"errors"
	"io/ioutil"
	"math"
	"testing"
)

func TestParseResponce(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		result     string
		confidence float64
		wantErr    error
	}{
		{
			name:       "success",
			in:         "../../testdata/nodejs/input1.json",
			result:     "MIT",
			confidence: 1,
		},
		{
			name:       "no license info",
			in:         "../../testdata/nodejs/input2.json",
			result:     "",
			confidence: 0,
			wantErr:    errNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := ioutil.ReadFile(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			result, confidence, err := parseResponce(b)
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Fatal(err)
			}
			if result != tt.result {
				t.Errorf("want: %s, got: %s", tt.result, result)
			}
			if math.Abs(confidence-tt.confidence) >= 1e-6 {
				t.Errorf("want: %f, got: %f", tt.confidence, confidence)
			}
		})
	}
}
