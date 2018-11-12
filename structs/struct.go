package main

import (
	"fmt"
)

type car struct {
	acceleratePedal uint16 // unsigned(as in no negitives -) int
	breakPedal      uint16 // min 0 max 65535
	steeringWheel   int16  // min -32768 max 32767
	topSpeedKpH     float64
}

func main() {
	aCar := car{acceleratePedal: 2234,
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
	fmt.Println(aCar)
	fmt.Println(aCar.acceleratePedal)
	fmt.Println(aCar.breakPedal)
	fmt.Println(aCar.steeringWheel)
	fmt.Println(aCar.topSpeedKpH)
}
