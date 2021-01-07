package evaluator

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		name      string
		condition Condition
		input     interface{}
		mock      func()
		want      func(bool, error)
	}{
		{
			name: "equal",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "string",
				Input:    "text",
				Operator: "equal",
				Value:    "Gianlugi Donurumma",
			},
			input: "Gianlugi Donurumma",
			mock:  func() {},
			want: func(b bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !b {
					t.Error("Should return true")
				}
			},
		},
		{
			name: "not_equal",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "string",
				Input:    "text",
				Operator: "not_equal",
				Value:    "Gianlugi Donurumma",
			},
			input: "Gianlugi",
			mock:  func() {},
			want: func(b bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !b {
					t.Error("Should return true")
				}
			},
		},

		{
			name: "in",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "string",
				Input:    "text",
				Operator: "in",
				Value:    []string{"Gianlugi Donurumma", "Christiano Rolando"},
			},
			input: "Gianlugi Donurumma",
			mock:  func() {},
			want: func(b bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !b {
					t.Error("Should return true")
				}
			},
		},
		{
			name: "not_in",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "string",
				Input:    "text",
				Operator: "not_in",
				Value:    []string{"Gianlugi Donurumma", "Christiano Rolando"},
			},
			input: "Christian Erickson",
			mock:  func() {},
			want: func(b bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !b {
					t.Error("Should return true")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := tt.condition.Evaluate(tt.input)
			tt.want(result, err)
		})
	}
}
