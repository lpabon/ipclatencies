//go:generate msgp
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
	mypkt "./mypkt"
	"fmt"
	"github.com/lpabon/godbc"
	"net"
	"os"
	"runtime"
)

func echoServer(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 4096)

	for {

		//c.SetDeadline(time.Now().Add(time.Second * 10))
		rpkt := mypkt.MyPkt{}
		l, err := c.Read(buf)
		if err != nil {
			fmt.Printf("R:Closed: %s\n", err.Error())
			return
		}
		rpkt.UnmarshalArrayMsg(buf[:l])
		fmt.Print(rpkt)

		pkt := mypkt.MyPkt{}
		pkt.Id = 1234
		pkt.Block = 1234567890
		pkt.NBlocks = 700
		pkt.Path = "/usr/home/path"
		//b := make([]byte, pkt.Msgsize())
		b, _ := pkt.MarshalArrayMsg()

		go func() {
			_, err = c.Write(b)
			if err != nil {
				fmt.Printf("W:Closed: %s\n", err.Error())
				return
			}
		}()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	os.Remove("/tmp/go.sock")
	l, err := net.Listen("unix", "/tmp/go.sock")
	godbc.Check(err == nil)

	defer l.Close()
	defer os.Remove("go.sock")

	for {
		fd, err := l.Accept()
		if err == nil {
			go echoServer(fd)
		} else {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}
	}
}
