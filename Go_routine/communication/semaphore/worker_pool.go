package semaphore

import (
	"fmt"
	"time"
)

/*
Bounded Worker Pool

Task:
Create a worker pool using a counting semaphore to limit the number of concurrent jobs running to 4, even if 20 jobs are submitted.

Goal:
Limit goroutine concurrency using semaphores instead of wait groups or worker channels.
*/
func Worker_pool() {
	jobs := 20
	semaphore := make(chan struct{}, 4)

	for i := 1; i <= jobs; i++ {
		semaphore <- struct{}{} // acquire token
		go func(id int) {
			fmt.Printf("Working on job %d\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("Finished job %d\n", id)
			<-semaphore // release token
		}(i)
	}

	// Wait until all jobs complete
	for i := 0; i < cap(semaphore); i++ {
		semaphore <- struct{}{}
	}
}

/*
//
func WorkerPoolCSP() {
	jobs := make(chan int, 20) // Buffered job queue
	done := make(chan bool)

	// Start 4 fixed workers
	for w := 1; w <= 4; w++ {
		go worker(w, jobs, done)
	}

	// Send 20 jobs
	for j := 1; j <= 20; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs

	// Wait for 4 workers to finish
	for w := 1; w <= 4; w++ {
		<-done
	}
}

func worker(id int, jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d at %v\n", id, job, time.Now())
		time.Sleep(5 * time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d at %v\n", id, job, time.Now())
	}
	done <- true // Signal that this worker is done
} */
