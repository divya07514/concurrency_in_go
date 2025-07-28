package utils

import "sync"

var Take = func(
	done <-chan any,
	valueStream <-chan any,
	num int,
) <-chan any {
	takeStream := make(chan any)
	go func() {
		defer close(takeStream)
		for i := num; i > 0 || i == -1; {
			if i != -1 {
				i--
			}
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

var RepeatFn = func(
	done <-chan any,
	fn func() any,
) <-chan any {
	valueStream := make(chan any)
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

var ToInt = func(
	done <-chan any,
	valueStream <-chan any,
) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

func PrimeFinder(done <-chan any, intStream <-chan int) <-chan any {
	primeStream := make(chan any)
	go func() {
		defer close(primeStream)
		for candidate := range intStream {
			select {
			case <-done:
				return
			default:
				if isPrime(candidate) {
					select {
					case <-done:
						return
					case primeStream <- candidate:
					}
				}
			}
		}
	}()
	return primeStream
}

var FanIn = func(
	done <-chan any,
	channels ...<-chan any,
) <-chan any {
	var wg sync.WaitGroup
	multiplexedStream := make(chan any)

	mutiplex := func(c <-chan any) {
		defer wg.Done()
		for v := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- v:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go mutiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var Repeat = func(
	done <-chan any,
	values ...any,
) <-chan any {
	valueStream := make(chan any)
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

var OrDone = func(done, c <-chan any) <-chan any {
	valStream := make(chan any)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

var Tee = func(done <-chan any, in <-chan any) (<-chan any, <-chan any) {
	out1 := make(chan any)
	out2 := make(chan any)
	go func() {
		defer close(out1)
		defer close(out2)
		for v := range OrDone(done, in) {
			var out1, out2 = out1, out2
			for range 2 {
				select {
				case <-done:
					return
				case out1 <- v:
					out1 = nil
				case out2 <- v:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}
