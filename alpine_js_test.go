package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestAlpine_3_15_11(t *testing.T) {
	output := AlpineJs_3_15_11()
	expected := "alpinejs@3.15.11/dist/cdn.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestAlpine_3_15_8(t *testing.T) {
	output := AlpineJs_3_15_8()
	expected := "alpinejs@3.15.8/dist/cdn.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestAlpine_3_13_8(t *testing.T) {
	output := AlpineJs_3_13_8()
	expected := "alpinejs@3.13.8/dist/cdn.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestAlpine_3_12_3(t *testing.T) {
	output := AlpineJs_3_12_3()
	expected := "alpinejs@3.12.3/dist/cdn.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
