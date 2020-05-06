package main

import (
	"fmt"
	"math/rand"
	"time"
)

func square(inputs <-chan float64, results chan<- float64) {

	factor := <-inputs

	// Unpredictable execution time
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

	result := factor * factor

	results <- result
}

func main() {

	defer func(currentTime time.Time) {
		passedTime := time.Since(currentTime)
		fmt.Println("Time passed:", passedTime)
	}(time.Now())

	results := make(chan float64)
	inputs := make(chan float64)

	defer close(inputs)

	go square(inputs, results)

	inputs <- float64(10)

	/*

	Selects are quite similar to switches:

		switch something {
		case 1: // do something
		case 2: // do something else
		default: // do something completely different
		}
	*/

	timeout := 1 * time.Second

	timeoutChannel := time.After(timeout)

	select {
		case r := <-results:
			fmt.Println(r)
			close(results)
		case <-timeoutChannel:
			fmt.Println("Timeout occurred")
	}
}
