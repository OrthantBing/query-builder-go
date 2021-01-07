package evaluator

import (
	"fmt"
	"time"

	"github.com/OrthantBing/query-builder-go/evaluator/datatype"
	"github.com/OrthantBing/query-builder-go/utils/convert"
)

type Condition struct {
	ID       string      `json:"id"`
	Field    string      `json:"field"`
	Type     string      `json:"type"`
	Input    string      `json:"input"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type ConditionEvaluator interface {
	Evaluate(input interface{}) bool
}

func (c Condition) Evaluate(input interface{}) (bool, error) {
	var x datatype.Resulter

	switch c.Type {
	case "integer":
		x = datatype.Integer(convert.GetInteger(input))

	case "double":
		x = datatype.Double(convert.GetFloat(input))

	case "string":
		x = datatype.String(convert.GetString(input))

	case "time":
		s := convert.GetString(input)
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return false, fmt.Errorf("Error in Parsing Time: %#v %#v", s, err)
		}

		t, _ = time.Parse(datatype.TimeLayout, t.Format(datatype.TimeLayout))
		x = datatype.Time(t)

	default:
		return false, fmt.Errorf("Error in Performing Operation: Unknown Type: %#v", c.Type)
	}

	result, err := x.Result(c.Operator, c.Value)
	if err != nil {
		return false, err
	}

	return result, nil
}

// Condition implements Stringer interface
func (c Condition) String() string {
	return fmt.Sprintf("%s %s %#v", c.Field, c.Operator, c.Value)
}
