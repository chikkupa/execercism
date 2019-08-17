package letter

// ConcurrentFrequency to calculate frequency of each character parallelly
func ConcurrentFrequency(sentences []string) FreqMap {
	frequencyMap := FreqMap{}

	c := make(chan FreqMap)

	for _, sentence := range sentences {
		go func(sentence string, c chan FreqMap) {
			c <- Frequency(sentence)
		}(sentence, c)
	}

	for range sentences {
		freqMap := <-c
		for index, count := range freqMap {
			frequencyMap[index] += count
		}
	}

	return frequencyMap
}
