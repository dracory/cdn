package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestChartsCss_1_2_0(t *testing.T) {
	output := ChartsCss_1_2_0()
	expected := "charts.css@1.2.0/dist/charts.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestChartsCss_1_1_0(t *testing.T) {
	output := ChartsCss_1_1_0()
	expected := "charts.css@1.1.0/dist/charts.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestChartsCss_0_9_3(t *testing.T) {
	output := ChartsCss_0_9_3()
	expected := "charts.css@0.9.0/dist/charts.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
