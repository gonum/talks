// +build OMIT

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func consumer(id int, c <-chan int, wg *sync.WaitGroup) {
	for i := range c {
		time.Sleep(time.Duration(200) * time.Millisecond)
		fmt.Println("id:", id, "received", i)
	}
	fmt.Println("Consumer", id, "finished")
	wg.Done()
}

func main() {
	c := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go consumer(i, c, wg)
	}
	for i := 0; i < 30; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
}
