package main

import (
	"concurrency_in_go/concurrency_patterns/pipelines/utils"
	"fmt"
)

func main() {
	genVals := func() <-chan <-chan any {
		chanStream := make(chan (<-chan any))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan any, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range utils.Bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
