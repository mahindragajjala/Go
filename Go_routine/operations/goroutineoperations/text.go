package goroutineoperations

/*
🔹 Basic Level
What is a Goroutine?
How to start a Goroutine using go keyword
Goroutines vs Functions (normal call vs concurrent call)
Anonymous functions with Goroutines
Main Goroutine and Program Termination
Multiple Goroutines and Scheduling (basic concept)

Race Condition Introduction

Wait using time.Sleep() (not recommended for sync, but for demo)

Performance benefits of Goroutines

🔹 Intermediate Level
Sync using sync.WaitGroup

Common mistakes with WaitGroup (e.g., defer inside loop)

Shared memory issues with Goroutines

Mutex (sync.Mutex) for synchronization

Deadlock Introduction with Mutex

Recursive function with Goroutines

Panic recovery in Goroutines

Worker Pool Pattern using Goroutines

Benchmarking Goroutines performance

🔹 Advanced Level
sync.RWMutex – Read/Write lock

sync.Once – One-time initialization

sync.Map – Concurrent-safe map

runtime.GOMAXPROCS and its effect on performance

Goroutine Leaks – What they are and how to detect

Goroutine Dump & Debugging (using runtime/pprof, runtime.Stack)

Controlling number of Goroutines (Semaphore pattern)

Goroutine Lifecycle Management with Context

Cancellation of Goroutines (via context.WithCancel)

Goroutines in Microservices (Real-world concurrency design)
*/
