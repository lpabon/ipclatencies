//
// Hello World Zeromq server
//
// Author: Aaron Raddon   github.com/araddon
// Requires: http://github.com/alecthomas/gozmq
//
package main

import (
	zmq "github.com/pebbe/zmq3"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	//defer context.Close()
	defer socket.Close()
	socket.Bind("ipc://latency.ipc")
	//socket.Bind("tcp://*:5555")

	// Wait for messages
	for {
		msg, _ := socket.RecvBytes(0)
		socket.SendBytes(msg, 0)
	}
}
