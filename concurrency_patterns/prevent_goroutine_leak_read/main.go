package main

import (
	"fmt"
	"time"
)

// If a goroutine is responsible for creating a goroutine, it is also responsible for ensur‐
// ing it can stop the goroutine.
func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer close(terminated)
			defer fmt.Print("do work terminated\n")
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("Cancelling do work goroutine...")
		close(done)
	}()
	<-terminated
	fmt.Println("Done")
}
