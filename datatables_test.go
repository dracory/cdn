package cdn

import (
	"strings"
	"testing"
)

func TestDataTablesCss_2_3_7(t *testing.T) {
	output := DataTablesCss_2_3_7()
	expected := "https://cdn.datatables.net/2.3.7/css/dataTables.dataTables.min.css"
	if !strings.Contains(output, expected) {
		t.Errorf("Does not contain '%s', Output: %s", expected, output)
	}
}

func TestDataTablesJs_2_3_7(t *testing.T) {
	output := DataTablesJs_2_3_7()
	expected := "https://cdn.datatables.net/2.3.7/js/dataTables.min.js"
	if !strings.Contains(output, expected) {
		t.Errorf("Does not contain '%s', Output: %s", expected, output)
	}
}

func TestDataTablesCss_2_2_2(t *testing.T) {
	output := DataTablesCss_2_2_2()
	expected := "https://cdn.datatables.net/2.2.2/css/dataTables.dataTables.min.css"
	if !strings.Contains(output, expected) {
		t.Errorf("Does not contain '%s', Output: %s", expected, output)
	}
}

func TestDataTablesJs_2_2_2(t *testing.T) {
	output := DataTablesJs_2_2_2()
	expected := "https://cdn.datatables.net/2.2.2/js/dataTables.min.js"
	if !strings.Contains(output, expected) {
		t.Errorf("Does not contain '%s', Output: %s", expected, output)
	}
}
