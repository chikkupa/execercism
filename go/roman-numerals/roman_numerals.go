package romannumerals

import (
	"errors"
	"sort"
)

// ToRomanNumeral convert normal numbers to roman numbers
func ToRomanNumeral(number int) (string, error) {
	var roman string

	if number > 3000 || number <= 0 {
		return roman, errors.New("not convertible")
	}
	romanNumSet := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	// Sort map by key order
	var dividends []int
	for k := range romanNumSet {
		dividends = append(dividends, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(dividends)))

	for _, dividend := range dividends {
		counter := number / dividend
		rep := romanNumSet[dividend]

		for i := 0; i < counter; i++ {
			roman += rep
		}

		number -= (counter * dividend)
	}

	return roman, nil
}
