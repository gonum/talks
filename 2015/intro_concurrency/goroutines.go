// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done() // Run this when the function returns
		time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
		fmt.Println("Function 1 says hi")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
		fmt.Println("Function 2 says hi")
	}()
	time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
	fmt.Println("Main says hi")
	wg.Wait() // Wait for both of the goroutines to finish
}
