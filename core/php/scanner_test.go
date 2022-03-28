package php

import (
	"errors"
	"io/ioutil"
	"math"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vulsio/licensecheck/shared"
	"github.com/vulsio/licensecheck/shared/mock"
)

func TestScanLicense(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		in         string
		version    string
		result     string
		confidence float64
		wantErr    error
	}{
		{
			name:       "success",
			in:         "../../testdata/php/input1.json",
			result:     "MIT",
			confidence: 1,
		},
		{
			name:       "no license info",
			in:         "../../testdata/php/input2.json",
			result:     "unknown",
			confidence: 0,
			wantErr:    shared.ErrNotFound,
		},
		{
			name:       "package that default is dev-master",
			in:         "../../testdata/php/input3.json",
			result:     "MIT",
			confidence: 1,
		},
		{
			name:       "success with version",
			in:         "../../testdata/php/input1.json",
			version:    "1.0.0",
			result:     "MIT",
			confidence: 1,
		},
		{
			name:       "no license info with version",
			in:         "../../testdata/php/input2.json",
			version:    "1.0.0",
			result:     "unknown",
			confidence: 0,
			wantErr:    shared.ErrNotFound,
		},
		{
			name:       "not exist version",
			in:         "../../testdata/php/input1.json",
			version:    "999",
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

			result, confidence, err := sc.ScanLicense("test", tt.version)
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
