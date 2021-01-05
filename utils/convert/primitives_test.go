package convert_test

import (
	"testing"

	"github.com/OrthantBing/query-builder-go/utils/convert"
)

func TestGetInteger(t *testing.T) {
	tcase := "TestGetInteger"

	inputs := []interface{}{
		int(30),
		int32(10),
		int64(20),
		float64(20.45),
		0,
	}

	outputs := []int64{
		30,
		10,
		20,
		20,
		0,
	}

	for i, input := range inputs {
		output := convert.GetInteger(input)
		if output != outputs[i] {
			t.Fatalf("%s %d Failed: Mismatching Results: %d %d", tcase, i, output, outputs[i])
		}
	}

	t.Logf("%s Passed", tcase)
}

func TestGetFloat(t *testing.T) {
	tcase := "TestGetFloat"

	inputs := []interface{}{
		float32(30),
		int32(10),
		int64(20),
		float64(20.45),
		0,
	}

	outputs := []float64{
		30,
		10,
		20,
		20.45,
		0,
	}

	for i, input := range inputs {
		output := convert.GetFloat(input)
		if output != outputs[i] {
			t.Fatalf("%s %d Failed: Mismatching Results: %f %f", tcase, i, output, outputs[i])
		}
	}

	t.Logf("%s Passed", tcase)
}
