package main

import (
	"fmt"
	"math"
)

func main() {

	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Printf("Please enter a (float) value for acceleration: ")
	_, errAcc := fmt.Scan(&acceleration)
	if errAcc != nil {
		fmt.Printf("Error occurred reading time from user input: %s", errAcc.Error())
	}

	fmt.Printf("Please enter a (float) value for initial velocity: ")
	_, errIV := fmt.Scan(&initialVelocity)
	if errIV != nil {
		fmt.Printf("Error occurred reading time from user input: %s", errIV.Error())
	}

	fmt.Printf("Please enter a (float) value for initial displacement: ")
	_, errID := fmt.Scan(&initialDisplacement)
	if errID != nil {
		fmt.Printf("Error occurred reading time from user input: %s", errID.Error())
	}

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Printf("Please enter a (float) value for the time passed: ")
	_, errTime := fmt.Scan(&time)
	if errTime != nil {
		fmt.Printf("Error occurred reading time from user input: %s", errTime.Error())
	}

	fmt.Printf("Resulting displacement: %f", fn(time))
}

func GenDisplaceFn(acc, init_v, init_d float64) func (float64) float64 {

	function := func (time float64) float64 {
		return (0.5 * acc * math.Pow(time,2)) + (init_v * time) + init_d
	}
	return function
}
