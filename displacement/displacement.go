package main

import (
	"fmt"
)

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	fn := func(t float64) float64 {
		var s float64 = 0.5*acceleration*t*t + initialVelocity*t + initialDisplacement
		return s
	}

	return fn
}

func main() {
	fmt.Println("Enter acceleration :")
	var acc float64
	fmt.Scan(&acc)

	fmt.Println("Enter initial velocity:")
	var vel float64
	fmt.Scan(&vel)

	fmt.Println("Enter initial displacement:")
	var dis float64
	fmt.Scan(&dis)

	fmt.Println("Enter time: ")
	var ti float64
	fmt.Scan(&ti)

	f1 := GenDisplaceFn(acc, vel, dis)
	fmt.Printf("The displacement is: %f", f1(ti))
}
