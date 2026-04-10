package cdn

import (
	"strings"
	"testing"
)

func TestVueElementPlusCss_2_13_7(t *testing.T) {
	output := VueElementPlusCss_2_13_7()
	expected := "element-plus@2.13.7/dist/index.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestVueElementPlusJs_2_13_7(t *testing.T) {
	output := VueElementPlusJs_2_13_7()
	expected := "element-plus@2.13.7/dist/index.full.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestVueElementPlusCss_2_13_5(t *testing.T) {
	output := VueElementPlusCss_2_13_5()
	expected := "element-plus@2.13.5/dist/index.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestVueElementPlusJs_2_13_5(t *testing.T) {
	output := VueElementPlusJs_2_13_5()
	expected := "element-plus@2.13.5/dist/index.full.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestVueElementPlusJs_2_3_8(t *testing.T) {
	output := VueElementPlusJs_2_3_8()
	expected := "element-plus@2.3.8/dist/index.full.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
