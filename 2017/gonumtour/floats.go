// +build OMIT

package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {
	s := []float64{9, -2, -6}
	max := floats.Max(s)
	fmt.Println("Max of s is", max)

	// Turn s into a unit vector.
	l := floats.Norm(s, 2)
	floats.Scale(1/l, s)
	fmt.Printf("unit vector: %0.4v\n", s)

	t := []float64{9.0 / 11, -2.0 / 11, -6.0 / 11}
	fmt.Println("s = t?", floats.EqualApprox(s, t, 1e-14))
}
