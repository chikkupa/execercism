package luhn

import (
	"fmt"
	"strings"
	"regexp"
)

// Valid check wether the number is valid
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	match,_ := regexp.MatchString(`[0-9]+`, input)
	if !match {
		return false
	}

	if len(input) <= 1 {
		return false
	}

	runeInput := []rune(input)
	
	oddFlag := true;
	if len(runeInput) % 2 == 0 {
		oddFlag = false
	}

	total := 0
	for i, number := range runeInput {
		number -= rune('0')
		if oddFlag {
			if i % 2 == 1 {
				number *= 2
				if number > 9 {
					number -= 9
				}
			}
		} else {
			if i % 2 == 0 {
				number *= 2
				if number > 9 {
					number -= 9
				}
			}
		}

		total += int(number)
	}

	fmt.Println(total)

	if total % 10 == 0 {
		return true
	}

	return false
}