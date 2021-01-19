package datatype

import (
	"fmt"
	"time"

	"github.com/OrthantBing/query-builder-go/utils/convert"
)

type Time time.Time

var TimeLayout = "03:04pm"

func (t Time) before(x Time) bool {
	p, q := time.Time(t), time.Time(x)
	return p.Before(q)
}

func (t Time) after(x Time) bool {
	p, q := time.Time(t), time.Time(x)
	return p.After(q)
}

func (t Time) equal(x Time) bool {
	p, q := time.Time(t), time.Time(x)
	return p.Equal(q)
}

type binaryDateRelationalFunc func(Time, Time) bool
type membershipDateFunc func(Time, []Time) bool

var dateRelationalFunctions = map[string]binaryDateRelationalFunc{
	LessThan:       func(x, y Time) bool { return x.before(y) },
	Less:           func(x, y Time) bool { return x.before(y) },
	LessOrEqual:    func(x, y Time) bool { return x.equal(y) || x.before(y) },
	GreaterThan:    func(x, y Time) bool { return x.equal(y) || x.after(y) },
	Greater:        func(x, y Time) bool { return x.equal(y) || x.after(y) },
	GreaterOrEqual: func(x, y Time) bool { return x.after(y) },
	Equal: func(x, y Time) bool {
		fmt.Println(time.Time(x).String(), time.Time(y).String())
		return x.equal(y)
	},
	NotEqual: func(x, y Time) bool { return !x.equal(y) },
}
var dateMembershipFunctions = map[string]membershipDateFunc{
	In: func(x Time, y []Time) bool {
		for _, i := range y {
			if i.equal(x) {
				return true
			}
		}
		return false
	},
	NotIn: func(x Time, y []Time) bool {
		for _, i := range y {
			if i.equal(x) {
				return false
			}
		}
		return true
	},
	Between: func(x Time, y []Time) bool {
		if len(y) < 2 {
			return false
		}
		return x.after(y[0]) && x.before(y[1])
	},
	NotBetween: func(x Time, y []Time) bool {
		if len(y) < 2 {
			return false
		}
		return x.before(y[0]) || x.after(y[1])
	},
}

func (t Time) Result(operator string, value interface{}) (bool, error) {
	var result bool
	switch operator {
	case In, NotIn, Between, NotBetween:
		dateFunc := dateMembershipFunctions[operator]
		v := value.([]string)
		s := make([]Time, len(v))
		for i, val := range v {
			t := convert.ParseTime(TimeLayout, convert.GetString(val))
			s[i] = Time(t)
		}
		result = dateFunc(t, s)

	case Equal, NotEqual, Less, LessThan, LessOrEqual, Greater, GreaterThan, GreaterOrEqual:
		dateFunc := dateRelationalFunctions[operator]
		s := convert.ParseTime(TimeLayout, convert.GetString(value))
		result = dateFunc(t, Time(s))

	default:
		return false, fmt.Errorf("Undefined Operation: %s", operator)
	}
	return result, nil
}
