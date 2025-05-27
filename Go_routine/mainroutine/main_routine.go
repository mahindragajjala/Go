package mainroutine

import "fmt"

/*
When you run a Go program, it starts with the main() function. This main() function runs in a special goroutine called the main goroutine.

What Happens When the Main Goroutine Exits?

When the main() function finishes running, the entire program ends, even if there are other goroutines running in the background.

This is very important:
The Go runtime does not wait for other goroutines to finish automatically.
Once the main goroutine exits, all other goroutines are killed, even if they were doing work.
*/

/*
How to Wait for Other Goroutines?
To keep the program running until other goroutines are done, use time.Sleep, sync.WaitGroup, or channels.
*/

func sendValues(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func MEMORY_MANAGEMENT() {
	ch := make(chan int)
	go sendValues(ch)
	for val := range ch {
		fmt.Println("Received:", val)
	}
}
