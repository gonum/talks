// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gonum/optimize/functions"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Uniform is a type that generates uniform random numbers with the given bounds
type Uniform struct {
	Min float64
	Max float64
}

// Rand is a method on Uniform that generates the random number
func (u Uniform) Rand() float64 {
	return rand.Float64()*(u.Max-u.Min) + u.Min
}

func main() {
	rosenbrock := functions.ExtendedRosenbrock{}
	p := Uniform{
		Max: -2.0,
		Min: 2.0,
	}
	n := 100
	dim := 3
	var sum float64
	x := make([]float64, dim)
	for i := 0; i < n; i++ {
		for j := range x {
			x[j] = p.Rand()
		}
		sum += rosenbrock.Func(x)
	}
	fmt.Println("Expected value is ", sum/float64(n))
}
