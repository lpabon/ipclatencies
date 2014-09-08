package main

import (
	"fmt"
	"github.com/lpabon/foocsim/utils"
	"time"
)

func main() {
	var td utils.TimeDuration
	for i := 0; i < 10000000; i++ {
		start := time.Now()
		end := time.Now()

		td.Add(end.Sub(start))
	}

	fmt.Printf("Time should be close to 0: %.4f ns\n", td.MeanTimeUsecs()*1000.0)
}
