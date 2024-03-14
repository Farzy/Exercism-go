package clock

import "fmt"

// Clock represents a timekeeping device that keeps track of hours and minutes.
type Clock struct {
	hours   int
	minutes int
}

// addMinutes adds minutes together, and returns the overflow as hours.
// Negative minutes are subtracted from the hour.
func addMinutes(minutes1, minutes2 int) (hours int, minutes int) {
	hours = (minutes1 + minutes2) / 60
	minutes = (minutes1 + minutes2) % 60
	if minutes < 0 {
		minutes += 60
		hours -= 1
	}
	return
}

// addHours adds hours together, the overflow is dropped and negative hours wrap around 24.
func addHours(hours1, hours2 int) (hours int) {
	hours = (hours1 + hours2) % 24
	if hours < 0 {
		hours += 24
	}
	return
}

// New creates a new Clock with the given hours and minutes.
//
// The function first adds the minutes to the initial 0 hours using the addMinutes function.
// Then it adds the hours to the result of addMinutes using the addHours function.
// The final hours and minutes are used to initialize a new Clock object.
//
// Example:
//
//	h, m := 5, 30
//	result := New(h, m)
//	// result = Clock{hours: 5, minutes: 30}
//
// Note: The Clock struct must be defined and the addMinutes and addHours functions must be implemented.
func New(h, m int) Clock {
	hours, minutes := addMinutes(0, m)
	hours = addHours(hours, h)
	return Clock{
		hours:   hours,
		minutes: minutes,
	}
}

// Add adds the specified number of minutes to the clock and returns the new clock time.
//
// Example:
// c := Clock{hours: 5, minutes: 30}
// newClock := c.Add(45) // newClock will be {hours: 6, minutes: 15}
func (c Clock) Add(m int) Clock {
	hours, minutes := addMinutes(c.minutes, m)
	hours = addHours(c.hours, hours)
	return Clock{
		hours:   hours,
		minutes: minutes,
	}
}

// Subtract subtracts the specified number of minutes from the clock and returns the new clock time.
func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

// String returns the string representation of the Clock in the format "HH:MM".
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
