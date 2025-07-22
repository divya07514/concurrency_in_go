package main

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i < 5; i++ {
				resultStream <- i * i
			}
		}()
		return resultStream
	}
	resultStream := chanOwner()
	for data := range resultStream {
		println(data)
	}
	println("done")
}
