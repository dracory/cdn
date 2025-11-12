package cdn

import (
	"strings"
	"testing"
)

func TestBootstrapCss_5_3_8(t *testing.T) {
	output := BootstrapCss_5_3_8()
	expected := "bootstrap@5.3.8/dist/css/bootstrap.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs_5_3_8(t *testing.T) {
	output := BootstrapJs_5_3_8()
	expected := "bootstrap@5.3.8/dist/js/bootstrap.bundle.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCss_5_3_7(t *testing.T) {
	output := BootstrapCss_5_3_7()
	expected := "bootstrap@5.3.7/dist/css/bootstrap.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs_5_3_7(t *testing.T) {
	output := BootstrapJs_5_3_7()
	expected := "bootstrap@5.3.7/dist/js/bootstrap.bundle.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCss_5_3_3(t *testing.T) {

	output := BootstrapCss_5_3_3()
	expected := "bootstrap@5.3.3/dist/css/bootstrap.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs_5_3_3(t *testing.T) {
	output := BootstrapJs_5_3_3()
	expected := "bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCss_5_3_2(t *testing.T) {
	output := BootstrapCss_5_3_2()
	expected := "bootstrap@5.3.2/dist/css/bootstrap.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs_5_3_2(t *testing.T) {
	output := BootstrapJs_5_3_2()
	expected := "bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCss_5_3_1(t *testing.T) {
	output := BootstrapCss_5_3_1()
	expected := "bootstrap@5.3.1/dist/css/bootstrap.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs_5_3_1(t *testing.T) {
	output := BootstrapJs_5_3_1()
	expected := "bootstrap@5.3.1/dist/js/bootstrap.bundle.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCss_5_2_3(t *testing.T) {
	output := BootstrapCss_5_2_3()
	expected := "bootstrap@5.2.3/dist/css/bootstrap.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs_5_2_3(t *testing.T) {
	output := BootstrapJs_5_2_3()
	expected := "bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCss_5_3_0(t *testing.T) {
	output := BootstrapCss_5_3_0()
	expected := "bootstrap@5.3.0/dist/css/bootstrap.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs_5_3_0(t *testing.T) {
	output := BootstrapJs_5_3_0()
	expected := "bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
