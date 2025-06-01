package buffered

import "fmt"

/*
Non-blocking send on buffered channel:
Write a program where the main goroutine sends 3 values to a buffered channel of size 3 (non-blocking).
Then try to send a fourth value and explain what happens.
*/
/*
PS C:\Users\mahindra.gajjala\Desktop\GO\Go_routine> go run main.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
routine/communication/buffered.Blocking()
        C:/Users/mahindra.gajjala/Desktop/GO/Go_routine/communication/buffered/blocking.go:15 +0x69
main.main()
        C:/Users/mahindra.gajjala/Desktop/GO/Go_routine/main.go:15 +0xf
exit status 2
PS C:\Users\mahindra.gajjala\Desktop\GO\Go_routine>
*/
func Blocking() {
	Buffered := make(chan int, 3)
	Buffered <- 1
	Buffered <- 2
	Buffered <- 3
	Buffered <- 4
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
}

/*
occurs when no goroutine can make progress, typically because they are all blocked waiting for an event that will never happen.

When you run a Go program:

""""""""""The function main() runs inside the main goroutine (implicitly created by the Go runtime).
So even without writing go, you are always running inside at least one goroutine — the main one.""""""""""

Start main goroutine:
    ↓
Create Buffered channel (capacity 3)
    ↓
Send 1 → channel (Buffered = [1])
Send 2 → channel (Buffered = [1, 2])
Send 3 → channel (Buffered = [1, 2, 3])
Send 4 → ❌ BLOCKS! (channel full)
No other goroutine exists to read from the channel → so the main goroutine is permanently stuck → deadlock.
*/
/*
The 4th send blocks, waiting for space to be available.
But there's no receiver running concurrently — your code runs sequentially.
*/

//Read before writing too much
func Blocking_solution() {
	Buffered := make(chan int, 3)
	Buffered <- 1
	Buffered <- 2
	Buffered <- 3
	fmt.Println(<-Buffered) // Read one to make space
	Buffered <- 4           // Now this will work
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
}

// Use a goroutine for sending or receiving
func Blockinguse_goroutine() {
	Buffered := make(chan int, 3)
	go func() {
		Buffered <- 1
		Buffered <- 2
		Buffered <- 3
		Buffered <- 4 // Won’t block because receiver is active
	}()

	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
}
