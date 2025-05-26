package workerpool

/*
A way to reuse a fixed number of goroutines to do a lot of work, instead of creating a new goroutine every time.

In Golang, a Worker Pool (also known as a Goroutine Pool) is a "concurrency pattern" used to "manage and limit the number of goroutines" (lightweight threads) running at the same time.

It is very useful when you want to process a large number of tasks (jobs) efficiently without overwhelming the system.

*/

/*
Why Use a Worker Pool?

Go can "create" or "start" thousands of goroutines, but spawning too many at once may consume too much memory/CPU.

A Worker Pool allows:
Controlled parallel execution
Better resource utilization
Improved system stability
*/
/*
Real-Life Analogy:
Imagine you're a manager at a car wash center. You have:
100 dirty cars waiting outside (tasks)
5 employees (workers) who can clean cars in parallel
You don't want to hire 100 employees because it wastes resources. Instead, you use only 5 employees in a loop to clean one car at a time. That’s how a Worker Pool works.
*/

/*
Components of Worker Pool:

Jobs Channel – List of tasks to process.
Worker Goroutines – A fixed number of workers that process the tasks.
Results Channel – Optional. Stores the results from each job.
Main Function – Sends jobs and collects results.
*/

/*
✅ What is a Semaphore?

A semaphore is a low-level synchronization primitive used to control access to a limited resource by multiple goroutines.
Think of it as a counter:
When a goroutine wants access, it decrements the counter.
When done, it increments it.
If the counter is 0, other goroutines must wait.
In Go, semaphores are often implemented using buffered channels.

✅ What is a Worker Pool?

A Worker Pool is a design pattern:
It creates a fixed number of worker goroutines
Distributes many tasks to these workers via channels
Each worker processe	s one task at a time
Think of it as a task queue + limited workers
*/
/*
A semaphore is a concurrency control mechanism that limits how many goroutines can access a shared resource at the same time. It helps prevent overuse of resources like memory, database connections, etc.

A goroutine pool (or worker pool) is a design pattern where a fixed number of goroutines are created to process many tasks. It helps limit the number of goroutines running at the same time, which reduces memory and CPU usage.
*/
