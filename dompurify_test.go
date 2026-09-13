package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestDOMPurify3_4_14(t *testing.T) {
	output := DOMPurify_3_4_14()
	expected := "dompurify@3.4.14"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '"+expected+"', Output:"+output)
	}
}

func TestDOMPurify3_4_14_CDNPrefix(t *testing.T) {
	t.Setenv("CDN_URL_PREFIX", "https://example.com/")
	output := DOMPurify_3_4_14()
	expected := "https://example.com/"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '"+expected+"', Output:"+output)
	}
}
