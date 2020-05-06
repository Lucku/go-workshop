package main

import (
	"fmt"
	"strings"
	"sync"
)

type job struct {
	message string
}

type result struct {
	processed string
}

func worker(id int, wg *sync.WaitGroup, jobs <-chan *job, results chan<- *result, errors chan<- error) {

	// fmt.Printf("Worker with ID %d started\n", id)
	// defer fmt.Printf("Worker with ID %d stopped\n", id)

	for j := range jobs {

		if j.message == "" {
			errors <- fmt.Errorf("message is empty")
			wg.Done()
			return
		}

		// Actual logic of worker
		res := processSingleMessage(j.message)

		results <- &result{processed: res}

		wg.Done()
	}
}

func processSingleMessage(msg string) string {
	sb := strings.Builder{}
	sb.WriteString(msg)
	sb.WriteString(" processed")
	return sb.String()
}

func processMessages(messages []string, np int) (resultsList []string) {

	jobs := make(chan *job, len(messages))
	results := make(chan *result, len(messages))
	errors := make(chan error, len(messages))

	var wg sync.WaitGroup

	for w := 1; w <= np; w++ {
		go worker(w, &wg, jobs, results, errors)
	}

	wg.Add(len(messages))

	for _, m := range messages {
		jobs <- &job{message: m}
	}

	close(jobs)

	wg.Wait()

	select {
	case err := <-errors:
		fmt.Println("Finished with error:", err)
	default:
	}

	close(errors)
	close(results)

	for r := range results {
		resultsList = append(resultsList, r.processed)
	}

	return resultsList
}

func main() {

	messages := []string{"These", "are" ,"a" ,"lot" ,"of" ,"messages"}

	result := processMessages(messages, 2)

	fmt.Println(result)
}
