package etl

import "strings"

// Transform ransform the legacy data format to the shiny new format
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)

	for index, items := range input {
		for _, item := range items {
			output[strings.ToLower(string(item))] = index
		}
	}

	return output
}
