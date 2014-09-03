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

func echoServer(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 512)
	c.SetReadDeadline(time.Now().Add(time.Millisecond * 10))
	for {
		nr, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Closed: " + err.Error())
			return
		}
		_, err = c.Write(buf[:nr])
		if err != nil {
			return
		}

	}
}

func main() {
	l, err := net.Listen("unix", "go.sock")
	defer l.Close()
	defer os.Remove("go.sock")
	godbc.Check(err == nil)

	for {
		fd, err := l.Accept()
		godbc.Check(err == nil)

		go echoServer(fd)
	}
}
