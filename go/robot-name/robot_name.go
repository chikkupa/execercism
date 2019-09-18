package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot Structure to save robot name
type Robot string

var record = make(map[Robot]bool)

// Name Generate a rondom name for robot and
// return the name
func (r *Robot) Name() (string, error) {
	if *r == "" {
		if len(record) > (26*26*10*10*10 - 1) {
			return "", errors.New("Namespace exhausted")
		}

		for true {
			r.generateName()
			if !record[*r] {
				break
			}
		}

		record[*r] = true

		return string(*r), nil
	}
	return string(*r), nil
}

func (r *Robot) generateName() {
	s1, s2 := string(rand.Intn(26)+'A'), string(rand.Intn(26)+'A')
	number := rand.Intn(1000)

	*r = Robot(fmt.Sprintf("%s%s%3d", s1, s2, number))
}

// Reset Delete the name of the robot and regenerate its name
func (r *Robot) Reset() (string, error) {
	for true {
		r.generateName()
		if !record[*r] {
			break
		}
	}

	record[*r] = true

	return string(*r), nil
}
