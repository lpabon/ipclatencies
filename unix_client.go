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
	"github.com/lpabon/foocsim/utils"
	"net"
	"sync"
	"time"
)

func client(c net.Conn, wg *sync.WaitGroup, id int) {

	defer wg.Done()

	defer c.Close()

	buf := make([]byte, 4096)
	var td utils.TimeDuration
	prev := td.Copy()

	for i := 0; i < 100000; i++ {

		start := time.Now()

		_, err := c.Write(buf)
		if err != nil {
			fmt.Printf("%d: %d Error writing: %s\n", id, errors, err.Error())
			return
		}

		_, err = c.Read(buf)
		if err != nil {
			fmt.Printf("%d: Error writing: %s\n", id, err.Error())
			return
		}

		td.Add(time.Now().Sub(start))

		if (i % 10000) == 0 {
			fmt.Printf("%d:%d: %.4f usecs\n", id, i, td.DeltaMeanTimeUsecs(prev))
			prev = td.Copy()
		}

		time.Sleep(time.Microsecond * 10)
	}

	fmt.Printf("%d:Latency %.4f usecs\n", id, td.MeanTimeUsecs())
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {

		c, err := net.Dial("unix", "go.sock")
		if err != nil {
			fmt.Printf("Last thread: %d\n", i)
			fmt.Printf("%d: Error: %s\n", err.Error())
			break
		}

		wg.Add(1)
		go client(c, &wg, i)
	}

	wg.Wait()
}
