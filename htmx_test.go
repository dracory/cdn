package cdn

import (
	"os"
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

var cdnURL = "https://my-custom-cdn.com/"

func TestHtmx_2_0_8(t *testing.T) {
	os.Setenv("CDN_URL_PREFIX", cdnURL)
	defer os.Unsetenv("CDN_URL_PREFIX")

	output := Htmx_2_0_8()
	expected := "htmx.org@2.0.8"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
	if !strings.HasPrefix(output, cdnURL) {
		t.Error("Does not have custom CDN prefix")
	}
}

func TestHtmx_2_0_0(t *testing.T) {
	os.Setenv("CDN_URL_PREFIX", cdnURL)
	defer os.Unsetenv("CDN_URL_PREFIX")

	output := Htmx_2_0_0()
	expected := "htmx.org@2.0.0"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
	if !strings.HasPrefix(output, cdnURL) {
		t.Error("Does not have custom CDN prefix")
	}
}

func TestHtmx_1_9_11(t *testing.T) {
	os.Setenv("CDN_URL_PREFIX", cdnURL)
	defer os.Unsetenv("CDN_URL_PREFIX")

	output := Htmx_1_9_11()
	expected := "htmx.org@1.9.11"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
	if !strings.HasPrefix(output, cdnURL) {
		t.Error("Does not have custom CDN prefix")
	}
}

func TestHtmx_1_9_9(t *testing.T) {
	os.Setenv("CDN_URL_PREFIX", cdnURL)
	defer os.Unsetenv("CDN_URL_PREFIX")

	output := Htmx_1_9_9()
	expected := "htmx.org@1.9.9"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
	if !strings.HasPrefix(output, cdnURL) {
		t.Error("Does not have custom CDN prefix")
	}
}

func TestHtmx_1_9_6(t *testing.T) {
	os.Setenv("CDN_URL_PREFIX", cdnURL)
	defer os.Unsetenv("CDN_URL_PREFIX")

	output := Htmx_1_9_6()
	expected := "htmx.org@1.9.6"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
	if !strings.HasPrefix(output, cdnURL) {
		t.Error("Does not have custom CDN prefix")
	}
}

func TestHtmx_1_9_5(t *testing.T) {
	os.Setenv("CDN_URL_PREFIX", cdnURL)
	defer os.Unsetenv("CDN_URL_PREFIX")

	output := Htmx_1_9_5()
	expected := "htmx.org@1.9.5"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
	if !strings.HasPrefix(output, cdnURL) {
		t.Error("Does not have custom CDN prefix")
	}
}

func TestHtmx_1_9_4(t *testing.T) {
	os.Setenv("CDN_URL_PREFIX", cdnURL)
	defer os.Unsetenv("CDN_URL_PREFIX")

	output := Htmx_1_9_4()
	expected := "htmx.org@1.9.4"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
	if !strings.HasPrefix(output, cdnURL) {
		t.Error("Does not have custom CDN prefix")
	}
}

func TestHtmx_1_9_2(t *testing.T) {
	os.Setenv("CDN_URL_PREFIX", cdnURL)
	defer os.Unsetenv("CDN_URL_PREFIX")

	output := Htmx_1_9_2()
	expected := "htmx.org@1.9.2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
	if !strings.HasPrefix(output, cdnURL) {
		t.Error("Does not have custom CDN prefix")
	}
}
