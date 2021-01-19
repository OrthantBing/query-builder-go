package evaluator

import (
	"encoding/json"
	"testing"
)

func TestRule(t *testing.T) {
	tests := []struct {
		name  string
		rule  func() Rule
		input map[string]interface{}
		mock  func()
		want  func(bool, error)
	}{
		{
			name: "Basic",
			rule: func() Rule {
				rBytes, err := json.Marshal(map[string]interface{}{
					"condition": "AND",
					"rules": []map[string]interface{}{
						{
							"id":       "a",
							"field":    "a",
							"type":     "integer",
							"input":    "x",
							"operator": "less_than",
							"value":    "10",
						},
						{
							"id":       "b",
							"field":    "b",
							"type":     "integer",
							"input":    "x",
							"operator": "less_than",
							"value":    "10",
						},
						{
							"condition": "OR",
							"rules": []map[string]interface{}{
								{
									"id":       "c",
									"field":    "c",
									"type":     "string",
									"input":    "x",
									"operator": "equal",
									"value":    "5",
								},
								{
									"id":       "d",
									"field":    "d",
									"type":     "integer",
									"input":    "x",
									"operator": "less_than",
									"value":    "3",
								}},
						},
					},
				})

				if err != nil {
					t.Fatalf("Error in creating rule")

				}
				var r Rule
				if err = json.Unmarshal(rBytes, &r); err != nil {
					t.Fatalf("Error in Creating Rule")
				}

				return r
			},
			input: map[string]interface{}{
				"a": 5,
				"b": 5,
				"c": "5",
				"d": 5,
			},
			mock: func() {},
			want: func(b bool, e error) {
				if e != nil {
					t.Error(e)
				}
				if !b {
					t.Error("Should Return true")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			rule := tt.rule()
			result, err := rule.Satisfy(tt.input)
			tt.want(result, err)

		})
	}
}
