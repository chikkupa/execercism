package grains

import (
	"errors"
	"math"
)

// total is a constant value so it is defined and calculated outside the function
var total = 2*uint64(math.Pow(2, 64)) - 1

// Square to find the number grains on a square
func Square(input int) (uint64, error) {
	if input < 1 {
		return 0, errors.New("number must be greater than 0")
	}

	if input > 64 {
		return 0, errors.New("number must be less than or equal to 64")
	}

	return uint64(math.Pow(2, float64(input-1))), nil
}

// Total to find the total number of grains
func Total() uint64 {
	// sum of squares of 2 is 2 * 2 ^ 64 - 1
	return total
}
