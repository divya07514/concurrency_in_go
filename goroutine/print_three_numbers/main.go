package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	one := make(chan int, 1)
	two := make(chan int, 1)
	three := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	limit := 12
	done := make(chan struct{})

	go func(name string) {
		defer wg.Done()
		for {
			select {
			case <-done:
				println("Exiting", name)
				return
			case num := <-one:
				if num > limit {
					close(done)
					return
				}
				fmt.Println(name, num)
				time.Sleep(500 * time.Millisecond)
				two <- num + 1
			}
		}
	}("thread-1:")

	go func(name string) {
		defer wg.Done()
		for {
			select {
			case <-done:
				println("Exiting", name)
				return
			case num := <-two:
				if num > limit {
					close(done)
					return
				}
				fmt.Println(name, num)
				time.Sleep(500 * time.Millisecond)
				three <- num + 1
			}
		}
	}("thread-2:")

	go func(name string) {
		defer wg.Done()
		for {
			select {
			case <-done:
				println("Exiting", name)
				return
			case num := <-three:
				if num > limit {
					close(done)
					return
				}
				fmt.Println(name, num)
				time.Sleep(500 * time.Millisecond)
				one <- num + 1
			}
		}
	}("thread-3:")

	one <- 1
	wg.Wait()
	fmt.Println("Done")
}
