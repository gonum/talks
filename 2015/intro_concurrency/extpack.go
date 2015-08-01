// +build OMIT

package main

import (
	"fmt"

	// Non-standard library package
	"github.com/gonum/optimize/functions"
)

func main() {
	// Struct type from the functions package
	rosenbrock := functions.ExtendedRosenbrock{}

	x := []float64{3, 4, 5} // Literal slice allocation
	fmt.Printf("Rosenbrock(%v) = %v", x, rosenbrock.Func(x))
}
