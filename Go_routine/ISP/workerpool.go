package isp

import "fmt"

/*
Sure! Let me explain this in simple English with a real-life example.



 🎯 Real-Life Situation

Imagine you run a photo editing website. People from all over the world upload pictures. Your system needs to:

1. Receive each uploaded image.
2. Resize or edit the image.
3. Send the result back to the user.



 👨‍🍳 Kitchen Analogy

Think of it like a restaurant kitchen:

* 🧾 The waiter brings orders from customers.
* 👨‍🍳 The chefs cook the food.
* 📣 The delivery staff serves the food to the customer.

Each person has one job.



 🧵 What is a Worker Pool?

A worker pool is a group of workers (in Go, goroutines) that do the same task in parallel. For example:

* 5 chefs (goroutines) work on 10 orders at the same time — each chef takes one order when they are free.



 ✅ Applying ISP (Interface Segregation Principle)

Instead of giving one person all the responsibilities (take order, cook, serve), we split it:

| Role     | What it does              | Go channel used                  |
| -- | - | -- |
| Waiter   | Sends orders              | `chan<- Job` (send-only)         |
| Chef     | Takes order and cooks it  | `<-chan Job` and `chan<- Result` |
| Delivery | Takes result and delivers | `<-chan Result` (receive-only)   |

Each part of the program only does what it’s supposed to — nothing extra.

This is the Interface Segregation Principle:

> Don't make a part of the system do more than what it needs to. Split responsibilities clearly.



 💻 Example in Go (Simple Explanation)

Here’s how you can imagine this in Go:


// Job is like a food order
type Job struct {
	ID int
}

// Result is like the finished dish
type Result struct {
	JobID int
	Msg   string
}

// Waiter: only sends jobs
func sendJobs(jobs chan<- Job) {
	for i := 1; i <= 5; i++ {
		jobs <- Job{ID: i}
	}
	close(jobs)
}

// Chef: gets job and sends back result
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		results <- Result{
			JobID: job.ID,
			Msg:   fmt.Sprintf("Worker %d processed Job %d", id, job.ID),
		}
	}
}

// Delivery: receives the result and shows message
func receiveResults(results <-chan Result) {
	for res := range results {
		fmt.Println(res.Msg)
	}
}

*/
// Job is like a food order
type Job struct {
	ID int
}

// Result is like the finished dish
type Result struct {
	JobID int
	Msg   string
}

// Waiter: only sends jobs
func sendJobs(jobs chan<- Job) {
	for i := 1; i <= 5; i++ {
		jobs <- Job{ID: i}
	}
	close(jobs)
}

// Chef: gets job and sends back result
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		results <- Result{
			JobID: job.ID,
			Msg:   fmt.Sprintf("Worker %d processed Job %d", id, job.ID),
		}
	}
}

// Delivery: receives the result and shows message
func receiveResults(results <-chan Result) {
	for res := range results {
		fmt.Println(res.Msg)
	}
}
