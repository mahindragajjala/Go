package goscheduler

/*
What Are Run Queues?

In Go’s scheduler:
A run queue is a list of "ready-to-run goroutines" waiting to be picked up by a thread and run on a processor.

Go uses two types of run queues:

Local Run Queue (per-P)
Global Run Queue (shared)

1. Local Run Queue

"Each Processor(P) has its own private queue of goroutines."
These are goroutines scheduled to run on that specific processor.
This makes scheduling faster, because P doesn’t need to coordinate with others.
Think of it as a personal task list that a processor (P) handles on its own.

2. Global Run Queue

A shared queue for all goroutines.
Goroutines go here when there's no available P at the moment.
If a local queue is empty, a P may steal a goroutine from the global queue.
Think of it as a backup waiting area that any processor can use when it runs out of tasks.

go func() {
   ... // Your function
}
↓
Creates G (Goroutine)
↓
Checks P’s Local Run Queue (runqput)
   ↳ If room → placed in Local Run Queue
   ↳ If full → placed in Global Run Queue (globrunqput)
↓
Scheduler (schedule) picks G from:
   1. Local Queue
   2. Global Queue
   3. Steal from other P's local queue
↓
M (thread) runs G

------------------------------------------------------------------------------------------------

What Are Run Queues?

Run queues are internal structures used by the Go runtime to manage Goroutines that are ready to run. There are two main types of run queues in Go:

1. Global Run Queue: This is the global queue where Goroutines are placed when they are ready to run, but cannot immediately execute due to a lack of available processor (or CPU core).

2. Local Run Queue: Each OS thread (or P - processor in Go terms) has its own local queue where Goroutines are scheduled to run. The local queue holds Goroutines that are associated with the thread and can be executed immediately when that thread is available.

Go Scheduling System: G, P, and M

To understand the run queues better, let's first discuss the Go runtime scheduling system:

G (Goroutine): Represents the unit of execution. Each Goroutine is a lightweight thread.
P (Processor): Represents a logical processor. It’s essentially a resource for scheduling Goroutines. There is typically one P per CPU core (and also potentially more depending on the Go runtime configuration).
M (Machine): Represents an operating system thread. A single M can run a single P at a time.

The Go scheduler uses this G-M-P model to handle Goroutines efficiently.

Creating Goroutines and Placing Them in the Run Queue

1. Goroutine Creation:
   When a Goroutine is created (using `go` keyword), the Go runtime creates a `G` object for that Goroutine. This object holds metadata related to the Goroutine such as its state (running, waiting, etc.), the stack, and other runtime details.

2. Placing the Goroutine in the Run Queue:

   * If the Goroutine is ready to run, it is placed in a run queue.
   * If it’s created by an idle thread (with no Goroutines currently running on that thread), it may be placed directly in the local run queue of the thread (P).
   * If the local run queue is full or the thread is busy, the Goroutine will be placed in the "global run queue".

Local Run Queue

* Each P(Processor) has its own local run queue. This local queue stores the Goroutines that are ready to run on that specific processor.
* When a Goroutine is ready to run, the Go scheduler tries to pick it up from the local run queue of the processor associated with the Goroutine (if it's available).

  Flow:

  * The Goroutine is initially placed in the local run queue of an idle processor (P).
  * If a processor (P) is idle, it pulls Goroutines from its own local run queue to execute.
  * If the local queue is empty, the processor may check other queues (such as the global run queue or other P's local queue).

  type P struct {
    runq     [256]G      // local run queue (ring buffer)
    runqhead uint32
    runqtail uint32
    ...
}
📍 So: Local run queue is part of the P struct in the Go runtime.
-------------------------------------------------------------------------

Global Run Queue

* The global run queue is a "shared queue" that holds Goroutines that cannot be immediately executed by a specific processor (P). This may happen if all processors are busy or a processor’s local run queue is full.
* The global run queue allows Goroutines to be picked up by any available processor (P) once it becomes idle.

  Flow:

  * When a Goroutine is created and there is no available processor (P) to handle it, it is placed in the global run queue.
  * As processors (P) become idle, they may pull Goroutines from the global run queue and execute them.
  * This ensures that no Goroutine is left unexecuted, even when local run queues are full or processors are busy.


Located inside the Go Scheduler (runtime/proc.go), not tied to any particular P.
var (
    sched struct {
        ...
        runq     gQueue // global run queue
        runqsize int32
        ...
    }
)

It is a single, shared queue used when:
Local queues are full
Goroutines need to be redistributed
-------------------------------------------------------------------------
Balancing Between Local and Global Run Queues

The Go runtime tries to maintain a balance between local and global run queues:

1.Work Stealing:

   * If a local run queue becomes empty but there are still Goroutines in other local run queues, the processor (P) can steal work from another processor’s local run queue.
   * This helps keep the CPU cores busy and prevents idling.

2. Processor Idle Time:

   * When a processor (P) finishes executing its local run queue, it will check the global run queue for more work.
   * If the global run queue is also empty, the processor will try to steal work from other processors’ local queues.

3. Goroutine Scheduling:

   * If a Goroutine performs a "blocking operation" (e.g., waiting for I/O), it is moved to a "waiting state" and is eventually scheduled again when the blocking operation completes.
   * Such Goroutines are not placed in the run queues unless they are ready to run again.
-------------------------------------------------------------------------
A Simple Flow of Execution

Let’s break down the execution flow in simpler terms:

1. Goroutine Creation:

   * A Goroutine is created using the `go` keyword.
   * Initially, it’s added to the local run queue of the processor (P) handling that task.

2. Processor (P) Picks a Goroutine**:

   * If the processor (P) is idle, it picks the Goroutine from its local queue.
   * If the local queue is empty, the processor checks the global run queue for tasks.

3. Executing the Goroutine:

   * The Goroutine executes until it either completes or blocks (for example, if it waits for some resource like I/O).
   * If blocked, it is not in the run queue until it is ready to run again.

4. Idle Processor (P):

   If the processor (P) finishes its local run queue and the global run queue is empty, it might steal work from another processor’s local queue.

-------------------------------------------------------------------------
Let’s visualize this:

1. Goroutine Creation:

   go func() {
       fmt.Println("Hello from Goroutine!")
   }()

   A Goroutine is created and added to the local run queue of the processor (P) handling it.

2. Goroutine Execution:

   Processor (P) picks the Goroutine from the local run queue and begins executing it.

3. Processor Busy:

   * If all processors are busy, the new Goroutine is added to the global run queue.
   * Once a processor becomes idle, it picks a Goroutine from the global run queue.

4. **Completion**:

   * The Goroutine finishes its task and exits.
-------------------------------------------------------------------------
But since local run queues have limited size (default 256), once they’re full, the rest of the goroutines overflow into the global run queue.
As processors finish local tasks and become idle, they pick Goroutines from the global run queue.

-------------------------------------------------------------------------

Goroutine is created using go func().
The Go runtime puts the new Goroutine (G) in:
The local run queue of the current logical processor (P), if there’s space.
Otherwise, it goes into the global run queue.
A bound thread (M) working with a processor (P) continuously:
Picks Goroutines from its local run queue.
Or from the global run queue (if its local one is empty).
Or steals work from other P’s queues (if needed).
The OS schedules M (thread) to run on a physical CPU core.
The CPU core executes the code inside the Goroutine (G).


*/
