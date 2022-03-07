package java

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToMavenPomName(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		version string
		want    string
	}{
		{
			name:    "example",
			in:      "org.example:example",
			version: "1",
			want:    "org/example/example/1/example-1.pom",
		},
		{
			name:    "actual",
			in:      "org.apache.logging.log4j:log4j-core",
			version: "2.17.1",
			want:    "org/apache/logging/log4j/log4j-core/2.17.1/log4j-core-2.17.1.pom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToMavenPomName(tt.in, tt.version)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(-got +want): %s", diff)
			}
		})
	}
}
