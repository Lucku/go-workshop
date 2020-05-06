package main

import (
	"fmt"
	"sync"
)

func count(i int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Count:", i)
}

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go count(i, &wg)
		wg.Add(1)
	}

	wg.Wait()

	// If not waiting:
	// main goroutine: 		XXXXX| STOP
	// count goroutine: 	----X|XXX (Last part of code can't be executed)
}
