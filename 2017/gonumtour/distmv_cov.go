// +build OMIT

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distmv"
)

func main() {
	// START OMIT
	// Estimate the covariance of a Normal distribution
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	dim := 3
	mean := []float64{2, 9, -1}
	cov := mat.NewSymDense(dim, []float64{
		5, 1, 0.3,
		1, 4, -0.1,
		0.3, -0.1, 3.5,
	})
	fmt.Printf("True Covariance\n%0.4f\n", mat.Formatted(cov))
	norm, ok := distmv.NewNormal(mean, cov, rnd)
	if !ok {
		log.Fatal("Covariance matrix not positive definite")
	}

	nSamp := 10000
	x := mat.NewDense(nSamp, dim, nil)
	for i := 0; i < nSamp; i++ {
		norm.Rand(x.RawRowView(i))
	}

	estCov := stat.CovarianceMatrix(nil, x, nil)
	fmt.Printf("Estimated Covariance\n%0.4f\n", mat.Formatted(estCov))
	// END OMIT
}
