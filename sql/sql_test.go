package sql

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/OrthantBing/query-builder-go/config"
)

type testCase struct {
	Input  map[string]interface{}
	Output string
}

func TestTransform(t *testing.T) {
	inpJSON, err := os.Open(fmt.Sprintf("%s/%s", config.TestFilePath, "1.json"))
	if err != nil {
		t.Error(err)
	}

	k := RuleFilter{}
	jsonParser := json.NewDecoder(inpJSON)
	if err = jsonParser.Decode(&k); err != nil {
		t.Error(err)
	}
	whereClause, err := k.Transform()
	if err != nil {
		t.Error(err)
	}
	if whereClause != "price < 10.25 AND (category = 2 OR category = 1 OR (name = 'asasdfasdfdf' AND category = 1))" {
		t.Errorf("Where clause doesnt match")
	}
	fmt.Println(whereClause)
	t.Log(whereClause)

}

func TestGenerateStringFromRule(t *testing.T) {
	tc := []testCase{
		testCase{
			Input: map[string]interface{}{
				"id":       "Pokemon",
				"field":    "Pokemon",
				"type":     "string",
				"input":    "text",
				"operator": "contains",
				"value":    "Zen Shaun",
			},
			Output: "Pokemon IN ('Zen', 'Shaun')",
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Pokemon",
				"field":    "Pokemon",
				"type":     "string",
				"input":    "text",
				"operator": "contains",
				"value":    "Zen, Shaun",
			},
			Output: "Pokemon IN ('Zen', 'Shaun')",
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Pokemon",
				"field":    "Pokemon",
				"type":     "string",
				"input":    "text",
				"operator": "begins_with",
				"value":    "Zen",
			},
			Output: "Pokemon LIKE 'Zen%' ",
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "IsPokemon",
				"field":    "IsPokemon",
				"type":     "boolean",
				"input":    "select",
				"operator": "equal",
				"value":    "true",
			},
			Output: "IsPokemon = 1",
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Hp",
				"field":    "Hp",
				"type":     "integer",
				"input":    "select",
				"operator": "not_between",
				"value":    []interface{}{500, 699},
			},
			Output: fmt.Sprintf("Hp NOT BETWEEN 500 AND 699"),
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Time",
				"field":    "Time",
				"type":     "datetime",
				"input":    "text",
				"operator": "between",
				"value": []interface{}{
					"12:53:00",
					"19:53:00",
				},
			},
			Output: fmt.Sprintf("Time BETWEEN '%s 12:53:00' AND '%s 19:53:00'", time.Now().Format("2006-02-01"), time.Now().Format("2006-02-01")),
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Time",
				"field":    "Time",
				"type":     "datetime",
				"input":    "text",
				"operator": "equal",
				"value":    "12:53:00",
			},
			Output: fmt.Sprintf("Time = '%s 12:53:00'", time.Now().Format("2006-02-01")),
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Pokemon",
				"field":    "Pokemon",
				"type":     "string",
				"input":    "checkbox",
				"operator": "in",
				"value": []interface{}{
					"SCYTHER",
					"PIKACHU",
				},
			},
			Output: fmt.Sprintf("Pokemon IN ('SCYTHER', 'PIKACHU')"),
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Shortlist",
				"field":    "Shortlist",
				"type":     "integer",
				"input":    "select",
				"operator": "less_or_equal",
				"value":    3,
			},
			Output: fmt.Sprintf("Shortlist <= 3"),
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Shortlist",
				"field":    "Shortlist",
				"type":     "double",
				"input":    "select",
				"operator": "less_or_equal",
				"value":    3,
			},
			Output: fmt.Sprintf("Shortlist <= 3"),
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Shortlist",
				"field":    "Shortlist",
				"type":     "double",
				"input":    "select",
				"operator": "less",
				"value":    3.6,
			},
			Output: fmt.Sprintf("Shortlist < 3.6"),
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Shortlist",
				"field":    "Shortlist",
				"type":     "integer",
				"input":    "select",
				"operator": "less_or_equal",
				"value":    "3",
			},
			Output: fmt.Sprintf("Shortlist <= 3"),
		},
		testCase{
			Input: map[string]interface{}{
				"id":       "Xp",
				"field":    "Xp",
				"type":     "double",
				"input":    "select",
				"operator": "less_or_equal",
				"value":    5.5,
			},
			Output: fmt.Sprintf("Xp <= 5.5"),
		},
	}

	for _, t1 := range tc {
		str, err := generateStringFromRule(t1.Input)

		if err != nil {
			t.Error(err)
		}
		if str != t1.Output {
			t.Errorf("Got: %s Expected: %s", str, t1.Output)
		}
		t.Log(str)
	}

}
