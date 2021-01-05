package convert_test

import (
	"testing"
	"time"

	"lynk/driver-attendance-lambda/internal/utils/convert"
)

func TestParseInteger(t *testing.T) {
	tcase := "TestParseInteger"

	inputs := []string{
		"10",
		"0",
		"11",
		"aaaa",
	}

	outputs := []int64{
		10,
		0,
		11,
		0,
	}

	for i, input := range inputs {
		output := convert.ParseInteger(input, 0)
		if output != outputs[i] {
			t.Fatalf("%s %d Failed: Mismatching Results: %d %d", tcase, i, output, outputs[i])
		}
	}

	t.Logf("%s Passed", tcase)
}

func TestParseFloat(t *testing.T) {
	tcase := "TestParseFloat"

	inputs := []string{
		"10.0001",
		"0",
		"11.55",
		"5",
		"aaaa",
	}

	outputs := []float64{
		10.0001,
		0,
		11.55,
		5.0,
		0,
	}

	for i, input := range inputs {
		output := convert.ParseFloat(input, 0)
		if output != outputs[i] {
			t.Fatalf("%s %d Failed: Mismatching Results: %f %f", tcase, i, output, outputs[i])
		}
	}

	t.Logf("%s Passed", tcase)
}

func TestParseTime(t *testing.T) {
	tcase := "TestParseTime"

	inputs := [][2]string{
		{time.RFC3339, "2005-10-12T13:00:00Z"},
		{time.RFC3339, "2020-06-01T01:00:30.999Z"},
		{"2006-01-02", "2020-05-02"},
		{"2006-01-02 03", "2020-07-02 01"},
		{"2006-01-02 15:04:05", "2020-05-02 17:50:22"},
		{"03:04pm", "12:00pm"},
	}

	outputs := []string{
		"2005-10-12 01:00:00pm",
		"2020-06-01 01:00:30am",
		"2020-05-02 12:00:00am",
		"2020-07-02 01:00:00am",
		"2020-05-02 05:50:22pm",
		"0000-01-01 12:00:00pm",
	}

	for i, input := range inputs {
		output := convert.ParseTime(input[0], input[1]).Format("2006-01-02 03:04:05pm")
		if output != outputs[i] {
			t.Fatalf("%s %d Failed: Mismatching Results: %s %s", tcase, i, output, outputs[i])
		}
	}

	t.Logf("%s Passed", tcase)
}
