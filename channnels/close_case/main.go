package main

import "fmt"

func main() {
	one := make(chan int)
	two := make(chan int)

	go func() {
		for i := 10; i < 100; i = i + 10 {
			one <- i
		}
		close(one)
	}()

	go func() {
		for i := 20; i >= 0; i-- {
			two <- i
		}
		close(two)
	}()

	result := readFromTwoChannels(one, two)
	fmt.Println(result)
}

func readFromTwoChannels(one chan int, two chan int) interface{} {
	var out []int
	for one != nil || two != nil {
		select {
		case v, ok := <-one: // channel nil pattern
			if !ok {
				one = nil
				print("one channel closed\n")
				continue
			}
			out = append(out, v)
		case v, ok := <-two:
			if !ok {
				two = nil
				print("two channel closed\n")
				continue
			}
			out = append(out, v)
		}
	}
	return out
}
