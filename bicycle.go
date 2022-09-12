package main

import "fmt"

const usixteenbitmax float64 = 65535
const gear_multiple float64 = 0.1
const gear_speed_limit float64 = 30

type bicycle struct {
	rpm           uint16
	breakPedal    uint16
	steeringWheel int16
	gear          uint16
}

func (c *bicycle) kmh() (speed float64) {
	speed = float64(c.rpm) * (gear_speed_limit * gear_multiple * float64(c.gear) / usixteenbitmax)
	return
}

func main() {
	var rpm uint16 = 65000
	var breakP uint16 = 0
	var steeringWheel int16 = 0
	var gear uint16 = 8

	bicycleA := bicycle{rpm, breakP, steeringWheel, gear}

	fmt.Println(bicycleA.kmh())

}
