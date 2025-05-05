package threadvsgoroutine

/*
| Aspect                 | Thread                                               | Goroutine                                                               |
| ---------------------- | -------------------------------------------------------- | --------------------------------------------------------------------------- |
| Definition         | A thread is a unit of execution managed by the OS.       | A goroutine is a lightweight, user-space thread managed by the Go runtime.  |
| Creation Cost      | Relatively heavy — involves OS-level overhead.           | Very lightweight — takes a few KBs of stack memory.                         |
| Managed By         | Operating System.                                        | Go runtime scheduler.                                                       |
| Memory Consumption | Each thread typically takes 1–2 MB of stack space.       | Each goroutine starts with \~2 KB of stack and grows as needed.             |
| Concurrency Model  | Preemptive multitasking by OS.                           | Cooperative scheduling by Go runtime.                                       |
| Scalability        | Limited — too many threads can crash or slow the system. | Highly scalable — thousands or millions of goroutines can run concurrently. |
| Communication      | Uses shared memory with locks (mutexes).                 | Encourages message passing via channels (CSP model).                        |
| Blocking           | Blocking a thread blocks the OS thread.                  | Blocking a goroutine does not block the OS thread — others can run.     |
| Context Switch     | Managed by OS — heavier, more costly.                    | Managed by Go runtime — lighter and faster.                                 |
| Use Case           | General purpose in all languages (Java, C++, etc.).      | Specific to Go — used for highly concurrent applications.                   |

*/
