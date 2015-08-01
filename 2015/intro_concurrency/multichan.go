// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	for i := 0; i < 4; i++ {
		go producer(i, c)
	}
	for {
		fmt.Println(<-c)
	}
}

func producer(id int, c chan<- string) {
	var counter int
	for {
		counter++
		c <- fmt.Sprintf("Id: %v, v: %v", id, counter)
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
}
