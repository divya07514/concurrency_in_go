package main

import (
	"fmt"
	"time"
)

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	ch := make(chan int)
	go func() {
		for _, val := range vals {
			ch <- val
		}
	}()
	result := process(ch)
	fmt.Println("Result:", result)
}

func process(ch chan int) []int {
	nums := 10
	out := make(chan int, nums)
	// Using a buffered channel to allow goroutines to run concurrently
	// only 10 results will be processed at a time even though ch contains more values

	for i := 0; i < nums; i++ {
		go func() {
			v := <-ch
			time.Sleep(100 * time.Millisecond) // Simulate some processing delay
			out <- v * v
		}()
	}
	result := make([]int, 0)
	for i := 0; i < nums; i++ {
		result = append(result, <-out)
	}
	return result
}
