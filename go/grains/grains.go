package grains

import (
	"errors"
)

// total is a constant value so it is defined and calculated outside the function
var total = uint64(1<<64 - 1)

// Square to find the number grains on a square
func Square(input int) (uint64, error) {
	if input < 1 {
		return 0, errors.New("number must be greater than 0")
	}

	if input > 64 {
		return 0, errors.New("number must be less than or equal to 64")
	}

	return uint64(1 << uint(input-1)), nil
}

// Total to find the total number of grains
func Total() uint64 {
	return total
}
