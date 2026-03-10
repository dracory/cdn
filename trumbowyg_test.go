package cdn

import (
	"strings"
	"testing"
)

func TestTrumbowygCss_2_31_0(t *testing.T) {
	output := TrumbowygCss_2_31_0()
	expected := "/2.31.0/ui/trumbowyg.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTrumbowygJs_2_31_0(t *testing.T) {
	output := TrumbowygJs_2_31_0()
	expected := "/2.31.0/trumbowyg.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTrumbowygCss_2_27_3(t *testing.T) {
	output := TrumbowygCss_2_27_3()
	expected := "/2.27.3/ui/trumbowyg.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTrumbowygJs_2_27_3(t *testing.T) {
	output := TrumbowygJs_2_27_3()
	expected := "/2.27.3/trumbowyg.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
