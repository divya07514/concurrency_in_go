package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Error    error
	Response *http.Response
}

var checkStatus = func(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)
		for _, url := range urls {
			var result Result
			resp, err := http.Get(url)
			result = Result{
				Error:    err,
				Response: resp,
			}
			select {
			case results <- result:
			case <-done:
				return
			}
		}
	}()
	return results
}

func main() {
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "a", "b", "c", "d"}
	counter := 0
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("Error: %s\n", result.Error.Error())
			counter++
			if counter >= 3 {
				fmt.Println("Stopping after 3 results")
				break
			}
			continue
		}
		fmt.Println("Response Status:", result.Response.Status)
	}
	time.Sleep(time.Second)
	fmt.Println("done")
}
