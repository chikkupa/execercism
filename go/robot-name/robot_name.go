package robotname

import (
	"errors"
	"math/rand"
)

// Robot Structure to save robot name
type Robot string

var record = make(map[Robot]bool)

// Name Generate a rondom name for robot and
// return the name
func (r *Robot) Name() (string, error) {
	if *r == "" {
		r.generateName()
	}
	return string(*r), nil
}

func (r *Robot) generateName() (string, error) {
	var bName []rune

	if len(record) >= (26*26*1000 - 1) {
		return "", errors.New("")
	}

	for true {
		bName = append(bName, rune('A')+rune(rand.Intn(26)))
		bName = append(bName, rune('A')+rune(rand.Intn(26)))
		bName = append(bName, rune('0')+rune(rand.Intn(10)))
		bName = append(bName, rune('0')+rune(rand.Intn(10)))
		bName = append(bName, rune('0')+rune(rand.Intn(10)))

		*r = Robot(bName)

		if !record[*r] {
			break
		}
	}

	record[*r] = true

	return string(*r), nil
}

// Reset Delete the name of the robot and regenerate its name
func (r *Robot) Reset() (string, error) {
	// record[*r] = false
	return r.generateName()
}
