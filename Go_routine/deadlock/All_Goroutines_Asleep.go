package deadlock

/*
All Goroutines Asleep — Nothing Left to Do

func main() {
	ch := make(chan int)
	<-ch // 🔒 Deadlock: Waiting forever
}
*/
/*
Go detects this situation at runtime. It means that:

All goroutines, including the main goroutine, are waiting (blocked).
There’s nothing left to do, because no goroutine can proceed.
So the Go runtime panics with this message.

1. Channel Read/Write Mismatch
You tried to read from or write to a channel, but there's no corresponding writer or reader, so the goroutine blocks forever.
func main() {
    ch := make(chan int)
    ch <- 10 // No goroutine to read from channel => blocks forever
}

2. Infinite Waiting on an Unbuffered Channel
Reading or writing to an unbuffered channel without a partner causes deadlock.
func main() {
    ch := make(chan int)
    <-ch // nothing ever sent => deadlock
}

3. No Goroutines Running
Sometimes the program creates goroutines, but the main goroutine exits before they run, or nothing meaningful is done.
func main() {
    go func() {
        fmt.Println("Hello")
    }()
    // Main exits immediately => goroutine may not run
}
	solution
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Hello")
    }()
    wg.Wait()
}

4. Loop Waiting on a Channel with No Input
You use range over a channel, but nobody closes the channel.
func main() {
    ch := make(chan int)
    for val := range ch {
        fmt.Println(val)
    }
}

solution
func main() {
    ch := make(chan int)
    go func() {
        ch <- 1
        ch <- 2
        close(ch)
    }()
    for val := range ch {
        fmt.Println(val)
    }
}

5. Missing select Case for Channel
You use select {} with no case, which causes the program to sleep forever.
func main() {
    select {} // Blocks forever, nothing happens
}


How to Avoid This Error

Ensure every write to a channel has a corresponding read, and vice versa.
Use sync.WaitGroup to ensure goroutines finish before main() exits.
Always close channels when done writing (especially when using range).
Be careful with infinite loops or select {} with no cases.
*/
