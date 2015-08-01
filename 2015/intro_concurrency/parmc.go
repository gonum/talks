// +build OMIT

package main

import (
	"fmt"
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

func MonteCarlo(f func([]float64) float64, d Distribution, samples, dim int) (ev float64) {
	x := make([]float64, dim)
	for i := 0; i < samples; i++ {
		for j := range x {
			x[j] = d.Rand()
		}
		ev += f(x)
	}
	return ev / float64(samples)
}

func ParallelMonteCarlo(f func([]float64) float64, d Distribution, samples, dim int) (ev float64) {
	sz := samples / runtime.GOMAXPROCS(0)
	mux := &sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v := MonteCarlo(f, d, sz, dim)
			mux.Lock()
			ev += v
			mux.Unlock()
		}()
	}
	wg.Wait()
	return ev / float64(runtime.GOMAXPROCS(0))
}

func main() {
	f := functions.ExtendedRosenbrock{}.Func
	p := Uniform{Max: -2.0, Min: 2.0}
	ev := ParallelMonteCarlo(f, p, 100000, 3)
	fmt.Println("Expected value is:", ev)
}
