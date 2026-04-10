package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestSlazy_0_5_0(t *testing.T) {
	output := Slazy_0_5_0()
	expected := "lesichkovm/slazy@0.5.0/dist/slazy.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
