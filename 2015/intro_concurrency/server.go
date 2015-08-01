// +build OMIT

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"github.com/gonum/optimize/functions"
)

const (
	myAddr   = ":5001"
	sendAddr = ":5000"
)

func main() {
	// Establish read/write connections
	listener, err := net.Listen("tcp", myAddr)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	send, err := net.Dial("tcp", sendAddr)
	if err != nil {
		log.Fatal(err)
	}
	// Encoders for serializing types into binary data
	encoder := gob.NewEncoder(send)
	decoder := gob.NewDecoder(conn)
	x := make([]float64, 0)
	for {
		err := decoder.Decode(&x)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Location received")
		ans := functions.ExtendedRosenbrock{}.Func(x)
		if err = encoder.Encode(ans); err != nil {
			log.Fatal(err)
		}
	}
}
