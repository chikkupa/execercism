package isogram

import (
	"strings"
)

// IsIsogram function to check whether a given string is isogram or not
func IsIsogram(value string) bool {
	value = strings.ToUpper(value)

	for i, char := range value {
		if char != ' ' && char != '-' && strings.Contains(value[i+1:], string(char)) {
			return false
		}
	}
	return true
}
