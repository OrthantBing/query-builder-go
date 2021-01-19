package datatype

import "testing"

func TestListString(t *testing.T) {

	tests := []struct {
		name      string
		condition string
		x         string
		y         []string
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "In Positive",
			condition: In,
			x:         "gianlugi Donurumma",
			y:         []string{"Gianlugi Donurumma", "Christiano Rolando"},
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should return true")
				}
			},
		},
		{
			name:      "In Negative",
			condition: In,
			x:         "Christian Erickson",
			y:         []string{"Gianlugi Donurumma", "Christiano Rolando"},
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if result {
					t.Error("Should return false")
				}
			},
		},
		{
			name:      "Not In Positive",
			condition: NotIn,
			x:         "Christian Erickson",
			y:         []string{"Gianlugi Donurumma", "Christiano Rolando"},
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should return true")
				}
			},
		},
		{
			name:      "Not In Negative",
			condition: NotIn,
			x:         "gianlugi donurumma",
			y:         []string{"Gianlugi Donurumma", "Christiano Rolando"},
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if result {
					t.Error("Should return false")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			x := String(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         string
		y         string
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "Equal",
			condition: Equal,
			x:         "gianlugi Donurumma",
			y:         "Gianlugi Donurumma",
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should return true")
				}
			},
		},
		{
			name:      "Contains",
			condition: Contains,
			x:         "gianlugi Donurumma",
			y:         "Gianlugi Donurumma",
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should return true")
				}
			},
		},
		{
			name:      "Contains",
			condition: Contains,
			x:         "gianlugi Donurumma",
			y:         "Gianlugi",
			mock:      func() {},
			want: func(result bool, err error) {
				if err != nil {
					t.Error(err)
				}
				if !result {
					t.Error("Should return true")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			x := String(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}
