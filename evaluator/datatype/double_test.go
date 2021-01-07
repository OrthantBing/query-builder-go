package datatype

import "testing"

func TestDoubleString(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         float64
		y         string
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "Equal",
			condition: Equal,
			x:         10.5,
			y:         "10.5",
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
			x := Double(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}

func TestDouble(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         float64
		y         float64
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "Equal",
			condition: Equal,
			x:         10.5,
			y:         10.5,
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
			x := Double(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}

func TestListDoubleString(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         float64
		y         []string
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "In Positive",
			condition: In,
			x:         10.5,
			y:         []string{"10.5", "20"},
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
			x:         10.5,
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
		{
			name:      "Between Positive",
			condition: Between,
			x:         10.6,
			y:         []string{"9.9", "15"},
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
			x:         10.7,
			y:         []string{"15.1", "25"},
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
			x := Double(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}

func TestListDouble(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		x         float64
		y         []float64
		mock      func()
		want      func(bool, error)
	}{
		{
			name:      "In Positive",
			condition: In,
			x:         10.5,
			y:         []float64{10.5, 20},
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
			x:         10.5,
			y:         []float64{50, 20},
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
			x:         10.5,
			y:         []float64{9.3, 15},
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
			x:         10.2,
			y:         []float64{15.6, 25},
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
			x := Double(tt.x)
			result, err := x.Result(tt.condition, tt.y)
			tt.want(result, err)
		})
	}
}
