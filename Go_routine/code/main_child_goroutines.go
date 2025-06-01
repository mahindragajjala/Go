package code

/*
Main goroutine vs child goroutines

Main goroutine:
The "main() function" runs in the main goroutine.
It is the "entry point" of the Go program.
When the main() function finishes execution, the "entire program terminates", even if other (child) "goroutines are still running".

This means: "child goroutines must complete before main() exits", or they will be "forcefully terminated".



Child Goroutines:

Any goroutine started using the go keyword (e.g., go someFunc()) is a child goroutine.
They run "concurrently with the main goroutine".

You need synchronization mechanisms like:
sync.WaitGroup
channels
time.Sleep (not recommended for real use) to ensure they complete before the program exits.

func child(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Hello from child goroutine")
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go child(&wg)
    fmt.Println("Hello from main goroutine")
    wg.Wait() // Wait for child to finish
}

*/
