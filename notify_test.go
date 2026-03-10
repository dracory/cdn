package cdn

import (
	"strings"
	"testing"
)

func TestNotify_3_0_0(t *testing.T) {
	output := Notify_3_0_0()
	expected := "notifyjs@3.0.0/dist/notify.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestNotify_0_4_2(t *testing.T) {
	output := Notify_0_4_2()
	expected := "cdnjs.cloudflare.com/ajax/libs/notify/0.4.2/notify.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
