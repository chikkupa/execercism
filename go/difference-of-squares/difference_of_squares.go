package diffsquares

// SquareOfSum square of sum of n numbers
func SquareOfSum(n int) int {
	return sumOfN(n) * sumOfN(n)
}

// SumOfSquares sum of squares of n numbers
func SumOfSquares(n int) int {
	return sumOfN(n) * (2*n + 1) / 3
}

// Difference difference of square of sum and sum of squares of n numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// To find the sum of n numbers
func sumOfN(n int) int {
	return n * (n + 1) / 2
}