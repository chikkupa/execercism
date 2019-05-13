package luhn

import (
	"regexp"
	"strings"
)

// Valid check wether the number is valid
func Valid(input string) bool {
	// Replacing all the spaces
	input = strings.Replace(input, " ", "", -1)
	// Check whether any non number value exist
	match, _ := regexp.MatchString(`[^0-9]+`, input)

	if match {
		return false
	}

	if len(input) <= 1 {
		return false
	}

	runeInput := []rune(input)

	// To identify the position of second digit to double
	oddFlag := 1
	if len(runeInput)%2 == 0 {
		oddFlag = 0
	}

	total := 0
	for i, number := range runeInput {
		// Character to number
		number -= rune('0')
		if i%2 == oddFlag {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}

		total += int(number)
	}

	if total%10 == 0 {
		return true
	}

	return false
}
