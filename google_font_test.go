package cdn

import (
	"testing"
)

func TestGoogleFont(t *testing.T) {
	tests := []struct {
		family string
		weight string
		want   string
	}{
		{"Roboto", "400", "https://fonts.googleapis.com/css2?family=Roboto:wght@400&display=swap"},
		{"Open Sans", "700", "https://fonts.googleapis.com/css2?family=Open+Sans:wght@700&display=swap"},
	}

	for _, tt := range tests {
		t.Run(tt.family+"-"+tt.weight, func(t *testing.T) {
			if got := GoogleFont(tt.family, tt.weight); got != tt.want {
				t.Errorf("GoogleFont(%q, %q) = %v, want %v", tt.family, tt.weight, got, tt.want)
			}
		})
	}
}
