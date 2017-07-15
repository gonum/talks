// +build OMIT

package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// START OMIT
	a := mat.NewDense(2, 3, []float64{
		3, 4, 5,
		1, 2, 3,
	})
	b := mat.NewDense(3, 3, []float64{
		1, 1, 8,
		1, 2, -3,
		5, 5, 7,
	})
	fmt.Println("tr(b) =", mat.Trace(b))

	c := &mat.Dense{}
	c.Mul(a, b)
	c.Add(c, a)
	c.Mul(c, b.T())
	fmt.Printf("%v\n", mat.Formatted(c))
	// END OMIT
}
