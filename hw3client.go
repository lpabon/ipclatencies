//
// Hello World Zeromq Client
//
// Author: Aaron Raddon   github.com/araddon
// Requires: http://github.com/alecthomas/gozmq
//
package main

import (
	"fmt"
	"github.com/lpabon/foocsim/utils"
	zmq "github.com/pebbe/zmq3"
	"time"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REQ)
	// defer context.Close()
	defer socket.Close()

	fmt.Printf("Connecting to latency server...\n")
	socket.Connect("ipc://latency.ipc")
	//socket.Connect("tcp://*:5555")

	msg := []byte{0}
	var td utils.TimeDuration
	prev := td.Copy()

	for i := 0; i < 1000000; i++ {
		start := time.Now()
		socket.SendBytes([]byte(msg), 0)
		socket.RecvBytes(0)
		end := time.Now()
		td.Add(end.Sub(start))
		if (i % 10000) == 0 {
			fmt.Printf("%d: %.4f usecs\n", i, td.DeltaMeanTimeUsecs(prev))
			prev = td.Copy()
		}
	}

	fmt.Printf("\nLatency %.4f usecs\n", td.MeanTimeUsecs())
}
