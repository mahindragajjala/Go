package looponchannels

/*
	Range over a channel
	Infinite Loop with Manual Break
	Loop with Select Statement
	Loop Until Timeout
	Loop with Done Channel (Cancellation Pattern)
	Looping Over Buffered Channels
	Loop with Goroutine Pool Pattern
	Loop with Fan-In Pattern
	Loop with Fan-Out Pattern
*/
/*
| Pattern                  | Description                        | Needs Close? |
| ------------------------ | ---------------------------------- | ------------ |
| `range ch`               | Simplest loop, auto-exits on close | ✅            |
| `for { <-ch }`           | Manual check with `ok`             | ✅            |
| `for { select { ... } }` | Multiplex channels                 | ❌            |
| `for with timeout`       | Non-blocking with timeout          | ❌            |
| `for with done`          | Graceful cancellation              | ✅            |
| Buffered channels        | Queue-like usage                   | ✅            |
| Goroutine pool           | Parallel processing                | ✅            |
| Fan-In                   | Merge many to one                  | ❌            |
| Fan-Out                  | Split one to many                  | ✅            |

*/
