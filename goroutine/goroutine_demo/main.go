package main

import (
	"fmt"
	"time"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := processConcurrently(x)
	fmt.Printf("%v\n", result)
}

func processConcurrently(x []int) []int {
	routines := 5
	in := make(chan int, routines)
	out := make(chan int, routines)

	for i := 0; i < routines; i++ {
		go func() {
			for val := range in {
				result := doubleIt(val)
				out <- result
				print("writing to out:", result)
				print("\n")
				time.Sleep(2 * time.Second) // Simulate some processing time
			}
		}()
	}

	go func() {
		for val := range x {
			in <- val
			print("reading from x:", val)
			print("\n")
			time.Sleep(1 * time.Second)
		}
	}()

	output := make([]int, 0, len(x))
	for i := 0; i < len(x); i++ {
		output = append(output, <-out)
	}
	return output
}

func doubleIt(val int) int {
	return val * 2
}
