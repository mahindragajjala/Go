package goscheduler

/*

* G (Goroutine): Represents a lightweight thread (task or function).
* P (Processor): Holds a local queue of Gs. It is responsible for scheduling Gs to Ms.
* M (Machine): An actual OS thread that runs the Gs.

Assignment Workflow:

1. Goroutine creation (G):

   * When a new goroutine is created (via `go func()`), it is placed in the local queue of the current P.
   * If the local queue is full, it's added to the global run queue.

2. Execution begins:

   * Each P is bound to an M (OS thread) and schedules goroutines from its local queue.
   * The M calls the Go scheduler, which picks a G from P's queue and starts executing it.

3. Running Gs:

   * The M continues to pick Gs from the local queue of its assigned P.
   * If P's local queue is empty, it can:

     * Steal Gs from other Ps (work stealing).
     * Pull Gs from the global run queue.

4. Blocking calls:

   * If a G does a blocking syscall (like I/O), the M can block.
   * The scheduler unbinds the M from the P, assigns a new M to the P, and continues running other Gs.

5. GOMAXPROCS:

   * This sets the maximum number of Ps, which limits the number of goroutines that can run in parallel.
   * Default is the number of CPU cores.



 💡 Visualization:


G (Goroutine) -> goes into -> P (Processor)'s queue -> scheduled on -> M (OS Thread)


 Example:

If `GOMAXPROCS = 2`, there are:

* 2 Ps (each with a queue of Gs).
* Multiple Ms created as needed, but only 2 M-P pairs can run Gs concurrently.


*/
