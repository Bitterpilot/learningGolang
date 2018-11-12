package main

import (
	"fmt"
)

const uSixteenBitMax float64 = 65535
const kphtomphMultiplier float64 = 1.60934

type car struct {
	acceleratePedal uint16 // unsigned(as in no negitives -) int
	breakPedal      uint16 // min 0 max 65535
	steeringWheel   int16  // min -32768 max 32767
	topSpeedKpH     float64
}

// value reciver method
func (c car) speedKpH() float64 {
	return float64(c.acceleratePedal) * (c.topSpeedKpH / uSixteenBitMax)
}
func (c car) speedMpH() float64 {
	return float64(c.acceleratePedal) * (c.topSpeedKpH / uSixteenBitMax / kphtomphMultiplier)
}

func main() {
	aCar := car{acceleratePedal: 65535,
		breakPedal:    0,
		steeringWheel: 6000,
		topSpeedKpH:   220}

	/* you can also just pass the values into the struct without naming them
	bCar := car{50500, 65535, 250, 180}

	the most likly use case is probally something like this
	cCar := car{acceleratePedalSensor,
				breakPedalSensor,
				steeringWheelSensor,
				topSpeedKpH:220}
	where you are parssing variables into the struct
	*/
	fmt.Println(aCar.acceleratePedal)
	fmt.Println(aCar.speedKpH())
	fmt.Println(aCar.speedMpH())
}
