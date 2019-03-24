package scrabble

var scores = map[int]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// Score to find the scrabble score of a string
func Score(value string) int {
	var score int
	for i := 0; i < len(value); i++ {
		if value[i] >= 65 && value[i] <= 90 {
			index := int(value[i])
			score += scores[index]
		} else if value[i] >= 97 && value[i] <= 122 {
			index := int(value[i])
			score += scores[index-32]
		}
	}

	return score
}
