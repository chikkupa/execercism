package diffsquares

import "math"

// SquareOfSum square of sum of n numbers
func SquareOfSum(n){
	int(math.Pow(n * (n + 1) / 2, 2))
}

// SumOfSquares sum of squares of n numbers
func SumOfSquares(n){
	return n * (n + 1) * (2 * n + 1) / 2
}

// Difference difference of square of sum and sum of squares of n numbers
func Difference(n){
	return n * (n + 1) * (n * ( 3 * n - 1) -2) / 12
}