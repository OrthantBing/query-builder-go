package utils_test

import (
	"reflect"
	"testing"

	"github.com/OrthantBing/query-builder-go/utils"
)

func TestAddMap(t *testing.T) {
	tcase := "TestAddMap"

	dst := map[string]interface{}{
		"a": 10,
		"b": "asdasd",
	}

	src := map[string]interface{}{
		"c": 100,
		"d": "asdasdd",
	}

	expectedResult := map[string]interface{}{
		"a": 10,
		"b": "asdasd",
		"c": 100,
		"d": "asdasdd",
	}

	result := utils.AddMap(dst, src)
	if reflect.DeepEqual(result, expectedResult) == false {
		t.Fatalf("%s Failed: Mismatching Details: %#v %#v", tcase, result, expectedResult)
	}

	t.Logf("%s Passed", tcase)
}

func TestUnwindKeyMap(t *testing.T) {
	tcase := "TestUnwindKeyMap"

	inputs := []map[string]interface{}{
		// Child Present with Other Siblings
		{
			"driver_name":           "Tony Stark",
			"driver_truck_category": "Airborne Titanium-Plated Human F-45",
			"total_trips": map[string]interface{}{
				"value": 2,
				"trips": []string{"LY98792342", "LY87623843"},
			},
			"total_cancelled_trips": map[string]interface{}{
				"value": 2,
				"trips": []string{"LY198792342", "LY187623843"},
			},
			"total_cash_trips": map[string]interface{}{
				"value": 2,
				"trips": []string{"LY98792342", "LY87623843"},
			},
		},
		// Only Child Present
		{
			"total_trips": map[string]interface{}{
				"value": 2,
			},
			"total_cancelled_trips": map[string]interface{}{
				"value": 2,
			},
			"total_cash_trips": map[string]interface{}{
				"value": 2,
			},
		},
		// Child not present
		{
			"total_trips":           2,
			"total_cancelled_trips": 2,
			"total_cash_trips":      3,
		},
	}

	expectedResults := []map[string]interface{}{
		{
			"total_trips":           2,
			"total_cancelled_trips": 2,
			"total_cash_trips":      2,
		},
		{
			"total_trips":           2,
			"total_cancelled_trips": 2,
			"total_cash_trips":      2,
		},
		{},
	}

	for i, input := range inputs {
		if result := utils.UnwindKeyMap(input, "value"); !reflect.DeepEqual(result, expectedResults[i]) {
			t.Fatalf("%s %d Failed: Mismatching Results: %#v %#v", tcase, i, result, expectedResults[i])
		}
	}

	t.Logf("%s Passed", tcase)
}

func TestFindMap(t *testing.T) {
	tcase := "TestFindMap"

	inputs := []map[string]interface{}{
		{
			"a": 10,
			"b": 10,
			"c": 10,
			"d": 10,
		},
		{
			"a": 10,
			"b": 10,
			"c": 10,
		},
	}

	keymaps := []map[string]interface{}{
		{
			"a": 10,
			"b": 10,
		},
		{
			"e": 10,
			"b": 10,
			"c": 10,
		},
	}

	expectedResults := []map[string]interface{}{
		{
			"a": 10,
			"b": 10,
		},
		{
			"b": 10,
			"c": 10,
		},
	}

	for i, input := range inputs {
		if result := utils.FindMap(input, keymaps[i]); !reflect.DeepEqual(result, expectedResults[i]) {
			t.Fatalf("%s %d Failed: Mismatching Results: %#v %#v", tcase, i, result, expectedResults[i])
		}
	}

	t.Logf("%s Passed", tcase)
}
