//
// Hello World Zeromq Client
//
// Author: Aaron Raddon   github.com/araddon
// Requires: http://github.com/alecthomas/gozmq
//
package main

import (
	"fmt"
	"github.com/lpabon/bufferio"
	"github.com/lpabon/foocsim/utils"
	"github.com/lpabon/godbc"
	zmq "github.com/pebbe/zmq3"
	"sync"
	"time"
)

type Message struct {
	Id     int32
	Count  int32
	Result int32
}

func client(wg sync.WaitGroup, context *zmq.Context, id int) {
	defer wg.Done()
	socket, _ := context.NewSocket(zmq.REQ)
	// defer context.Close()
	defer socket.Close()

	fmt.Printf("%d: Connecting to latency server...\n", id)
	socket.Connect("ipc://latency.ipc")
	//socket.Connect("tcp://*:5555")

	var msg = Message{Id: int32(id)}
	buf := bufferio.NewBufferIOMake(32)

	var td utils.TimeDuration
	prev := td.Copy()

	for i := 1; i < 1000000; i++ {
		buf.Reset()
		msg.Count = int32(i)
		err := buf.WriteDataLE(msg)
		godbc.Check(err == nil, fmt.Sprintln(err))

		time.Sleep(time.Microsecond * 50)

		start := time.Now()
		socket.SendBytes(buf.Bytes(), 0)
		ret, _ := socket.RecvBytes(0)
		end := time.Now()

		var retmsg Message
		retbuf := bufferio.NewBufferIO(ret)
		err = retbuf.ReadDataLE(&retmsg)
		godbc.Check(err == nil)
		godbc.Check(retmsg.Id == int32(id))
		godbc.Check(retmsg.Result == (msg.Count + 10))

		td.Add(end.Sub(start))
		if (i % 10000) == 0 {
			fmt.Printf("%d:%d: %.4f usecs\n", id, i, td.DeltaMeanTimeUsecs(prev))
			prev = td.Copy()
		}
	}

	fmt.Printf("\n%d:Latency %.4f usecs\n", id, td.MeanTimeUsecs())
}

func main() {

	var wg sync.WaitGroup
	context, _ := zmq.NewContext()
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go client(wg, context, i)
	}

	wg.Wait()
}
