package hamming

import "errors"

// Distance function to calculate hamming distance between toe DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The sequences are of unequal length")
	}

	var hammingDifference int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDifference++
		}
	}

	return hammingDifference, nil
}
