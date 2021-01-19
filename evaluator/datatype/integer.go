package datatype

import (
	"fmt"

	"github.com/OrthantBing/query-builder-go/utils/convert"
)

type Integer int64

type binaryIntegerRelationalFunc func(Integer, Integer) bool
type membershipIntegerRelationalFunc = func(Integer, []Integer) bool

var integerRelationalFunctions = map[string]binaryIntegerRelationalFunc{
	LessThan:       func(x, y Integer) bool { return x < y },
	Less:           func(x, y Integer) bool { return x < y },
	LessOrEqual:    func(x, y Integer) bool { return x <= y },
	GreaterThan:    func(x, y Integer) bool { return x > y },
	Greater:        func(x, y Integer) bool { return x > y },
	GreaterOrEqual: func(x, y Integer) bool { return x >= y },
	Equal:          func(x, y Integer) bool { return x == y },
	NotEqual:       func(x, y Integer) bool { return x != y },
}

var integerMembershipFunctions = map[string]membershipIntegerRelationalFunc{
	In: func(x Integer, y []Integer) bool {
		for _, s := range y {
			if s == x {
				return true
			}
		}
		return false
	},
	NotIn: func(x Integer, y []Integer) bool {
		for _, s := range y {
			if s == x {
				return false
			}
		}
		return true
	},
	Between: func(x Integer, y []Integer) bool {
		if len(y) != 2 {
			return false
		}
		return x > y[0] && x < y[1]
	},
	NotBetween: func(x Integer, y []Integer) bool {
		if len(y) != 2 {
			return false
		}

		return x <= y[0] || x >= y[1]
	},
}

func (i Integer) Result(operator string, value interface{}) (bool, error) {

	var result bool
	x := i
	switch operator {
	case In, NotIn, Between, NotBetween:
		intFunc := integerMembershipFunctions[operator]

		if convert.IsStringSlice(value) {
			v := value.([]string)
			list := make([]Integer, len(v))
			for i, val := range v {
				list[i] = Integer(convert.ParseInteger(convert.GetString(val), 0))
			}
			result = intFunc(x, list)
		} else {
			v := value.([]int64)
			list := make([]Integer, len(v))
			for i, val := range v {
				list[i] = Integer(val)
			}
			result = intFunc(x, list)
		}

	case Equal, NotEqual, Contains, NotContains, Less, LessOrEqual, LessThan,
		Greater, GreaterOrEqual, GreaterThan:
		intFunc := integerRelationalFunctions[operator]
		var y Integer
		y = Integer(convert.GetInteger(value))
		if convert.IsString(value) {
			y = Integer(convert.ParseInteger(convert.GetString(value), 0))
		}

		result = intFunc(x, y)
	default:
		return false, fmt.Errorf("Undefined Operation: %s", operator)
	}
	return result, nil
}
