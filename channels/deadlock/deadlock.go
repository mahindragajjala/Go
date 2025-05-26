package deadlock

/*
A deadlock occurs when two or more goroutines are waiting on each other to release resources, but none ever do, causing the program to halt indefinitely. In Go, the most common resources that lead to deadlocks are channels and mutexes.

Why Does Deadlock Happen in Go?
In Go, deadlocks typically happen because of incorrect synchronization when using:

Channels (especially unbuffered)
Mutexes
sync.WaitGroup
sync.Cond
Go detects many deadlocks at runtime, especially the main goroutine being blocked with no other runnable goroutines.

Unbuffered Channel with No Receiver
func main() {
	ch := make(chan int)
	ch <- 10 // 🔒 Deadlock: No receiver
}
The main goroutine is blocked forever trying to send on ch.
No goroutine is receiving from it.
*/
