package evaluator

import "testing"

func TestDobule(t *testing.T) {
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
				Type:     "double",
				Input:    "text",
				Operator: "equal",
				Value:    "10.5",
			},
			input: 10.5,
			mock: func() {

			},
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
				Type:     "double",
				Input:    "text",
				Operator: "not_equal",
				Value:    "10.5",
			},
			input: 11.5,
			mock: func() {

			},
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
			name: "less_than",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "double",
				Input:    "text",
				Operator: "less_than",
				Value:    "10.5",
			},
			input: 5.2,
			mock: func() {

			},
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
			name: "between",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "double",
				Input:    "text",
				Operator: "between",
				Value:    []string{"10.1", "50.4"},
			},
			input: 24.5,
			mock: func() {

			},
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
			name: "between negative",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "double",
				Input:    "text",
				Operator: "between",
				Value:    []string{"10.1", "50.4"},
			},
			input: 5.1,
			mock: func() {

			},
			want: func(b bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if b {
					t.Error("Should not return true")
				}
			},
		},

		{
			name: "not_between",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "double",
				Input:    "text",
				Operator: "not_between",
				Value:    []string{"10", "50"},
			},
			input: 5,
			mock: func() {

			},
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
