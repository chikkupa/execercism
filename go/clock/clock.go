package clock

import (
	"fmt"
)

// Clock An integer to represnt time in minutes
type Clock int

// New creates a new Clock
func New(hour, minute int) Clock {
	min := hour*60 + minute

	min %= 24 * 60
	if min < 0 {
		min = 24*60 + min
	}
	return Clock(min)
}

// String stringifies the Clock
func (c Clock) String() string {
	hours := c / 60
	mins := c % 60

	return fmt.Sprintf("%02d:%02d", int(hours), int(mins))
}

// Add adds minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

// Subtract substracts minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}
