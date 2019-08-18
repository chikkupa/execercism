package summultiples

// SumMultiples returns the sum of multiples
func SumMultiples(limit int, divisors []int) int {

}

func getMultiples(limit int, divisor int) []int {
	var multiples []int

	for i := divisor; i <= divisor; i += divisor {
		multiples = append(multiples, i)
	}
	return multiples
}

func mergeArrays(array1 []int, array2 []int) []int {
	for i := 0; i < len(array2); i++ {
		for j := 0; j < len(array1); j++ {
			if array2[i] < array1[j] {
				array1 = append(array1[:j], array2[i:i], array1[j:])
			}
		}
	}
	return array1
}
