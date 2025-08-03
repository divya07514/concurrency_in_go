package main

import (
	"fmt"
	"log"
	"os"

	errorpropagation "concurrency_in_go/concurrency_patterns/at_scale/error_propagation"
)

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}
func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	err := errorpropagation.RunJob("1")
	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."
		if _, ok := err.(errorpropagation.IntermediateErr); ok {
			msg = err.Error()
		}
		handleError(1, err, msg)
	}
}
