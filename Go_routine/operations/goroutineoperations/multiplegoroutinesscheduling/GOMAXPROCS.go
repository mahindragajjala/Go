package multiplegoroutinesscheduling

import (
	"fmt"
	"runtime"
	"time"
)

/*
GOMAXPROCS sets the maximum number of OS threads that can execute Go code simultaneously. It tells the Go scheduler how many logical CPUs it can use to run goroutines in parallel.

This does not limit how many goroutines you can create. You can spawn thousands of goroutines, but GOMAXPROCS controls how many of them can run at the same time.

By default, GOMAXPROCS is set to the number of logical CPUs on your system.
*/
func GOMAXPROCS() {
	fmt.Println(runtime.GOMAXPROCS(0)) // Get current value
}

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func Limited_GOMAXPROCS() {
	/*
		Limits the Go runtime to only one OS thread.
		Meaning: only one goroutine can execute Go code at a time, even if you have many CPU cores.
		This simulates a single-core environment. */

	runtime.GOMAXPROCS(1)              // Limit Go to 1 OS thread (CPU core)
	fmt.Println(runtime.GOMAXPROCS(1)) // Get current value
	/*
		go keyword creates two goroutines to run task() concurrently.
		They're scheduled to run by the Go runtime, but due to GOMAXPROCS(1), only one can run at a time.
	*/
	go task("Goroutine-1")
	go task("Goroutine-2")

	time.Sleep(1 * time.Second) // Wait to let goroutines complete

	fmt.Println("Main Done")
}

//Using All CPU Cores

func heavyTask(id int) {
	start := time.Now()
	sum := 0
	for i := 0; i < 1e8; i++ {
		sum += i
	}
	fmt.Printf("Task %d completed in %v\n", id, time.Since(start))
}

func Using_All_CPU_Cores() {
	cores := runtime.NumCPU()
	fmt.Println("Available CPUs:", cores)

	runtime.GOMAXPROCS(cores)

	for i := 1; i <= cores; i++ {
		go heavyTask(i)
	}

	time.Sleep(3 * time.Second)
}

/*
Write a program that prints even and odd numbers using two separate goroutines.
Create two goroutines that print "Ping" and "Pong" alternatively using time.Sleep(). - polling
Write a program that starts 5 goroutines, each printing a number from 1 to 5.

Create a goroutine that counts up to 10 and a main routine that waits for it to finish.

🧪 Intermediate Level
Write a function that spawns 10 goroutines, each computing the square of a number.

Use an anonymous goroutine to print "Hello from anonymous goroutine".

Create a program where 3 goroutines update a shared variable without synchronization. Observe the issue.

Fix the above problem using sync.Mutex or sync/atomic.

🧪 Advanced Level
Simulate a CPU-bound task (e.g., prime number check) across multiple goroutines and observe scheduling with GOMAXPROCS.

Create 10 goroutines that all send messages to a channel. Use a WaitGroup to wait until all are done.

Use runtime.Gosched() to yield control from one goroutine to others and explain behavior.

Observe and demonstrate the Go scheduler’s fairness by running infinite loops in multiple goroutines.
*/
