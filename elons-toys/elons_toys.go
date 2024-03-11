package elon

import "fmt"

// Drive drives the car if it has enough battery.
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance returns the LED display for the distance.
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns the LED display for the battery percentage.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish returns true if the car can finish the race, otherwise returns false.
func (c Car) CanFinish(trackDistance int) bool {
	rounds := float64(trackDistance) / float64(c.speed)
	if float64(c.battery) >= rounds*float64(c.batteryDrain) {
		return true
	} else {
		return false
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
