package raindrops

import "strconv"

// Convert Convert a number to string based on conditions
func Convert(number int) string {
	var output string
	flag := false

	if number%3 == 0 {
		output += "Pling"
		flag = true
	}

	if number%5 == 0 {
		output += "Plang"
		flag = true
	}

	if number%7 == 0 {
		output += "Plong"
		flag = true
	}

	if !flag {
		return strconv.Itoa(number)
	}

	return output
}
