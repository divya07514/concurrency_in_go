package main

import (
	"fmt"
	"sync"
)

func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {
	out := make(chan R, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	var result []R
	for r := range out {
		result = append(result, r)
	}
	return result
}

func main() {
	in := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			in <- i
		}
		close(in)
	}()
	processor := func(i int) int {
		return i * 2
	}
	gather := processAndGather(in, processor, 3)
	fmt.Print(gather)
}
