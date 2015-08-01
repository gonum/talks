// +build OMIT

package main

import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	go func() {
		s := <-c2
		fmt.Println(s)
		c3 <- "Hi main goroutine!"
	}()
	go func() {
		s := <-c1
		fmt.Println(s)
		c2 <- "Hi goroutine 1!"
	}()
	c1 <- "Hi goroutine 2!"
	s := <-c3
	fmt.Println(s)
}
