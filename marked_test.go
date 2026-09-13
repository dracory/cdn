package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestMarked18_0_11(t *testing.T) {
	output := Marked_18_0_11()
	expected := "marked@18.0.11"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '"+expected+"', Output:"+output)
	}
}

func TestMarked18_0_11_CDNPrefix(t *testing.T) {
	t.Setenv("CDN_URL_PREFIX", "https://example.com/")
	output := Marked_18_0_11()
	expected := "https://example.com/"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '"+expected+"', Output:"+output)
	}
}
