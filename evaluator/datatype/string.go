package datatype

import (
	"fmt"

	"strings"

	"github.com/OrthantBing/query-builder-go/utils/convert"
)

type String string

func (s String) contains(list String) bool {
	return strings.Contains(string(list), string(s))
}

type binaryStringRelationalFunc func(String, String) bool
type membershipStringRelationalFunc func(String, []String) bool

var stringRelationalFunctions = map[string]binaryStringRelationalFunc{
	LessThan:       func(x, y String) bool { return x < y },
	Less:           func(x, y String) bool { return x < y },
	LessOrEqual:    func(x, y String) bool { return x <= y },
	GreaterThan:    func(x, y String) bool { return x > y },
	Greater:        func(x, y String) bool { return x > y },
	GreaterOrEqual: func(x, y String) bool { return x >= y },
	Equal:          func(x, y String) bool { return x == y },
	NotEqual:       func(x, y String) bool { return x != y },
	Contains:       func(x, list String) bool { return list.contains(x) },
	NotContains:    func(x, list String) bool { return !list.contains(x) },
}

var stringMembershipFunctions = map[string]membershipStringRelationalFunc{
	In: func(x String, y []String) bool {
		for _, s := range y {
			if s == x {
				return true
			}
		}
		return false
	},
	NotIn: func(x String, y []String) bool {
		for _, s := range y {
			if s == x {
				return false
			}
		}
		return true
	},
}

// Implement Resulter interface
func (s String) Result(operator string, value interface{}) (bool, error) {
	var result bool

	x := String(strings.ToUpper(string(s)))
	switch operator {
	case In, NotIn:
		strFunc := stringMembershipFunctions[operator]
		v := value.([]string)
		list := make([]String, len(v))
		for i, val := range v {
			list[i] = String(strings.ToUpper(convert.GetString(val)))
		}
		result = strFunc(x, list)

	case Equal, NotEqual, Contains, NotContains, Less, LessOrEqual, LessThan,
		Greater, GreaterOrEqual, GreaterThan:
		strFunc := stringRelationalFunctions[operator]
		y := String(strings.ToUpper(convert.GetString(value)))
		result = strFunc(x, y)

	default:
		return false, fmt.Errorf("Undefined Operation: %s", operator)
	}

	return result, nil
}
