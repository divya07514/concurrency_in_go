package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	limit, err := timeLimit(doSomeWork, 2*time.Second)
	fmt.Println(limit, err)
}

func timeLimit[T any](worker func() T, timeout time.Duration) (T, error) {
	out := make(chan T, 1)
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	go func() {
		out <- worker()
	}()

	select {
	case <-ctx.Done():
		var zero T
		return zero, errors.New("timeout")
	case v := <-out:
		return v, nil
	}

}

func doSomeWork() int {
	if x := rand.Int(); x%2 == 0 {
		return x
	} else {
		time.Sleep(10 * time.Second)
		return 100
	}
}
