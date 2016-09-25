package clock

import (
	"fmt"
)

const testVersion = 4
const HOURS = 24
const MINUTES = 60

type Clock struct {
	Hour   int
	Minute int
}

// math.Abs receives float
func abs(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}

func prepareHours(hour int) int {
	// Compute value for hours
	if abs(hour) >= HOURS {
		hour = hour % HOURS
	}

	// Hours should be positive
	if hour < 0 {
		hour = HOURS + hour
	}

	return hour
}

func prepareMinutes(minute int) (minutes, carry int) {
	carry = 0
	minutes = minute

	// Compute carry for hours and minutes
	if abs(minute) >= MINUTES {
		carry = minute / MINUTES
		minutes = minute % MINUTES
	}

	// Even if minutes are less than MINUTES, but minutes are negatieve e.g -40
	// carry should be -1 New(1, -40) == "0: 20"
	if minutes < 0 {
		carry = carry - 1
	}

	// Compute positive values for minutes
	if minutes < 0 {
		minutes = MINUTES + minutes
	}

	return
}

func New(hour, minute int) Clock {
	minute, carry := prepareMinutes(minute)
	hour = prepareHours(hour + carry)

	return Clock{
		Hour:   hour,
		Minute: minute,
	}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.Hour, clock.Minute)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.Hour, clock.Minute+minutes)
}
