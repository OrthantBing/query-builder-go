package sql

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type RuleFilter struct {
	Condition string        `json:"condition"`
	Rules     []interface{} `json:"rules"`
}

type Rule2 map[string]interface{}

type Rule struct {
	ID       string      `json:"id"`
	Field    string      `json:"field"`
	Type     string      `json:"type"`
	Input    string      `json:"input"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type Rule1 struct {
	ID       string  `json:"id"`
	Field    string  `json:"field"`
	Type     string  `json:"type"`
	Input    string  `json:"input"`
	Operator string  `json:"operator"`
	Value    float64 `json:"value"`
}

type TransformInfo struct {
	Op  string `json:"op"`
	Sep string `json:"sep"`
	Mod string `json:"mod"`
}

type SQLOperators map[string]TransformInfo

var sqlOperators = SQLOperators{
	"equal": TransformInfo{
		Op: "= %s",
	},
	"not_equal": TransformInfo{
		Op: "!= %s",
	},
	"in": TransformInfo{
		Op:  "IN(%s)",
		Sep: ", ",
	},
	"not_in": TransformInfo{
		Op:  "NOT IN(%s)",
		Sep: ", ",
	},
	"less": TransformInfo{
		Op: "< %s",
	},
	"less_or_equal": TransformInfo{
		Op: "<= %s",
	},
	"greater": TransformInfo{
		Op: "> %s",
	},
	"greater_or_equal": TransformInfo{
		Op: ">= %s",
	},
	"between": TransformInfo{
		Op: "BETWEEN %s AND %s",
	},
	"not_between": TransformInfo{
		Op: "NOT BETWEEN %s AND %s",
	},
	"begins_with": TransformInfo{
		Op:  "LIKE(%s)",
		Mod: "{0}%",
	},
	"not_begins_with": TransformInfo{
		Op:  "LIKE(%s)",
		Mod: "{0}%",
	},
	"contains": TransformInfo{
		Op:  "LIKE(%s)",
		Mod: "%{0}%",
	},
	"not_contains": TransformInfo{
		Op:  "NOT LIKE(%s)",
		Mod: "%{0}%",
	},
	"ends_with": TransformInfo{
		Op:  "LIKE(%s)",
		Mod: "%{0}",
	},
	"not_ends_with": TransformInfo{
		Op:  "NOT LIKE(%s)",
		Mod: "%{0}",
	},
	"is_empty": TransformInfo{
		Op: "= ''",
	},
	"is_not_empty": TransformInfo{
		Op: "!= ''",
	},
	"is_null": TransformInfo{
		Op: "IS NULL",
	},
	"is_not_null": TransformInfo{
		Op: "IS NOT NULL",
	},
}

type Transformer interface {
	Transform() string
}

func (rf *RuleFilter) Transform() string {
	return transform(rf)
}

func transform(rf *RuleFilter) string {
	condition := rf.Condition
	ruleArr := rf.Rules
	returnString := ""
	for _, val := range ruleArr {
		if r, ok := val.(map[string]interface{}); ok {
			if _, ok := r["condition"]; ok {
				b, err := json.Marshal(r)
				fmt.Println(err)
				var rf RuleFilter
				err = json.Unmarshal(b, &rf)
				fmt.Println(err)
				if returnString == "" {
					returnString = "(" + transform(&rf) + ")"
				} else {
					returnString = returnString + " " + condition + " " + "(" + transform(&rf) + ")"
				}

			} else {
				str := generateStringFromRule(r)
				if returnString == "" {
					returnString = str
				} else {
					returnString = returnString + " " + condition + " " + str
				}

			}
		}
	}

	return returnString
}

func generateStringFromRule(r map[string]interface{}) string {
	op := r["operator"].(string)
	rval := r["value"]
	switch r["type"].(string) {
	case "integer":
		fmt.Println(r["type"])
		fmt.Println(r["value"])
		t := sqlOperators[op]
		var val string
		if op == "between" || op == "not_between" {
			v := rval.([]int)
			val = fmt.Sprintf(t.Op, strconv.Itoa(v[0]), strconv.Itoa(v[1]))
		} else {
			val = fmt.Sprintf(t.Op, strconv.FormatFloat(rval.(float64), 'f', -1, 64))
		}

		return fmt.Sprintf("%s %s", r["field"], val)

	case "double":
		t := sqlOperators[op]
		var val string
		if op == "between" || op == "not_between" {
			v := rval.([]float64)
			val = fmt.Sprintf(t.Op, strconv.FormatFloat(v[0], 'f', -1, 64), strconv.FormatFloat(v[1], 'f', -1, 64))
		} else {
			val = fmt.Sprintf(t.Op, strconv.FormatFloat(rval.(float64), 'f', -1, 64))
		}

		return fmt.Sprintf("%s %s", r["field"], val)

	case "string":
		t := sqlOperators[op]
		var val string
		if op == "in" || op == "not_in" {
			s := strings.Split(rval.(string), ",")
			var qStringArr []string
			for _, val := range s {
				qStringArr = append(qStringArr, fmt.Sprintf("'%s'", val))
			}
			val = fmt.Sprintf(t.Op, strings.Join(qStringArr, ","))
		} else {
			val = fmt.Sprintf(t.Op, "'"+rval.(string)+"'")
		}

		return fmt.Sprintf("%s %s", r["field"].(string), val)

	default:
		return ""
	}
}
