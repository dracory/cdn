package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestNotiflix3_2_8(t *testing.T) {
	output := Notiflix_3_2_8()
	expected := "notiflix-aio-3.2.8.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestNotiflix3_2_8_CSS(t *testing.T) {
	output := Notiflix_3_2_8_CSS()
	expected := "notiflix.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
