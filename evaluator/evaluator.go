package evaluator

import (
	"fmt"
)

type Rule struct {
	Connective string        `json:"condition"`
	Conditions []interface{} `json:"rules"`
}

type Satisfier interface {
	Satisfy(input interface{}) (bool, error)
}

// BinaryLogicalFunc is a function signature for boolean operations
type binaryLogicalFunc func(bool, bool) bool

var boolLogicalFunc = map[string]binaryLogicalFunc{
	"AND": func(x, y bool) bool { return x && y },
	"OR":  func(x, y bool) bool { return x || y },
}

func evaluateCondition(rmap map[string]interface{}, input map[string]interface{}) (bool, error) {
	c := Condition{
		ID:       rmap["id"].(string),
		Field:    rmap["field"].(string),
		Type:     rmap["type"].(string),
		Operator: rmap["operator"].(string),
		Value:    rmap["value"].(interface{}),
	}

	cResult, err := c.Evaluate(input[c.Field])
	if err != nil {
		return false, fmt.Errorf("Error in Evaluation Condition: [%s]: %#v", c, err)
	}

	return cResult, nil

}

func (r Rule) Satisfy(input map[string]interface{}) (bool, error) {
	// Empty Filter
	if len(r.Conditions) == 0 {
		return true, nil
	}

	result := false
	if r.Connective == "AND" {
		result = true
	}

	logicalFunc := boolLogicalFunc[r.Connective]
	for i, rule := range r.Conditions {
		rmap, ok := rule.(map[string]interface{})
		if !ok {
			return false, fmt.Errorf("Error in Type Assertion: Rules: %+v", rule)
		}

		if _, ok := rmap["condition"]; !ok {
			if cResult, err := evaluateCondition(rmap, input); err == nil {
				result = logicalFunc(result, cResult)
			} else {
				return false, err
			}

		} else {
			r := Rule{
				Connective: rmap["condition"].(string),
				Conditions: rmap["rules"].([]interface{}),
			}

			rResult, err := r.Satisfy(input)
			if err != nil {
				fmt.Errorf("Error in Evaluating Rule[%d]: %s: %#v", i, r, err)
				return false, fmt.Errorf("Error in Evaluating Rule[%d]: %s: %#v", i, r, err)
			}
			result = logicalFunc(result, rResult)
			fmt.Printf("After combining Rule: [%s] Result: %t, Result Update: %t\n", r, rResult, result)
		}
		fmt.Printf("Result for %s: %t\n", r, result)
	}
	return result, nil
}

// Rule implements Stringer interface
func (r Rule) String() string {
	// Thing is we need to unravel the whole recursive rules
	// into one flattened rule
	result := fmt.Sprintf("[%s][", r.Connective)

	for i, rule := range r.Conditions {
		rmap, ok := rule.(map[string]interface{})
		if !ok {
			return ""
		}

		// Check if its a simple condition or a complex rule
		if _, ok := rmap["condition"]; !ok {
			c := Condition{
				ID:       rmap["id"].(string),
				Field:    rmap["field"].(string),
				Type:     rmap["type"].(string),
				Operator: rmap["operator"].(string),
				Value:    rmap["value"].(interface{}),
			}

			result += fmt.Sprintf("(Condition: %s)", c)
			continue
		}

		// This is a rule
		r := Rule{
			Connective: rmap["condition"].(string),
			Conditions: rmap["rules"].([]interface{}),
		}

		result += fmt.Sprintf("Rule: [%d] %s", i, r)
	}
	result += "]"
	return result
}
