package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = func(done <-chan any, filePath string) (<-chan any, <-chan error) {
	lines := make(chan any)
	errs := make(chan error, 1)

	go func() {
		defer close(lines)
		defer close(errs)

		file, err := os.Open(filePath)
		if err != nil {
			errs <- err
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			select {
			case <-done:
				fmt.Println("done reading file")
				return
			default:
				lines <- scanner.Text()
			}
		}
		if err := scanner.Err(); err != nil {
			errs <- err
		}

	}()
	return lines, errs
}

var writer = func(done <-chan any, data <-chan any, errs <-chan error) (<-chan any, error) {
	lines := make(chan any)
	go func() {
		defer close(lines)
		for {
			select {
			case <-done:
				fmt.Println("done writing to channel")
				return
			case err, ok := <-errs:
				if ok && err != nil {
					fmt.Println("error from reader", err)
				}
				return
			case line, ok := <-data:
				if !ok {
					fmt.Println("input channel closed")
					return
				}
				select {
				case <-done:
					fmt.Println("done writing to channel")
					return
				case lines <- strings.ToUpper(line.(string)):
				}
			}
		}

	}()
	return lines, nil
}

func main() {
	done := make(chan any)
	readerChan, errs := reader(done, "/Users/divya.thakur/Desktop/nutanix/concurrency_in_go/concurrency_patterns/pipelines/generator/file.txt")

	writerChan, err := writer(done, readerChan, errs)
	if err != nil {
		fmt.Println(err)
		return
	}
	counter := 0
	for line := range writerChan {
		if counter == 10 {
			close(done)
			time.Sleep(2 * time.Second)
			return
		}
		fmt.Println(line)
		counter++
	}
	close(done)
	fmt.Println("done final processing")
}
