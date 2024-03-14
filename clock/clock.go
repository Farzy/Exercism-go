package clock

import "fmt"

const (
	dayMinutes  = 1440
	hourMinutes = 60
)

// Clock Struct representing a Clock with minutes as its only property.
type Clock struct {
	minutes int
}

// normalize is a helper function to adjust minutes to fall within the range 0 to dayMinutes (1440).
func normalize(minutes int) Clock {
	if minutes < 0 {
		return Clock{dayMinutes + (minutes % dayMinutes)}
	} else if minutes >= dayMinutes {
		return Clock{minutes % dayMinutes}
	}
	return Clock{minutes}
}

// New function to initialize a new clock set to a specific time represented by hours and minutes.
func New(h, m int) Clock {
	return normalize(h*hourMinutes + m)
}

// Add method adding a specified amount of minutes to the Clock time.
func (c Clock) Add(m int) Clock {
	return normalize(c.minutes + m)
}

// Subtract method subtracting a specified amount of minutes from the Clock time.
func (c Clock) Subtract(m int) Clock {
	return normalize(c.minutes - m)
}

// String method returning the string representation of the Clock in the "HH:MM" format.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/hourMinutes, c.minutes%hourMinutes)
}
