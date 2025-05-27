package isp

import "fmt"

/*
Interface Segregation for Goroutines
Enforcing single responsibility (e.g., a sender only sends, a receiver only receives)

Useful in pipelines and worker pools


The Interface Segregation Principle (ISP) is one of the SOLID principles in software engineering. It states:

"No client should be forced to depend on methods it does not use."

Applied to goroutines and channels in Go, this means segregating responsibilities—clearly separating senders, receivers, and processors—so that each part of a concurrent system does one job, and can’t do more even by accident.

This is very powerful for writing safe, scalable, and maintainable concurrent Go code.
Segregating responsibilities, also known as separation of duties, is a fundamental principle in risk management and internal controls that involves distributing different tasks and responsibilities among multiple individuals.

Enforcing Single Responsibility
Instead of allowing a goroutine to both send and receive, you define roles:

Send-only goroutine (chan<- T)
Receive-only goroutine (<-chan T)
Processor goroutine that takes in a value, transforms it, and passes it to another channel


This is like breaking a large interface into smaller, focused interfaces — but with channels and goroutines.
*/
// --- Step 1: Generator ---
// Send-only channel (chan<- int)
func generate(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// --- Step 2: Square ---
// Receives from input (<-chan int), sends to output (chan<- int)
func square(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

// --- Step 3: Printer ---
// Receive-only channel (<-chan int)
func printResults(ch <-chan int) {
	for result := range ch {
		fmt.Println("Result:", result)
	}
}

func INTERFACE_SEGREGATION_FOR_GOROUTINE() {
	// Channels
	nums := make(chan int)
	squares := make(chan int)

	// Launch segregated goroutines
	go generate(nums)        // Only sends
	go square(nums, squares) // Processes
	printResults(squares)    // Only receives
}

/*
Let's explain the Interface Segregation Principle (ISP) with a real-time, real-world example, and then show how it applies to Go (especially goroutines).

🧠 Real-Time Example of Interface Segregation Principle (ISP)
🚗 Scenario: Car Interfaces

Let’s say you're designing a system for different types of vehicles:

* Electric Car (Tesla)
* Petrol Car (Maruti)
* Motorcycle
* Bicycle

Now imagine you define this interface:
type Vehicle interface {
	StartEngine()
	StopEngine()
	ChargeBattery()
	Refuel()
}
This violates ISP because:

* 🚲 Bicycle doesn’t have an engine, so it shouldn't implement `StartEngine`, `StopEngine`
* 🚗 Petrol Car doesn’t charge battery, so `ChargeBattery()` is irrelevant
* ⚡ Electric Car doesn't `Refuel()`

👉 They are forced to implement methods they don’t need. That’s a direct violation of ISP.

✅ ISP-Compliant Version

Break this big interface into smaller, more focused ones:

type Engine interface {
	StartEngine()
	StopEngine()
}

type Electric interface {
	ChargeBattery()
}

type Petrol interface {
	Refuel()
}


Now each class/struct only implements the interface(s) it needs.

 🔄 Applying This to Go Goroutines & Channels

Let’s relate this to Go’s concurrency model with goroutines and channels.

 🎯 Real-World Use Case: Order Processing System

Imagine a backend system like Amazon or Flipkart:

1. 📦 Order Receiver – receives new orders
2. 🏗️ Order Processor – verifies, processes order
3. 🧾 Invoice Generator – generates invoice and sends it to email system

 ❌ Anti-ISP Example (Violating ISP)

func orderHandler(ch chan string) {
	// receives order
	order := <-ch

	// processes order
	process(order)

	// generates invoice
	invoice(order)
}

* This single goroutine does everything — receive, process, generate invoice.
* It’s tightly coupled, hard to test, and hard to scale.

✅ ISP-Compliant Goroutine Design (Real-Time Parallelism)

Split them using ISP principles and Go’s channel directionality:


// Receive orders (Receive-only)
func receiveOrders(out chan<- string) {
	orders := []string{"Order#1", "Order#2", "Order#3"}
	for _, o := range orders {
		out <- o
	}
	close(out)
}

// Process orders (Receive from one, send to another)
func processOrders(in <-chan string, out chan<- string) {
	for order := range in {
		out <- "Processed-" + order
	}
	close(out)
}

// Generate invoices (Receive-only)
func generateInvoices(in <-chan string) {
	for invoice := range in {
		fmt.Println("Generated Invoice for:", invoice)
	}
}

func main() {
	orderChan := make(chan string)
	processedChan := make(chan string)

	go receiveOrders(orderChan)
	go processOrders(orderChan, processedChan)
	generateInvoices(processedChan)
}

 ✅ Real-Time System Result

This mimics a real-time pipeline:

* `receiveOrders` only sends (send-only `chan<-`)
* `processOrders` both receives and sends
* `generateInvoices` only receives (receive-only `<-chan`)

Each goroutine has a single, clear responsibility, in line with ISP.


 💡 Key Takeaways

* ISP = “Don’t give a function/goroutine more power than it needs.”
* Use send-only (`chan<-`) and receive-only (`<-chan`) to enforce responsibility boundaries.
* Great for pipelines, queues, microservices, and scalable systems.


*/
