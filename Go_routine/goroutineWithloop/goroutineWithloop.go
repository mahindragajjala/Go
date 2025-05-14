package goroutinewithloop

import (
	"fmt"
	"time"
)

/*
"Goroutines are launched during each loop iteration, but they may start executing after the loop completes."

goroutines inside loops are a common pattern but need special care due to variable capture issues and synchronization

Loop Variable Allocation:
The variable i is declared once before the loop.
In each iteration, its value is updated, but the memory address stays the same.
All goroutines created within the loop share access to the same memory address of i

Goroutine Creation:
In each iteration, a new goroutine is launched via go func() { fmt.Println(i) }().
The anonymous function is a closure, which captures i by reference, not by value.

Memory Management:
Because i is captured by reference, each goroutine reads from the same memory location.
At runtime, goroutines may execute after the loop finishes, by which time i may have reached 5.
As a result, most or all goroutines may print 5, not 0, 1, 2, 3, 4.

Garbage Collection & Escape Analysis:

Since i is referenced by a goroutine (which escapes the current stack frame), Go’s escape analysis promotes i to the heap.
Go runtime's garbage collector (GC) is responsible for cleaning this up once no goroutine refers to it anymore.

Time →
────────────────────────────────────────────────────────

Main Goroutine (loop):
[ i=0 ] ──▶ Launch goroutine #1
[ i=1 ] ──▶ Launch goroutine #2
[ i=2 ] ──▶ Launch goroutine #3
[ i=3 ] ─▶ Launch goroutine #4
[ i=4 ] ─▶ Launch goroutine #5
[ i=5 ] ─▶ Loop ends

Goroutine #1:                           ┌──────────────▶ Print(i) → maybe 5
Goroutine #2:                             ┌────────────▶ Print(i) → maybe 5
Goroutine #3:                               ┌──────────▶ Print(i) → maybe 5
Goroutine #4:                                 ┌────────▶ Print(i) → maybe 5
Goroutine #5:                                   ┌──────▶ Print(i) → maybe 5
*/
func Incorrect_Capture() {
	fmt.Println("Incorrect_Capture---->")
	for i := 0; i < 5; i++ {
		fmt.Printf("OUT LOOP\n")
		go func() {
			fmt.Println("GO ROUTINE")
			fmt.Println(i)
		}()
		//time.Sleep(2 * time.Millisecond)
	}
	time.Sleep(time.Second) // Give goroutines time to run

}

/*
1. Loop Initialization
Variable i is declared and initialized to 0.
Memory for i is allocated ("likely on the heap, due to goroutine closure capturing it").

2. Each Loop Iteration
For each i from 0 to 4:
a. Loop increments i
i = 0 -> 1 -> 2 -> 3 -> 4

b. Goroutine is launched

	go func() {
	    fmt.Println(i)
	}()

3. Goroutine Scheduling
Goroutines are scheduled independently of the loop.
"The for loop might complete all 5 iterations before any goroutine executes."
All goroutines read from the same memory location of i, which is now 5 after the loop.

4. Inside Each Goroutine
Each goroutine executes:

fmt.Println(i)
Since i is captured by reference, and the loop has finished, each goroutine prints 5.
The actual value depends on when the goroutine executes relative to the loop's progression.
*/

func Using_parameter() {
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
}
