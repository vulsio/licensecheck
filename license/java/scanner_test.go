package java

import (
	"errors"
	"io/ioutil"
	"math"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vuls-saas/licensecheck/license/shared"
	"github.com/vuls-saas/licensecheck/license/shared/mock"
)

func TestScanLicense(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		in         string
		result     string
		confidence float64
		wantErr    error
	}{
		{
			name:       "xml comment",
			in:         "../../testdata/java/input1.xml",
			result:     "Apache-2.0",
			confidence: 0.911111,
		},
		{
			name:       "xml body",
			in:         "../../testdata/java/input2.xml",
			result:     "Apache License, Version 2.0",
			confidence: 1,
		},
		{
			name:       "no license info",
			in:         "../../testdata/java/input3.xml",
			result:     "unknown",
			confidence: 0,
			wantErr:    shared.ErrNotFound,
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

			result, confidence, err := sc.ScanLicense("test", "")
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
