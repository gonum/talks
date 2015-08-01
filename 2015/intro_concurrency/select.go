// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func worker(c1 <-chan float64, c2 <-chan bool, c3 chan<- float64) {
	var sum float64
	for {
		select {
		case v := <-c1:
			sum += v
		case <-c2:
			c3 <- sum
			return
		}
	}
}

func main() {
	c1 := make(chan float64)
	c2 := make(chan bool)
	c3 := make(chan float64)
	go worker(c1, c2, c3)
	for i := 0; i < 10; i++ {
		c1 <- rand.Float64()
	}
	c2 <- true
	fmt.Println("Sum of values is ", <-c3)
}
