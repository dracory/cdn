package cdn

import (
	"strings"
	"testing"
)

func TestFontAwesomeCss_7_2_0(t *testing.T) {
	output := FontAwesomeCss_7_2_0()
	expected := "@fortawesome/fontawesome-free@7.2.0/css/fontawesome.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestFontAwesomeCss_6_5_2(t *testing.T) {
	output := FontAwesomeCss_6_5_2()
	expected := "@fortawesome/fontawesome-free@6.5.2/css/fontawesome.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestFontAwesomeCss_6_4_2(t *testing.T) {
	output := FontAwesomeCss_6_4_2()
	expected := "@fortawesome/fontawesome-free@6.4.2/css/fontawesome.min.cs"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestFontAwesomeCss_6_1_2(t *testing.T) {
	output := FontAwesomeCss_6_1_2()
	expected := "@fortawesome/fontawesome-free@6.1.2/css/fontawesome.min.cs"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
