package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestBootstrapIconsCss_1_13_1(t *testing.T) {
	output := BootstrapIconsCss_1_13_1()
	expected := "bootstrap-icons@1.13.1/font/bootstrap-icons.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapIconsCss_1_11_3(t *testing.T) {

	output := BootstrapIconsCss_1_11_3()
	expected := "bootstrap-icons@1.11.3/font/bootstrap-icons.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapIconsCss_1_11_2(t *testing.T) {
	output := BootstrapIconsCss_1_11_2()
	expected := "bootstrap-icons@1.11.2/font/bootstrap-icons.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapIconsCss_1_11_0(t *testing.T) {
	output := BootstrapIconsCss_1_11_0()
	expected := "bootstrap-icons@1.11.0/font/bootstrap-icons.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapIconsCss_1_10_5(t *testing.T) {
	output := BootstrapIconsCss_1_10_5()
	expected := "bootstrap-icons@1.10.5/font/bootstrap-icons.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapIconsCss_1_10_2(t *testing.T) {
	output := BootstrapIconsCss_1_10_2()
	expected := "bootstrap-icons@1.10.2/font/bootstrap-icons.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapIconsCss_1_9_1(t *testing.T) {
	output := BootstrapIconsCss_1_9_1()
	expected := "bootstrap-icons@1.9.1/font/bootstrap-icons.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
