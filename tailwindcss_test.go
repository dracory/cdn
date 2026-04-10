package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestTailwindCss_4_2_2(t *testing.T) {
	output := TailwindCss_4_2_2()
	expected := "tailwindcss.com/4.2.2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTailwindCss_4_2_1(t *testing.T) {
	output := TailwindCss_4_2_1()
	expected := "tailwindcss.com/4.2.1"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTailwindCss_Browser_4(t *testing.T) {
	output := TailwindCss_Browser_4()
	expected := "@tailwindcss/browser@4"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTailwindCss_3_4_4(t *testing.T) {

	output := TailwindCss_3_4_4()
	expected := "cdn.tailwindcss.com/3.4.4"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTailwindCss_3_3_3(t *testing.T) {
	output := TailwindCss_3_3_3()
	expected := "cdn.tailwindcss.com/3.3.3"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
