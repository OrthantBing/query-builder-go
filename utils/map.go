package utils

import "github.com/OrthantBing/query-builder-go/utils/convert"

// AddMap adds all keys in src to dst map
// and Returns it the resulting map
// NOTE: Make sure there are no overlapping keys
func AddMap(dst, src map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	// Copy dst map first
	for key := range dst {
		output[key] = dst[key]
	}

	// Copy src map next
	for key := range src {
		output[key] = src[key]
	}

	return output
}

// UnwindKeyMap takes a nested map as an input and returns a new map that contains the childs
// of all the fields against the original field. If key not present, that field wont be present
// in the result. It reduces one level of nesting
// NOTE: All other children will be abandoned
func UnwindKeyMap(input map[string]interface{}, child string) map[string]interface{} {
	output := make(map[string]interface{})
	for key := range input {
		innermap := convert.GetMap(input[key])

		// Check if the child is present in the field
		if _, ok := innermap[child]; !ok {
			continue
		}

		// Append only if Present
		output[key] = innermap[child]
	}

	return output
}

// FindMap extracts from input whats present in key
// If the key is not there in input, its ignored
func FindMap(input, keymap map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	for key := range keymap {
		// Not present in input
		if _, ok := input[key]; !ok {
			continue
		}
		// Otherwise, Add it
		output[key] = input[key]
	}

	return output
}
