package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Uniform is a type that generates uniform random numbers with the given bound
type Uniform struct {
	Min float64
	Max float64
}

// Rand is a method on Uniform that generates the random number
func (u Uniform) Rand() float64 {
	return rand.Float64()*(u.Max-u.Min) + u.Min
}

// init functions are run before main starts.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	p := Uniform{
		Max: -2.0,
		Min: 2.0,
	}
	n := 5
	// A slice is a dynamically-sized array. Allocated with make.
	x := make([]float64, n) // Initialized to zeros
	for i := range x {      // Range over indices of x
		x[i] = p.Rand()
	}
	for i, v := range x { // Range over indices and elements of x
		fmt.Printf("x[%v] = %v\n", i, v)
	}
}
