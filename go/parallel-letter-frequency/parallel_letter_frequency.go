package letter

import "sync"

// RuneSyncMap extends sync.RWMutex for FreqMap
type RuneSyncMap struct {
	sync.RWMutex
	freqMap FreqMap
}

// ConcurrentFrequency to calculate frequency of each character parallelly
func ConcurrentFrequency(sentences []string) FreqMap {
	runeMap := RuneSyncMap{}
	runeMap.freqMap = FreqMap{}

	c := make(chan int)

	for _, sentence := range sentences {
		go runeMap.getFrequency(sentence, c)
	}

	for range sentences {
	_:
		<-c
	}

	return runeMap.freqMap
}

// Calculate the frequency of each character in the string
func (rm *RuneSyncMap) getFrequency(s string, c chan int) {
	ch := make(chan int)

	for index, count := range Frequency(s) {
		go rm.incrementFrequency(index, count, ch)
	}

	for range Frequency(s) {
	_:
		<-ch
	}

	c <- 1
}

// Increment the frequency count
func (rm *RuneSyncMap) incrementFrequency(index rune, count int, c chan int) {
	rm.Lock()
	rm.freqMap[index] += count
	rm.Unlock()

	c <- 1
}
