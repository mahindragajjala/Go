package looponchannels

import "fmt"

//Loop with Goroutine Pool Pattern
func Pool_Pattern() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	worker := func(id int, jobs <-chan int, results chan<- int) {
		for job := range jobs {
			fmt.Printf("Worker %d processing job %d\n", id, job)
			results <- job * 2
		}
	}

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}

}

/*
Use case: Concurrent workers processing jobs in parallel.
*/
