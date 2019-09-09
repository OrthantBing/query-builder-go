package sql

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RuleFilter struct {
	Condition string        `json:"condition"`
	Rules     []interface{} `json:"rules"`
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

func (rf *RuleFilter) Transform() (string, error) {
	return transform(rf)
}

func transform(rf *RuleFilter) (string, error) {

	condition := rf.Condition
	ruleArr := rf.Rules
	returnString := ""
	for _, val := range ruleArr {
		if r, ok := val.(map[string]interface{}); ok {
			if _, ok := r["condition"]; ok {
				b, err := json.Marshal(r)
				if err != nil {
					return "", err
				}
				var rf RuleFilter
				err = json.Unmarshal(b, &rf)
				if err != nil {
					return "", err
				}

				transformed, err := transform(&rf)
				if err != nil {
					return "", err
				}
				if returnString == "" {

					returnString = "(" + transformed + ")"
				} else {
					returnString = returnString + " " + condition + " " + "(" + transformed + ")"
				}

			} else {
				str, err := generateStringFromRule(r)
				if err != nil {
					return "", err
				}
				if returnString == "" {
					returnString = str
				} else {
					returnString = returnString + " " + condition + " " + str
				}

			}
		}
	}

	return returnString, nil
}

func generateStringFromRule(r map[string]interface{}) (string, error) {
	op := r["operator"].(string)
	inp := r["input"].(string)
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
			//val = fmt.Sprintf(t.Op, strconv.FormatFloat(rval.(float64), 'f', -1, 64))
			val = fmt.Sprintf(t.Op, rval.(string))
		}

		return fmt.Sprintf("%s %s", r["field"], val), nil

	case "double":
		t := sqlOperators[op]
		var val string
		if op == "between" || op == "not_between" {
			v := rval.([]float64)
			val = fmt.Sprintf(t.Op, strconv.FormatFloat(v[0], 'f', -1, 64), strconv.FormatFloat(v[1], 'f', -1, 64))
		} else {
			val = fmt.Sprintf(t.Op, strconv.FormatFloat(rval.(float64), 'f', -1, 64))
		}

		return fmt.Sprintf("%s %s", r["field"], val), nil

	case "string":
		t := sqlOperators[op]
		var val string
		if op == "in" || op == "not_in" {
			var qStringArr []string
			if inp == "checkbox" {
				sarr := rval.([]interface{})
				for _, val := range sarr {
					qStringArr = append(qStringArr, fmt.Sprintf("'%s'", val.(string)))
				}
			} else {
				s := strings.Split(rval.(string), ",")
				for _, val := range s {
					qStringArr = append(qStringArr, fmt.Sprintf("'%s'", val))
				}
			}

			val = fmt.Sprintf(t.Op, strings.Join(qStringArr, ","))
		} else {
			val = fmt.Sprintf(t.Op, "'"+rval.(string)+"'")
		}

		return fmt.Sprintf("%s %s", r["field"].(string), val), nil

	case "boolean":
		t := sqlOperators[op]
		var val string
		if rval.(string) == "true" {
			val = fmt.Sprintf(t.Op, "1")
		} else if rval.(string) == "false" {
			val = fmt.Sprintf(t.Op, "0")
		}
		return fmt.Sprintf("%s %s", r["field"].(string), val), nil

	case "datetime":
		t := sqlOperators[op]
		var val string
		if op == "between" || op == "not_between" {
			v := rval.([]interface{})

			timeStr := []string{}
			for _, data := range v {
				validtime, err := is24HHMMSS(data.(string))
				if err != nil {
					return "", err
				}
				if validtime {
					timeStr = append(timeStr, fmt.Sprintf("%s %s", time.Now().Format("2006-02-01"), data.(string)))
				} else {
					return "", errors.New("Invalid Time format")
				}

			}
			val = fmt.Sprintf(t.Op, "'"+timeStr[0]+"'", "'"+timeStr[1]+"'")

		} else {
			validtime, err := is24HHMMSS(rval.(string))
			if err != nil {
				return "", err
			}
			if validtime {
				timestr := fmt.Sprintf("%s %s", time.Now().Format("2006-02-01"), rval.(string))
				val = fmt.Sprintf(t.Op, "'"+timestr+"'")

			} else {
				return "", errors.New("Invalid timestring")
			}
		}

		return fmt.Sprintf("%s %s", r["field"], val), nil
	default:
		return "", nil
	}
}

func is24HHMMSS(dateStr string) (bool, error) {
	r, err := regexp.Compile("[0-9]{2}:[0-9]{2}:[0-9]{2}")
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return r.MatchString(dateStr), nil
}
