// +build OMIT

package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// START OMIT
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 8, 2, 1, 6, 6, 5})

	// Make a symmetric matrix from 2*a*a^T.
	s := &mat.SymDense{}
	s.SymOuterK(2, a)
	fmt.Printf("s = %v\n", mat.Formatted(s, mat.Prefix("    ")))

	// Compute a^T * s.
	b := &mat.Dense{}
	b.Mul(a.T(), s)

	// Construct a diagonal matrix and scale b.
	d := mat.NewDiagonal(3, []float64{0.25, 0.4, 0.2})
	fmt.Printf("d = %v\n", mat.Formatted(d, mat.Prefix("    ")))
	b.Mul(d, b)
	fmt.Printf("final =\n%0.4v\n", mat.Formatted(b))
	// END OMIT
}
