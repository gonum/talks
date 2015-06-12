package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"

	"github.com/gonum/optimize"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
}

type Rastrigin struct{}

func (Rastrigin) Func(x []float64) float64 {
	A := 10.0
	f := A * float64(len(x))
	for _, v := range x {
		f += v*v - A*math.Cos(2*math.Pi*v)
	}
	return f
}

func (r Rastrigin) Grad(x []float64, grad []float64) {
	A := 10.0
	f := A * float64(len(x))
	for i, v := range x {
		f += v*v - A*math.Cos(2*math.Pi*v)
		grad[i] = 2*v + 2*math.Pi*A*math.Sin(2*math.Pi*v)
	}
}

/*
func (r Rastrigin) Grad(x, grad []float64) {
	settings := fd.DefaultSettings()
	settings.Concurrent = true
	fd.Gradient(grad, r.Func, x, settings)
}
*/

func worker(f optimize.Function, locs chan []float64, minima chan float64) {
	for x := range locs {
		settings := optimize.DefaultSettings()
		settings.Recorder = nil
		settings.GradientThreshold = 1e-4
		result, err := optimize.Local(f, x, settings, &optimize.LBFGS{})
		if err != nil {
			minima <- math.Inf(1)
		} else {
			minima <- result.F
		}
	}
}

func randx() []float64 {
	xMin, xMax := -10.0, 10.0
	nDim := 5
	x := make([]float64, nDim)
	for i := 0; i < nDim; i++ {
		x[i] = rand.Float64()*(xMax-xMin) + xMin
	}
	return x
}

func main() {
	nRuns := 30000
	nWorkers := runtime.GOMAXPROCS(0)
	f := Rastrigin{}
	minima := make(chan float64)
	locs := make(chan []float64)

	// Launch workers
	for i := 0; i < nWorkers; i++ {
		go worker(f, locs, minima)
	}

	// Launch job producer
	go func() {
		for i := 0; i < nRuns; i++ {
			x := randx()
			locs <- x
		}
		close(locs) // Close the channel to cancel the workers
	}()
	best := math.Inf(1)
	for i := 0; i < nRuns; i++ {
		minimum := <-minima
		if minimum < best {
			best = minimum
		}
	}
	fmt.Println("Best minimum found is:", best)
}
