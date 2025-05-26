package workerpool

import (
	"fmt"
	"time"
)

// Job holds the number to process
type Job struct {
	ID     int
	Number int
}

// Result holds the result of processing
type Result struct {
	Job    Job
	Square int
}

// Worker: picks jobs from the jobs channel, processes them, and sends results to the results channel
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		// Simulate work
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(time.Second)
		output := Result{Job: job, Square: job.Number * job.Number}
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
		results <- output
	}
}

func Worker_Pool() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start worker pool
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: i, Number: i}
	}
	close(jobs)

	// Receive results
	for i := 1; i <= numJobs; i++ {
		result := <-results
		fmt.Printf("Result: JobID %d, Number %d, Square %d\n",
			result.Job.ID, result.Job.Number, result.Square)
	}
}
