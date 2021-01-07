package datatype

import (
	"fmt"
	"testing"
	"time"

	"github.com/OrthantBing/query-builder-go/utils/convert"
)

func TestListTime(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         time.Time
		y         []string
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "In Positive",
			condition: In,
			x:         convert.ParseTime(TimeLayout, "04:10pm"),
			y:         []string{"04:10pm", "05:10pm"},
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should Return true")
				}
			},
		},
		{
			name:      "Between Positive",
			condition: Between,
			x:         convert.ParseTime(TimeLayout, "04:10pm"),
			y:         []string{"10:00am", "05:10pm"},
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should Return true")
				}
			},
		},

		{
			name:      "Not Between Positive",
			condition: NotBetween,
			x:         convert.ParseTime(TimeLayout, "04:10pm"),
			y:         []string{"05:10pm", "6:10pm"},
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should Return true")
				}
			},
		},
		{
			name:      "Between Negative",
			condition: Between,
			x:         convert.ParseTime(TimeLayout, "04:10pm"),
			y:         []string{"05:10pm", "6:10pm"},
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if result {
					t.Error("Should not Return true")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			x := Time(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}

func TestTime(t *testing.T) {

	tests := []struct {
		name      string
		condition string
		x         time.Time
		y         string
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "Equal",
			condition: Equal,
			x:         convert.ParseTime(TimeLayout, "04:10pm"),
			y:         "04:10pm",
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should be true")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			x := Time(tt.x)
			result, err := x.Result(tt.condition, tt.y)

			fmt.Println(result)
			tt.want(result, err)
		})
	}
}
