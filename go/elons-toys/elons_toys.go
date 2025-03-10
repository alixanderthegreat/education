package elon

import (
	"strconv"
)

func (c *Car) Drive() {
	if c.battery-c.batteryDrain > 0 {
		c.battery -= c.batteryDrain
		c.distance = c.speed
	}
}

func (c Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(c.distance) + " meters"
}

func (c Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(c.battery) + "%"

}

func (c Car) CanFinish(trackDistance int) bool {
	return trackDistance <= (c.battery/c.batteryDrain)*(c.speed)
}
