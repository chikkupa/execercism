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

	for _, r := range s {
		go rm.incrementFrequency(r, ch)
	}

	for range s {
	_:
		<-ch
	}

	c <- 1
}

// Increment the frequency count
func (rm *RuneSyncMap) incrementFrequency(r rune, c chan int) {
	rm.Lock()
	rm.freqMap[r]++
	rm.Unlock()

	c <- 1
}
