//
// Hello World Zeromq Client
//
// Author: Aaron Raddon   github.com/araddon
// Requires: http://github.com/alecthomas/gozmq
//
package main

import (
	"fmt"
	zmq "github.com/alecthomas/gozmq"
	"github.com/lpabon/foocsim/utils"
	"time"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REQ)
	defer context.Close()
	defer socket.Close()

	fmt.Printf("Connecting to latency server...\n")
	socket.Connect("ipc://latency.ipc")
	//socket.Connect("tcp://*:5555")

	msg := []byte{0}
	var td utils.TimeDuration

	for i := 0; i < 1000000; i++ {
		start := time.Now()
		socket.Send([]byte(msg), 0)
		socket.Recv(0)
		end := time.Now()
		td.Add(end.Sub(start))
		if (i % 10000) == 0 {
			fmt.Print(".")
		}
	}

	fmt.Printf("\nLatency %.4f usecs\n", td.MeanTimeUsecs())
}
