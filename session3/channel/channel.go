package main

import (
	"fmt"
	"time"
)

func square(inputs <-chan float64, results chan<- float64) {

	// 1st synchronization point: Block until receiving a value from inputs
	factor := <-inputs

	time.Sleep(1 * time.Second)

	result := factor * factor

	// 2. out <- result
	results <- result
}

func main() {

	// Total time of n parallel square functions
	//
	// gr0: -------(time.Sleep)------------>
	// gr1: -------(time.Sleep)------->
	// gr2: --------(time.Sleep)----------------> | total execution time = max(execution time of grN) (N=0..9)
	// ...

	defer func(currentTime time.Time) {
		passedTime := time.Since(currentTime)
		fmt.Println("Time passed:", passedTime)
	}(time.Now())

	results := make(chan float64)
	inputs := make(chan float64)

	defer close(results)
	defer close(inputs)

	for i := 0; i < 10; i++ {
		go square(inputs, results)
	}

	for i := 0; i < 10; i++ {
		inputs <- float64(i)
	}

	// 2st synchronization point: Block 10 times until receiving a value from results
	for i := 0; i < 10; i++ {
		fmt.Println(<-results)
	}
}
