package main

import (
	"concurrency_in_go/concurrency_patterns/pipelines/utils"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand := func() any { return rand.Intn(500000000-200000000+1) + 200000000 }
	done := make(chan any)
	defer close(done)

	start := time.Now()
	randIntStream := utils.ToInt(done, utils.RepeatFn(done, rand))

	fmt.Println("Primes:")

	nums := runtime.NumCPU()
	finders := make([]<-chan any, nums)
	for i := range nums {
		finders[i] = utils.PrimeFinder(done, randIntStream)
	}

	primeStream := utils.FanIn(done, finders...)

	for prime := range utils.Take(done, primeStream, 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(start))
}
