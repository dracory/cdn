package cdn

import (
	"strings"
	"testing"
)

func TestJqueryUiCss_1_14_2(t *testing.T) {
	output := JqueryUiCss_1_14_2()
	expected := "/ui/1.14.2/themes/smoothness/jquery-ui.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryUiJs_1_14_2(t *testing.T) {
	output := JqueryUiJs_1_14_2()
	expected := "/ui/1.14.2/jquery-ui.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryUiCss_1_13_1(t *testing.T) {
	output := JqueryUiCss_1_13_1()
	expected := "/ui/1.13.1/themes/smoothness/jquery-ui.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryUiJs_1_13_1(t *testing.T) {
	output := JqueryUiJs_1_13_1()
	expected := "/ui/1.13.1/jquery-ui.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
