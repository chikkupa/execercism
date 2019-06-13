package letter

import (
	"sync"
)

// RuneSyncMap extends sync.RWMutex for FreqMap
type RuneSyncMap struct {
	sync.RWMutex
	freqMap FreqMap
}

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
	_:
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
func (rm *syncMap) getFrequency(s string, c chan int) {
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
func (rm *syncMap) incrementFrequency(index rune, count int, c chan int) {
	currentFrequency, ok := rm.Load(index)

	if ok {
		rm.Store(index, count+currentFrequency.(int))
	} else {
		rm.Store(index, count)
	}

	c <- 1
}