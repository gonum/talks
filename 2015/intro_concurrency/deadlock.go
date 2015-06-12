package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			c <- rand.Float64() > 0.5
		}
	}()
	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
}
