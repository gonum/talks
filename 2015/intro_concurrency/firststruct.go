package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Uniform is a type that generates uniform random numbers with the given bounds
type Uniform struct {
	Min float64
	Max float64
}

// Rand is a method on Uniform
func (u Uniform) Rand() float64 {
	return rand.Float64()*(u.Max-u.Min) + u.Min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := Uniform{
		Min: -2,
		Max: 2,
	}
	for i := 0; i < 5; i++ { // C-stlye range syntax
		fmt.Println("Random value is: ", p.Rand())
	}
}
