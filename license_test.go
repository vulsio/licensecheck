package licensecheck

import (
	"errors"
	"io/ioutil"
	"math"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vulsio/licensecheck/shared/mock"
)

func TestScan(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		in         string
		result     string
		confidence float64
		wantErr    error
		pkgName    string
		version    string
		scanType   int
	}{
		{
			name:       "GitHub",
			in:         "./testdata/github/MIT_sample.txt",
			result:     "MIT",
			confidence: 1,
			pkgName:    "test",
			scanType:   GitHub,
		},
		{
			name:       "Go",
			in:         "./testdata/go/input1.html",
			result:     "MIT",
			confidence: 1,
			pkgName:    "test",
			version:    "v1.0",
			scanType:   Go,
		},
		{
			name:       "Java",
			in:         "./testdata/java/input1.xml",
			result:     "Apache-2.0",
			confidence: 0.911111,
			pkgName:    "test",
			version:    "v1.0",
			scanType:   Java,
		},
		{
			name:       "node",
			in:         "./testdata/nodejs/input1.json",
			result:     "MIT",
			confidence: 1,
			pkgName:    "test",
			version:    "v1.0",
			scanType:   Nodejs,
		},
		{
			name:       "Python",
			in:         "./testdata/python/input1.json",
			result:     "MIT",
			confidence: 1,
			pkgName:    "test",
			version:    "v1.0",
			scanType:   Python,
		},
		{
			name:       "Ruby",
			in:         "./testdata/ruby/input1.json",
			result:     "MIT",
			confidence: 1,
			pkgName:    "test",
			version:    "v1.0",
			scanType:   Ruby,
		},
		{
			name:       "Rust",
			in:         "./testdata/rust/input1.json",
			result:     "MIT",
			confidence: 1,
			pkgName:    "test",
			version:    "v1.0",
			scanType:   Rust,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := ioutil.ReadFile(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			sc := new(Scanner)
			cl := mock.NewMockCrawler(ctrl)
			cl.EXPECT().Crawl(gomock.Any()).Return(b, nil)
			sc.Crawler = cl

			result, confidence, err := sc.Scan(tt.pkgName, tt.version, tt.scanType)
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

func TestScanInvalidArguments(t *testing.T) {
	result, confidence, err := new(Scanner).Scan("", "", 999)
	if !errors.Is(err, ErrUnKnownScanType) {
		t.Error(err)
	}
	if result != "unknown" {
		t.Errorf("want: unknown, got: %s", result)
	}
	if confidence != 0 {
		t.Errorf("want: 0, got: %f", confidence)
	}

}
