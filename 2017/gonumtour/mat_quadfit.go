// +build OMIT

package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// START OMIT
	n := 100
	x := make([]float64, n)
	for i := range x {
		x[i] = float64(i)
	}

	// Set output data
	y := make([]float64, n)
	for i, v := range x {
		y[i] = 3 + 0.5*v + 0.7*v*v + 0.01*rand.NormFloat64()
	}

	// A = [1, x, x^2]
	A := mat.NewDense(n, 3, nil)
	for i, v := range x {
		A.Set(i, 0, 1)
		A.Set(i, 1, v)
		A.Set(i, 2, v*v)
	}
	// A * beta = y  =>  beta = A^-1 * y
	beta := &mat.Vector{}
	beta.SolveVec(A, mat.NewVector(n, y))
	fmt.Printf("%0.4v\n", mat.Formatted(beta))
	// END OMIT
}
