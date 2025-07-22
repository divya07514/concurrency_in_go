package main

import (
	"math/rand"
	"time"
)

// If a goroutine is responsible for creating a goroutine, it is also responsible for ensur‐
// ing it can stop the goroutine.
func main() {

	numStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer close(randStream)
			defer println("rand stream closed")
			for {
				select {
				case <-done:
					return
				default:
					randStream <- rand.Intn(100)
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := numStream(done)
	println("generating 3 random numbers")
	for i := 0; i < 3; i++ {
		println(<-randStream)
	}
	close(done)
	time.Sleep(time.Second)
	println("done")
}
