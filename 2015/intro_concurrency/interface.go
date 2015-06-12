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

// Uniform is a type that generates uniform random numbers with the given bound
type Uniform struct {
	Min float64
	Max float64
}

// Rand is a method on Uniform that generates the random number
func (u Uniform) Rand() float64 {
	return rand.Float64()*(u.Max-u.Min) + u.Min
}

type Distribution interface {
	Rand() float64
}

// MonteCarlo estimates the expected value of f under the distribution d using
// given number of samples.
// Note that f is of function type.
func MonteCarlo(f func([]float64) float64, d Distribution, samples, dim int) float64 {
	var sum float64
	x := make([]float64, dim)
	for i := 0; i < samples; i++ {
		for j := range x {
			x[j] = d.Rand()
		}
		sum += f(x)
	}
	return sum / float64(samples)
}

func main() {
	p := Uniform{Max: -2.0, Min: 2.0}
	samples, dim := 100, 3
	// Pass the type methods as the function.
	fmt.Println("Rosenbrock EV:", MonteCarlo(functions.ExtendedRosenbrock{}.Func, p, samples, dim))
	fmt.Println("Beale EV:", MonteCarlo(functions.Linear{}.Func, p, samples, dim))
}
