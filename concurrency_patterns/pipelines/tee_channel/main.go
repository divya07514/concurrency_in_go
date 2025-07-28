package main

import (
	"concurrency_in_go/concurrency_patterns/pipelines/utils"
	"fmt"
)

func main() {
	done := make(chan any)
	defer close(done)

	out1, out2 := utils.Tee(done, utils.Take(done, utils.Repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
