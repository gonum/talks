// +build OMIT

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Declare a variable. Note type is after the name.
	var xMin float64 = -2.0
	xMax := 2.0 // Type inferred automatically

	// Function call
	r := uniform(xMin, xMax)
	fmt.Println("Random value is:", r)
}

// Declare a function with two inputs and one output
func uniform(xMin, xMax float64) float64 {
	// Generate a random variable using the "rand" package and scale it
	return rand.Float64()*(xMax-xMin) + xMin
}
