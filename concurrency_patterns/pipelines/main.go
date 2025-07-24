package main

import "fmt"

var generator = func(done <-chan any, integers ...int) <-chan int {
	intStream := make(chan int, len(integers))
	go func() {
		defer close(intStream)
		for _, in := range integers {
			select {
			case <-done:
				return
			case intStream <- in:
			}
		}
	}()
	return intStream
}

var multiply = func(done <-chan any, intStream <-chan int, multiplier int) <-chan int {
	mulStream := make(chan int)
	go func() {
		defer close(mulStream)
		for v := range intStream {
			select {
			case <-done:
				return
			case mulStream <- v * multiplier:
			}
		}
	}()
	return mulStream
}

var add = func(done <-chan any, intStream <-chan int, additive int) <-chan int {
	addStream := make(chan int)
	go func() {
		defer close(addStream)
		for v := range intStream {
			select {
			case <-done:
				return
			case addStream <- v + additive:
			}
		}
	}()
	return addStream
}

func main() {
	done := make(chan any)
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}

}
