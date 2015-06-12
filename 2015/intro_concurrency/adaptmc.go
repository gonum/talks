package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/gonum/optimize/functions"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
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

// START GENERATOR OMIT
type Generator struct {
	F   func([]float64) float64
	D   Distribution
	Dim int
}

// Generate generates a sample with the given function and distribution
func (g Generator) Generate() float64 {
	x := make([]float64, g.Dim)
	for i := range x {
		x[i] = g.D.Rand()
	}
	return g.F(x)
}

type Result struct {
	Mean  float64
	SumSq float64
}

// END GENERATOR OMIT

// START HELPER OMIT

func update(weight, mean, sumSq, sampleWeight, sample, sampleSumSq float64) (newWeight, newMean, newSumSq float64) {
	newWeight = weight + sampleWeight
	newMean = (weight*mean + sample*sampleWeight) / newWeight
	newSumSq = sumSq + sampleSumSq + (sample-mean)*(sample-mean)*weight*sampleWeight/newWeight
	return newWeight, newMean, newSumSq
}

// END HELPER OMIT

// START_WORKER OMIT
func worker(sz int, results chan<- Result, quit <-chan bool, done chan<- bool, g Generator) {
	for {
		var number, mean, sumSq float64
		for i := 0; i < sz; i++ {
			v := g.Generate()
			// update accumulates the total sample statistics
			number, mean, sumSq = update(number, mean, sumSq, 1, v, 0)
		}
		select {
		case results <- Result{Mean: mean, SumSq: sumSq}:
		case <-quit:
			return
		}
	}
}

// END_WORKER OMIT

// START EXPECTED OMIT

func ExpectedValue(g Generator, tol float64, sz int) (ev float64, nSamples int) {
	nWorkers := runtime.GOMAXPROCS(0)
	results := make(chan Result)
	quit := make(chan bool)
	done := make(chan bool)
	for i := 0; i < nWorkers; i++ {
		go worker(sz, results, quit, done, g)
	}
	var weight, mean, sumSq float64
	for result := range results {
		weight, mean, sumSq = update(weight, mean, sumSq, float64(sz), result.Mean, result.SumSq)
		std := math.Sqrt(sumSq / (weight - 1))
		eim := std / math.Sqrt(weight)
		if eim < tol {
			break
		}
	}
	close(quit)
	return mean, int(weight)
}

// END EXPECTED OMIT

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		g1 := Generator{
			F:   functions.ExtendedRosenbrock{}.Func,
			D:   Uniform{Min: -2, Max: 2},
			Dim: 3,
		}
		ev, samples := ExpectedValue(g1, 0.5, 100)
		fmt.Println("Rosen: EV =", ev, "n =", samples)
	}()
	go func() {
		defer wg.Done()
		g1 := Generator{
			F:   functions.Beale{}.Func,
			D:   Uniform{Min: -2, Max: 2},
			Dim: 2,
		}
		ev, samples := ExpectedValue(g1, 0.1, 100)
		fmt.Println("Beale: EV =", ev, "n =", samples)
	}()
	wg.Wait()
}
