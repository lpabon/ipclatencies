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

func client(wg sync.WaitGroup, id int) {

	defer wg.Done()

	c, _ := net.Dial("unix", "go.sock")
	defer c.Close()

	buf := make([]byte, 8)
	var td utils.TimeDuration
	prev := td.Copy()

	for i := 0; i < 100000000; i++ {

		start := time.Now()

		c.Write(buf)
		c.Read(buf)

		end := time.Now()
		td.Add(end.Sub(start))

		if (i % 10000) == 0 {
			fmt.Printf("%d:%d: %.4f usecs\n", id, i, td.DeltaMeanTimeUsecs(prev))
			prev = td.Copy()
		}

		time.Sleep(time.Microsecond * 10)
	}

	fmt.Printf("\n%d:Latency %.4f usecs\n", 0, td.MeanTimeUsecs())
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		go client(wg, i)
		wg.Add(1)
	}

	wg.Wait()
}

/*
func main() {

	c, _ := net.Dial("unix", "go.sock")
	defer c.Close()

	buf := make([]byte, 8)
	var td utils.TimeDuration

	for i := 0; i < 100000; i++ {

		start := time.Now()

		c.Write(buf)
		c.Read(buf)

		end := time.Now()
		td.Add(end.Sub(start))
		time.Sleep(time.Microsecond * 10)

	}

	fmt.Printf("\n%d:Latency %.4f usecs\n", 0, td.MeanTimeUsecs())

}
*/
