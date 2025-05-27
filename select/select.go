package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		inRoutine := 1
		ch1 <- inRoutine
		fromMain := <-ch2
		fmt.Println("goroutine", inRoutine, fromMain) // this is never executed !!!
	}()

	inMain := 2
	var fromRoutine int
	select {
	case fromRoutine = <-ch1:
	case ch2 <- inMain:
	}

	fmt.Println("in main", inMain, fromRoutine)
	time.Sleep(2 * time.Second)
}
