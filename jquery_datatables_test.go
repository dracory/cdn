package cdn

import (
	"strings"
	"testing"
)

func TestJqueryDataTablesCss_2_3_7(t *testing.T) {
	output := JqueryDataTablesCss_2_3_7()
	expected := "cdn.datatables.net/2.3.7/css/jquery.dataTables.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryDataTablesJs_2_3_7(t *testing.T) {
	output := JqueryDataTablesJs_2_3_7()
	expected := "cdn.datatables.net/2.3.7/js/jquery.dataTables.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryDataTablesCss_1_13_4(t *testing.T) {
	output := JqueryDataTablesCss_1_13_4()
	expected := "cdn.datatables.net/1.13.4/css/jquery.dataTables.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryDataTablesJs_1_13_4(t *testing.T) {
	output := JqueryDataTablesJs_1_13_4()
	expected := "cdn.datatables.net/1.13.4/js/jquery.dataTables.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
