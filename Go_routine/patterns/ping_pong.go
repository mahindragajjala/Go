package patterns

import (
	"fmt"
	"sync"
	"time"
)

/*
Ping-pong using goroutine communication in Go is a classic concurrency exercise.

It demonstrates how two goroutines can communicate and coordinate using channels — Go’s concurrency primitive.
*/
/*
Two goroutines (ping and pong) send messages back and forth — like a ball being hit between two players.
*/
/*
Iteration 0:
Goroutine Ping(0):        Goroutine Pong():
      |                        |
      | -- send(0) --------->  |
      |                        |
      |                        | -- receive(0)
      |                        | -- print("Pong received: 0")

Iteration 1:
Goroutine Ping(1):        Goroutine Pong():
      |                        |
      | -- send(1) --------->  |
      |                        |
      |                        | -- receive(1)
      |                        | -- print("Pong received: 1")

Iteration 2:
Goroutine Ping(2):        Goroutine Pong():
      |                        |
      | -- send(2) --------->  |
      |                        |
      |                        | -- receive(2)
      |                        | -- print("Pong received: 2")

...

Iteration 9:
Goroutine Ping(9):        Goroutine Pong():
      |                        |
      | -- send(9) --------->  |
      |                        |
      |                        | -- receive(9)
      |                        | -- print("Pong received: 9")
*/

type CHANNEL struct {
	wg   *sync.WaitGroup
	data chan int
}

func (ch *CHANNEL) Ping(n int) {
	defer ch.wg.Done()
	ch.data <- n
}

func (ch *CHANNEL) Pong() {
	defer ch.wg.Done()
	val := <-ch.data
	fmt.Println("Pong received:", val)
}

func Ping_Pong() {
	ch := CHANNEL{
		wg:   &sync.WaitGroup{},
		data: make(chan int), // initialize the channel
	}

	for i := 0; i < 10; i++ {
		ch.wg.Add(2)
		go ch.Ping(i)
		go ch.Pong()
		ch.wg.Wait()
	}
}

/*
Goroutine A:         Goroutine B:
    |                     |
    | -- "ping" ------->  |
    |                     |
    |  <--- "pong" ------ |
    |                     |
    | -- "ping" ------->  |
    |                     |
    |  <--- "pong" ------ |

	+------------+       pingCh        +------------+       pongCh       +------------+
|            | <------------------ |            | <----------------- |            |
|   Ping     |                    ^|   Pong     |                   ^|   Main      |
| Goroutine  | ------------------> | Goroutine  | -----------------> | Goroutine  |
|            |       pongCh        |            |       pingCh       |            |
+------------+                    v+------------+                    v+------------+
        ^                                                                 |
        |                                                                 |
        +----------------- done chan (completion signal) <---------------+


*/

// Ping function sends to pongCh and receives from pingCh
func ping(pingCh, pongCh chan string, done chan bool, count int) {
	for i := 0; i < count; i++ {
		msg := <-pingCh // Wait for "ping" to come
		fmt.Println("Ping received:", msg)

		time.Sleep(500 * time.Millisecond)    // Simulate work
		pongCh <- fmt.Sprintf("pong %d", i+1) // Send "pong" back
	}
	done <- true // Notify main that we're done
}

// Pong function sends to pingCh and receives from pongCh
func pong(pingCh, pongCh chan string, count int) {
	for i := 0; i < count; i++ {
		msg := <-pongCh // Wait for "pong" to come
		fmt.Println("Pong received:", msg)

		time.Sleep(500 * time.Millisecond)    // Simulate work
		pingCh <- fmt.Sprintf("ping %d", i+1) // Send "ping" back
	}
}

func PING_PONG() {
	pingCh := make(chan string)
	pongCh := make(chan string)
	done := make(chan bool)

	count := 5 // Number of ping-pong exchanges

	// Start goroutines
	go ping(pingCh, pongCh, done, count)
	go pong(pingCh, pongCh, count)

	// Start the first message
	pingCh <- "start"

	// Wait for the process to complete
	<-done

	fmt.Println("Ping-pong finished.")
}
