package code

/*
Goroutine scheduling

Goroutine scheduling in Go is a lightweight, efficient way to manage "concurrency".

It's handled by "Go's runtime scheduler", which is designed to "multiplex thousands or even millions" of goroutines onto a small number of operating system (OS) threads.

Goroutines are managed by the "Go runtime scheduler", which maps multiple goroutines onto a small number of OS threads
This is achieved using a model called M scheduling, where "M goroutines are multiplexed onto N OS threads".


1. Goroutine (G)
A goroutine is a lightweight, user-space thread managed by the Go runtime.
Created using the go keyword.
Goroutines are much cheaper than OS threads in terms of memory and performance.

These represent individual functions or tasks that run concurrently within a Go application.

2. M (Machine)
Represents an OS thread.
It’s what actually executes the code on a CPU.
This represents an OS thread that is capable of executing goroutines.
Each thread has its associated CPU state, memory, and stack, as managed by the OS.

3. P (Processor)
A logical processor, not a CPU core.
It holds the run queue of goroutines.
Number of Ps = GOMAXPROCS (defaults to the number of logical CPUs).
Each P is assigned to one M at a time.

A logical resource that holds a “run queue” of goroutines that need execution.
Processors are the bridge between goroutines and OS threads, with a one-to-one mapping to CPU cores by default (controlled by the GOMAXPROCS environment variable).
The scheduler assigns each P to an OS thread (M), which then executes goroutines from P’s run queue.

Goroutine States

_Gidle: Just allocated but not initialized
_Grunnable: Ready to be executed but not currently running
_Grunning: Currently executing on an M
_Gsyscall: Executing a system call
_Gwaiting: Blocked, waiting for some condition
_Gdead: No longer used, waiting to be reused
_Gcopystack: Stack is being moved to a new location
_Gpreempted: Preempted during execution

| Concept     | Real-life Pizza Example | In Scheduler                        |
| ----------- | ----------------------- | ----------------------------------- |
| Function    | Pizza order             | Work unit / goroutine               |
| Thread      | Chef                    | OS thread                           |
| Scheduler   | Kitchen manager         | Your scheduler logic                |
| System call | Hiring/firing a chef    | Creating/deleting threads           |
| Core        | Stove/oven              | CPU core (runs one thing at a time) |


https://medium.com/@sanilkhurana7/understanding-the-go-scheduler-and-looking-at-how-it-works-e431a6daacf


https://leapcell.medium.com/gos-concurrency-decoded-goroutine-scheduling-594e67ac8e9d


*/
