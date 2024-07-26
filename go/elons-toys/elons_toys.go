package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery > c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}

	fmt.Printf("car is now Car{speed: %v, batteryDrain: %v, battery: %v, distance: %v}",
		c.speed, c.batteryDrain, c.battery, c.distance)
}

func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
	return (c.battery / c.batteryDrain) * c.speed >= trackDistance
}