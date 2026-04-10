package cdn

import (
	"strings"
	"testing"
)

func TestVueJs_3_5_32(t *testing.T) {
	output := VueJs_3_5_32()
	expected := "vue@3.5.32/dist/vue.global.prod.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestVueJs_3_5_30(t *testing.T) {
	output := VueJs_3_5_30()
	expected := "vue@3.5.30/dist/vue.global.prod.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestVueJs_3(t *testing.T) {
	output := VueJs_3()
	expected := "vue@3/dist/vue.global.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
