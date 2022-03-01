package shared

import (
	"errors"
	"fmt"
	"net/url"
	"testing"
)

func TestCrawl(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantErr error
	}{
		{
			name:    "safty fail test",
			in:      "http://localhost/invalid-url",
			wantErr: &url.Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := new(DefaultCrawler)
			_, err := cl.Crawl(tt.in)
			if err != nil && !errors.As(tt.wantErr, &err) {
				t.Error(fmt.Sprintf("%T", err))
			}
		})
	}

}
