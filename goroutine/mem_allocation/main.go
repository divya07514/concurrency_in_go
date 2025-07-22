package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { // this blocks forever
		wg.Done()
		<-c
	}

	const numGoroutines = 1e5
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := 0; i < numGoroutines; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
