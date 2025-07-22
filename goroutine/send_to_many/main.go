package main

import (
	"fmt"
	"time"
)

var sendToMany = func(done <-chan interface{}, ch <-chan int, threadName string) {
	go func() {
		for {
			select {
			case <-done:
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Printf("%s: %d\n", threadName, v)
			}
		}
	}()
}

func main() {
	done := make(chan interface{})
	channels := make([]chan int, 5)

	for i := 0; i < 5; i++ {
		ch := make(chan int)
		threadName := fmt.Sprintf("Thread-%d", i)
		sendToMany(done, ch, threadName)
		channels[i] = ch
	}
	for i := 0; i < 3; i++ {
		for _, ch := range channels {
			ch <- i
		}
	}
	for _, ch := range channels {
		close(ch)
	}
	close(done)
	time.Sleep(time.Second)
	fmt.Println("done")
}
