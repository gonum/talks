// +build OMIT

package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// START OMIT
	d := mat.NewDense(3, 4, nil)

	// Set d to have random integer elements.
	m, n := d.Dims()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			d.Set(i, j, float64(rand.Intn(10)))
		}
	}
	fmt.Printf("%v\n", mat.Formatted(d))

	// Print the transpose.
	fmt.Printf("%v\n", mat.Formatted(d.T()))
	// END OMIT
}
