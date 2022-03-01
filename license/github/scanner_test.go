package github

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vuls-saas/license-scanner/license/shared/mock"
)

func TestScanLicense(t *testing.T) {
	ctrl := gomock.NewController(t)
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
		t.Run(tt.name, func(t *testing.T) {
			b, err := ioutil.ReadFile(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			sc := new(Scanner)
			cl := mock.NewMockCrawler(ctrl)
			cl.EXPECT().Crawl("https://raw.githubusercontent.com/test/master/LICENSE").Return(nil, errors.New("test"))
			cl.EXPECT().Crawl(gomock.Any()).Return(b, nil)
			sc.Crawler = cl

			result, confidence, err := sc.ScanLicense("test", "")
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
