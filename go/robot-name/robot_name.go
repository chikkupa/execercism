package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot Structure to save robot name
type Robot struct{ RobotName string }

var record = make(map[string]bool)

// Name Generate a rondom name for robot and
// return the name
func (r *Robot) Name() (string, error) {
	if r.RobotName == "" {
		if len(record) > (26*26*10*10*10 - 1) {
			return "", errors.New("Namespace exhausted")
		}

		for true {
			r.generateName()
			if !record[r.RobotName] {
				break
			}
		}

		record[r.RobotName] = true
	}
	return r.RobotName, nil
}

func (r *Robot) generateName() {
	s1, s2 := string(rand.Intn(26)+'A'), string(rand.Intn(26)+'A')
	number := rand.Intn(1000)

	r.RobotName = fmt.Sprintf("%s%s%3d", s1, s2, number)
}

// Reset Delete the name of the robot and regenerate its name
func (r *Robot) Reset() (string, error) {
	for true {
		r.generateName()
		if !record[r.RobotName] {
			break
		}
	}

	record[r.RobotName] = true

	return r.RobotName, nil
}
