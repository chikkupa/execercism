package letter

import "sync"

type syncMap struct {
	sync.Map
}

var frequencyMap syncMap

// ConcurrentFrequency to calculate frequency of each character parallelly
func ConcurrentFrequency(sentences []string) FreqMap {
	frequencyMap = syncMap{}

	c := make(chan int)

	for _, sentence := range sentences {
		go frequencyMap.getFrequency(sentence, c)
	}

	for range sentences {
		<-c
	}

	var freqMap = FreqMap{}

	frequencyMap.Range(func(key, value interface{}) bool {
		freqMap[key.(rune)] = value.(int)
		return true
	})

	return freqMap
}

// Calculate the frequency of each character in the string
func (sm *syncMap) getFrequency(s string, c chan int) {
	ch := make(chan int)

	for index, count := range Frequency(s) {
		go sm.incrementFrequency(index, count, ch)
	}

	for range Frequency(s) {
		<-ch
	}

	c <- 1
}

// Increment the frequency count
func (sm *syncMap) incrementFrequency(index rune, count int, c chan int) {
	currentFrequency, ok := sm.Load(index)

	if ok {
		sm.Store(index, count+currentFrequency.(int))
	} else {
		sm.Store(index, count)
	}

	c <- 1
}
