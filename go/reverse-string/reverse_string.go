package reverse

// String reverse a given string
func String(val string) string {
	runeVal := []rune(val)
	valLen := len(runeVal)
	for i := 0; i < valLen / 2; i++ {
		runeVal[i], runeVal[valLen - (i + 1)] = runeVal[valLen - (i + 1)], runeVal[i]
	}
	return string(runeVal)
}