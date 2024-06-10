package elon

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	result := fmt.Sprintf("Driven %d meters", c.distance)
	return result
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	result := fmt.Sprintf("Battery at %d%%", c.battery)
	return result
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
	carUses := c.battery / c.batteryDrain
	return carUses*c.speed >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
