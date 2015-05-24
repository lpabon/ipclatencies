//
// Copyright (c) 2014 Luis Pabon
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/lpabon/godbc"
	"net"
	"os"
	"time"
)

type Message struct {
	conn net.Conn
	buf  []byte
}

func writer(reply chan *Message) {
	for msg := range reply {
		_, err := msg.conn.Write(msg.buf)
		if err != nil {
			fmt.Printf("W:Closed: %s\n", err.Error())
			return
		}
	}
}

func worker(workers chan *Message, reply chan *Message) {

	for msg := range workers {
		time.Sleep(time.Microsecond)
		reply <- msg
	}

}

func reader(c net.Conn,
	workers chan *Message) {
	defer c.Close()

	buf := make([]byte, 4096)

	for {
		msg := &Message{
			conn: c,
			buf:  buf,
		}

		//c.SetDeadline(time.Now().Add(time.Second * 10))
		_, err := c.Read(msg.buf)
		if err != nil {
			fmt.Printf("R:Closed: %s\n", err.Error())
			return
		}

		workers <- msg

	}
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	os.Remove("go.sock")
	l, err := net.Listen("unix", "go.sock")
	godbc.Check(err == nil)

	defer l.Close()
	defer os.Remove("go.sock")

	workers := make(chan *Message)
	reply := make(chan *Message)

	go writer(reply)
	for i := 0; i < 32; i++ {
		go worker(workers, reply)
	}

	for {
		fd, err := l.Accept()
		if err == nil {
			go reader(fd, workers)
			//go echoServer(fd)
		} else {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}
	}
}