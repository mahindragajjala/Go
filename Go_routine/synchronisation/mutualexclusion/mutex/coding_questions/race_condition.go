package codingquestions

//RACE CONDITION
/*
🔸 1. Demonstrate a Race Condition
A race condition occurs when multiple goroutines (threads) access and modify shared data at the same time without proper synchronization.

You need to write a Go program that uses shared data accessed by multiple goroutines, but does not use any synchronization like Mutex.

This will result in unexpected or incorrect output due to race conditions.

🔸 2. Fix the Race Condition using sync.Mutex
Then, modify the same program and use a sync.Mutex to ensure only one goroutine accesses/modifies the shared data at a time.

This will fix the race condition and produce the correct output.
*/
