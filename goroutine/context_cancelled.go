package main

import "context"

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 1; i <= max; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	to := countTo(ctx, 50)

	for i := range to {
		if i == 30 {
			break
		}
		print(i, " ")
	}

}
