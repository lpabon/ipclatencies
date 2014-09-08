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
	"runtime"
	"time"
)

func echoServer(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 512)
	fmt.Print("Created connection\n")

	for {
		var nr int
		var err error

		c.SetDeadline(time.Now().Add(time.Second * 10))
		for errors := 0; ; errors++ {
			nr, err = c.Read(buf)
			if err != nil {
				fmt.Printf("R:Closed: %s\n", err.Error())
				if errors >= 10 {
					return
				}
			} else {
				//fmt.Print("R")
				break
			}
		}

		for errors := 0; ; errors++ {
			_, err = c.Write(buf[:nr])
			if err != nil {
				fmt.Printf("W:Closed: %s\n", err.Error())
				if errors >= 10 {
					return
				}
				return
			} else {
				//fmt.Print("W")
				break
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	os.Remove("go.sock")
	l, err := net.Listen("unix", "go.sock")
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
