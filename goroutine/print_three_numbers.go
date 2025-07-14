package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	one := make(chan int)
	two := make(chan int)
	three := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)
	limit := 66

	go func(name string) {
		defer wg.Done()
		for {
			num, ok := <-one
			if !ok {
				return
			}
			if num > limit {
				close(two)
				close(three)
				return
			}
			println(name, num)
			time.Sleep(500 * time.Millisecond) // Simulate some work
			two <- num + 1
		}
	}("thread-1:")

	go func(name string) {
		defer wg.Done()
		for {
			num, ok := <-two
			if !ok {
				return
			}
			if num > limit {
				close(three)
				close(one)
				return
			}
			println(name, num)
			time.Sleep(500 * time.Millisecond) // Simulate some work
			three <- num + 1
		}
	}("thread-2:")

	go func(name string) {
		defer wg.Done()
		for {
			num, ok := <-three
			if !ok {
				return
			}
			if num > limit {
				close(one)
				close(two)
				return
			}
			println(name, num)
			time.Sleep(500 * time.Millisecond) // Simulate some work
			one <- num + 1
		}
	}("thread-3:")
	one <- 1
	wg.Wait()
	fmt.Println("Done")
}
