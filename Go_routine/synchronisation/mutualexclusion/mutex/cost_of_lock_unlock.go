package mutex

/*
⚙️ What Happens When You Lock/Unlock a sync.Mutex?

Locking (mu.Lock())
Tries to acquire the lock.

If lock is already held, the goroutine is added to a queue (FIFO) and sleeps (context switch).
Wakes up and gets the lock once it’s free.
Unlocking (mu.Unlock())
Releases the lock.
If there are waiting goroutines, it wakes one of them.

⏱️ Cost Factors
Fast Path (No Contention): Only a few CPU instructions. Extremely fast (~10-20 ns).

Slow Path (With Contention):

Involves syscalls, context switches.
Cost goes up to microseconds or more.
Blocking one goroutine suspends it and wakes another, which involves scheduler overhead.
Mutexes are fast when uncontended, but contention causes context switching, which is expensive.
*/
