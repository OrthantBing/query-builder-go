package evaluator

import "testing"

func TestTime(t *testing.T) {
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
				Type:     "time",
				Input:    "text",
				Operator: "equal",
				Value:    "04:10pm",
			},
			input: "2006-01-02T16:10:00Z",
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
				Type:     "time",
				Input:    "text",
				Operator: "not_equal",
				Value:    "05:10pm",
			},
			input: "2006-01-02T16:10:00Z",
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
			name: "less_than",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "time",
				Input:    "text",
				Operator: "less_than",
				Value:    "05:10pm",
			},
			input: "2006-01-02T16:10:00Z",
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
			name: "between",
			condition: Condition{
				ID:       "a",
				Field:    "a",
				Type:     "time",
				Input:    "text",
				Operator: "between",
				Value:    []string{"03:10pm", "06:10pm"},
			},
			input: "2006-01-02T16:10:00Z",
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
