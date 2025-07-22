package main

import (
	"fmt"
	"sync"
)

func main() {
	oddCh := make(chan int)
	evenCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	limit := 100

	go func() {
		defer wg.Done()
		for {
			num, ok := <-oddCh
			if !ok {
				return
			}
			if num > limit {
				close(evenCh)
				return
			}
			println(num)
			evenCh <- num + 1
		}
	}()

	go func() {
		defer wg.Done()
		for {
			num, ok := <-evenCh
			if !ok {
				return
			}
			if num > limit {
				close(oddCh)
				return
			}
			println(num)
			oddCh <- num + 1
		}
	}()

	oddCh <- 1

	wg.Wait()
	fmt.Print("number generation completed\n")

}
