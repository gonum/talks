// +build OMIT

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	myAddr   = ":5000"
	sendAddr = ":5001"
)

func main() {
	send, err := net.Dial("tcp", sendAddr)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp", myAddr)
	if err != nil {
		log.Fatal(err)
	}
	listen, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	encoder := gob.NewEncoder(send)
	decoder := gob.NewDecoder(listen)
	nDim := 10
	x := make([]float64, nDim)
	for {
		time.Sleep(300 * time.Millisecond)
		for i := range x {
			x[i] = rand.Float64()
		}
		if err := encoder.Encode(x); err != nil {
			log.Fatal(err)
		}
		var ans float64
		if err = decoder.Decode(&ans); err != nil {
			log.Fatal(err)
		}
		fmt.Println("answer is ", ans)
	}
}
