package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(30)

	for i := 0; i <= 9; i++ {
		go func() {
			wg.Done()
			fmt.Print(i, ", ")
		}()

	}

	for i := 0; i <= 9; i++ {
		go func(i int) {
			wg.Done()
			fmt.Print(i, ", ")
		}(i)
	}

	for i := 0; i <= 9; i++ {
		i := i
		go func() {
			wg.Done()
			fmt.Print(i, ", ")
		}()
	}

	wg.Wait()
}
