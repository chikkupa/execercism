package letter

// ConcurrentFrequency to calculate frequency of each character parallelly
func ConcurrentFrequency(sentences []string) FreqMap {
	frequencyMap := FreqMap{}

	c := make(chan FreqMap)

	for _, sentence := range sentences {
		go getFrequency(sentence, c)
	}

	for range sentences {
		freqMap := <-c
		for index, count := range freqMap {
			frequencyMap[index] += count
		}
	}

	return frequencyMap
}

// Calculate the frequency of each character in the string
func getFrequency(s string, c chan FreqMap) {
	freqMap := FreqMap{}

	for index, count := range Frequency(s) {
		freqMap[index] += count
	}

	c <- freqMap
}
