// +build OMIT

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Both goroutines write to a with no synchronization
	var a float64
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		a = 5
	}()
	go func() {
		defer wg.Done()
		a = 3
	}()
	wg.Wait()
	fmt.Println(a)
}
