package evaluator

import (
	"encoding/json"
	"fmt"
)

type RuleFilter struct {
	Condition string        `json:"condition"`
	Rules     []interface{} `json:"rules"`
}

type Evaluator interface {
	Evaluate() bool
}

func (rf *RuleFilter) Evaluate(payload interface{}) (bool, error) {
	return evaluate(rf, payload)
}

func generateRuleFilter(r map[string]interface{}) (*RuleFilter, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	var rf RuleFilter
	err = json.Unmarshal(b, &rf)
	if err != nil {
		return nil, err
	}

	return &rf, nil
}

func evaluate(rf *RuleFilter, payload interface{}) (bool, error) {
	condition := rf.Condition
	ruleArr := rf.Rules
	returnBool := true

	for _, val := range ruleArr {
		if r, ok := val.(map[string]interface{}); ok {
			if _, ok := r["condition"]; ok {
				rf, err := generateRuleFilter(r)
				if err != nil {
					return false, err
				}

				evaluated, err := evaluate(rf, payload)
				fmt.Println(evaluated)
				if err != nil {
					return false, err
				}

			} else {

			}
		}
	}

	fmt.Println(condition, returnBool)
	return false, nil
}
