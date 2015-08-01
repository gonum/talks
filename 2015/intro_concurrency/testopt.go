// +build OMIT

package main

import (
	"fmt"
	"log"
	"math"

	"github.com/gonum/optimize"
)

type Rastrigin struct{}

func (Rastrigin) Func(x []float64) float64 {
	A := 10.0
	f := A * float64(len(x))
	for _, v := range x {
		f += v*v - A*math.Cos(2*math.Pi*v)
	}
	fmt.Println("Rast Func = ", f)
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

func main() {
	settings := optimize.DefaultSettings()
	settings.Recorder = nil
	settings.GradientThreshold = 1e-6
	f := Rastrigin{}
	x := []float64{9.50160783681757, 0.3523567525151421, -8.042810467718468, -9.320723586564494, 0.025196429450302205}
	result, err := optimize.Local(f, x, settings, &optimize.LBFGS{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.F)
}
