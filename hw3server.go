//
// Hello World Zeromq server
//
// Author: Aaron Raddon   github.com/araddon
// Requires: http://github.com/alecthomas/gozmq
//
package main

import (
	"github.com/lpabon/bufferio"
	"github.com/lpabon/godbc"
	zmq "github.com/pebbe/zmq3"
	"runtime"
	"sync"
)

type Message struct {
	Id     int32
	Count  int32
	Result int32
}

func server(wg sync.WaitGroup, context *zmq.Context) {
	socket, _ := context.NewSocket(zmq.REP)
	//defer context.Close()
	defer socket.Close()
	socket.Bind("ipc://latency.ipc")
	//socket.Bind("tcp://*:5555")

	// Wait for messages
	for {
		msg, _ := socket.RecvBytes(0)

		var retmsg Message
		retbuf := bufferio.NewBufferIO(msg)
		err := retbuf.ReadDataLE(&retmsg)
		godbc.Check(err == nil)

		retmsg.Result = retmsg.Count + 10
		retbuf.Reset()
		retbuf.WriteDataLE(retmsg)

		socket.SendBytes(msg, 0)
	}
}

func main() {

	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	context, _ := zmq.NewContext()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go server(wg, context)
	}

	wg.Wait()
}
