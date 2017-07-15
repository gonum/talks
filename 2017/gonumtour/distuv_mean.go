// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// START OMIT
	// Estimate the mean of a ChiSquared random variable

	// Set a random number seed
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	dist := distuv.ChiSquared{K: 10, Src: rnd}
	nSamp := 10000
	x := make([]float64, nSamp)
	for i := range x {
		x[i] = dist.Rand()
	}
	mean := stat.Mean(x, nil)
	fmt.Println("True mean:", dist.Mean())
	fmt.Println("Estimated mean:", mean)
	// END OMIT
}
