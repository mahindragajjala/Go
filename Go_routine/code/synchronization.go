package code

/*
Goroutines are lightweight threads used for concurrency in Go. When multiple goroutines access shared resources (like variables, files, maps, etc.), synchronization is essential to prevent:

"Race conditions"

"Data corruption"

"Unpredictable behavior"

Synchronization ensures that only one goroutine can access the critical section (shared resource) at a time, or coordinates when one should wait for another to complete.
*/
