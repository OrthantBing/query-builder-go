package convert

// GetAnySlice converts interface{} to []interface{}
// Mostly used for Unstructured Data
func GetAnySlice(x interface{}) []interface{} {
	s, ok := x.([]interface{})
	if !ok {
		return []interface{}{}
	}
	return s
}
