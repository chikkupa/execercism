package twofer

import (
	"fmt"
)

// ShareWith function to return greeting
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	greeting := fmt.Sprintf("One for %s, one for me.", name)
	
	return greeting
}
