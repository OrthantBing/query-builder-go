package datatype

import (
	"fmt"

	"github.com/OrthantBing/query-builder-go/utils/convert"
)

type Double float64

type binaryDoubleRelationalFunc func(Double, Double) bool
type membershipDoubleRelationalFunc = func(Double, []Double) bool

var doubleRelationalFunctions = map[string]binaryDoubleRelationalFunc{
	LessThan:       func(x, y Double) bool { return x < y },
	Less:           func(x, y Double) bool { return x < y },
	LessOrEqual:    func(x, y Double) bool { return x <= y },
	GreaterThan:    func(x, y Double) bool { return x > y },
	Greater:        func(x, y Double) bool { return x > y },
	GreaterOrEqual: func(x, y Double) bool { return x >= y },
	Equal:          func(x, y Double) bool { return x == y },
	NotEqual:       func(x, y Double) bool { return x != y },
}

var doubleMembershipFunctions = map[string]membershipDoubleRelationalFunc{
	In: func(x Double, y []Double) bool {
		for _, s := range y {
			if s == x {
				return true
			}
		}
		return false
	},
	NotIn: func(x Double, y []Double) bool {
		for _, s := range y {
			if s == x {
				return false
			}
		}
		return true
	},
	Between: func(x Double, y []Double) bool {
		if len(y) != 2 {
			return false
		}
		return x > y[0] && x < y[1]
	},
	NotBetween: func(x Double, y []Double) bool {
		if len(y) != 2 {
			return false
		}

		return x <= y[0] || x >= y[1]
	},
}

func (i Double) Result(operator string, value interface{}) (bool, error) {

	var result bool
	x := i
	switch operator {
	case In, NotIn, Between, NotBetween:
		doubleFunc := doubleMembershipFunctions[operator]

		if convert.IsStringSlice(value) {
			v := value.([]string)
			list := make([]Double, len(v))
			for i, val := range v {
				list[i] = Double(convert.ParseFloat(convert.GetString(val), 0))
			}
			result = doubleFunc(x, list)
		} else {
			v := value.([]float64)
			list := make([]Double, len(v))
			for i, val := range v {
				list[i] = Double(val)
			}
			result = doubleFunc(x, list)
		}

	case Equal, NotEqual, Contains, NotContains, Less, LessOrEqual, LessThan,
		Greater, GreaterOrEqual, GreaterThan:
		doubleFunc := doubleRelationalFunctions[operator]
		var y Double
		y = Double(convert.GetFloat(value))
		if convert.IsString(value) {
			y = Double(convert.ParseFloat(convert.GetString(value), 0))
		}

		result = doubleFunc(x, y)
	default:
		return false, fmt.Errorf("Undefined Operation: %s", operator)
	}
	return result, nil

}
