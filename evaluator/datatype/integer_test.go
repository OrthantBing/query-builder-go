package datatype

import "testing"

func TestIntegerString(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         int64
		y         string
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "Equal",
			condition: Equal,
			x:         10,
			y:         "10",
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
			x := Integer(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}

func TestInteger(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         int64
		y         int64
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "Equal",
			condition: Equal,
			x:         10,
			y:         10,
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
			x := Integer(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}

func TestListIntegerString(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         int64
		y         []string
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "Between Positive",
			condition: Between,
			x:         10,
			y:         []string{"9", "15"},
			mock:      func() {},
			want: func(b bool, e error) {
				if e != nil {
					t.Error(e)
				}

				if !b {
					t.Error("Should Return true")
				}
			},
		},
		{
			name:      "Not Between Positive",
			condition: NotBetween,
			x:         10,
			y:         []string{"15", "25"},
			mock:      func() {},
			want: func(b bool, e error) {
				if e != nil {
					t.Error(e)
				}
				if !b {
					t.Error("Should return true")
				}
			},
		},
		{
			name:      "In Positive",
			condition: In,
			x:         10,
			y:         []string{"10", "20"},
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
			name:      "Not In Positive",
			condition: NotIn,
			x:         10,
			y:         []string{"50", "20"},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			x := Integer(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}
func TestListInteger(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         int64
		y         []int64
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "In Positive",
			condition: In,
			x:         10,
			y:         []int64{10, 20},
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
			name:      "Not In Positive",
			condition: NotIn,
			x:         10,
			y:         []int64{50, 20},
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
			x:         10,
			y:         []int64{9, 15},
			mock:      func() {},
			want: func(b bool, e error) {
				if e != nil {
					t.Error(e)
				}

				if !b {
					t.Error("Should Return true")
				}
			},
		},
		{
			name:      "Not Between Positive",
			condition: NotBetween,
			x:         10,
			y:         []int64{15, 25},
			mock:      func() {},
			want: func(b bool, e error) {
				if e != nil {
					t.Error(e)
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
			x := Integer(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}
