package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestJqTreeCss_1_8_8(t *testing.T) {
	output := JqTreeCss_1_8_8()
	expected := "jqtree/1.8.8/jqtree.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeJs_1_8_8(t *testing.T) {
	output := JqTreeJs_1_8_8()
	expected := "jqtree/1.8.8/tree.jquery.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeCss_1_8_3(t *testing.T) {
	output := JqTreeCss_1_8_3()
	expected := "jqtree/1.8.3/jqtree.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeJs_1_8_3(t *testing.T) {
	output := JqTreeJs_1_8_3()
	expected := "jqtree/1.8.3/tree.jquery.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeCss_1_8_2(t *testing.T) {
	output := JqTreeCss_1_8_2()
	expected := "jqtree/1.8.2/jqtree.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeJs_1_8_2(t *testing.T) {
	output := JqTreeJs_1_8_2()
	expected := "jqtree/1.8.2/tree.jquery.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeCss_1_8_0(t *testing.T) {
	output := JqTreeCss_1_8_0()
	expected := "jqtree/1.8.0/jqtree.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeJs_1_8_0(t *testing.T) {
	output := JqTreeJs_1_8_0()
	expected := "jqtree/1.8.0/tree.jquery.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeCss_1_7_0(t *testing.T) {
	output := JqTreeCss_1_7_0()
	expected := "jqtree/1.7.0/jqtree.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeJs_1_7_0(t *testing.T) {
	output := JqTreeJs_1_7_0()
	expected := "jqtree/1.7.0/tree.jquery.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
