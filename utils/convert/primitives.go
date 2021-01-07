// Package convert performs type assertions, string parsing. Its helpful while unmarshalling
// and converting JSON types to Go types
package convert

// GetInteger Converts interface{} to Int64 if Applicable
// Otherwise returns Zero Value

func GetInteger(x interface{}) int64 {
	// Depending on Type, Perform Required Conversion
	switch x.(type) {
	case int64:
		return x.(int64)
	case int32:
		return int64(x.(int32))
	case int:
		return int64(x.(int))
	case float32, float64:
		return int64(GetFloat(x))
	default:
		return 0
	}
}

// GetFloat Converts interface{} to Float64 if Applicable
// Otherwise returns Zero Value
func GetFloat(x interface{}) float64 {
	// Depending on Type, Perform Required Conversion
	switch x.(type) {
	case int, int32, int64:
		return float64(GetInteger(x))
	case float32:
		return float64(x.(float32))
	case float64:
		return x.(float64)
	default:
		return 0
	}
}
func IsString(x interface{}) bool {
	_, ok := x.(string)
	return ok
}

func IsStringSlice(x interface{}) bool {
	_, ok := x.([]string)
	return ok
}

// GetString converts interface{} to String
func GetString(x interface{}) string {
	s, ok := x.(string)
	if !ok {
		return ""
	}
	return s
}

// GetBoolean Converts interface{} to Bool if Applicable
// Otherwise returns False Value
func GetBoolean(x interface{}) bool {
	value, ok := x.(bool)
	if !ok {
		return false
	}
	return value
}

// GetMap Converts interface{} to a Map if Applicable
// Otherwise returns an Empty Map
func GetMap(x interface{}) map[string]interface{} {
	value, ok := x.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}
	return value
}
