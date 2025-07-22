package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0, 10)

	deque := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		item := queue[0]
		queue = queue[1:]
		fmt.Print("Dequeued: ", item, "\n")
		c.L.Unlock()
		c.Signal() // Notify one waiting goroutine
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		queue = append(queue, i)
		fmt.Print("Enqueued: ", i, "\n")
		go deque(time.Second)
		c.L.Unlock()
	}
}
