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

func main() {
	go func() {
		time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
		fmt.Println("Function 1 says hi")
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
		fmt.Println("Function 2 says hi")
	}()
	time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
	fmt.Println("Main says hi")
}
