package emptystructgo

import (
	"fmt"
	"time"
)

/*
 What is struct{} in Go?

struct{} is a zero-size struct in Go.
It takes up zero bytes of memory.
It is often used as a placeholder, a signal, or a marker.
Unlike other structs, it has no fields, hence it's called empty.
*/
/*
These look similar but mean slightly different things:
struct{} refers to the type — it defines a struct with no fields.
struct{}{} is a value — an instance of the empty struct.
*/

//As a Set Element (for unique keys)
/*
Since it takes no memory, map[string]struct{} is used instead of map[string]bool when you just care about existence.
*/
func Unique_keys() {
	visited := map[string]struct{}{}

	visited["node1"] = struct{}{}
	visited["node2"] = struct{}{}

	if _, ok := visited["node1"]; ok {
		fmt.Println("node1 is visited")
	}
	/*
	   Why? Because struct{} uses zero memory, whereas bool takes one byte.
	*/
}

// As a Signal (Especially in Channels and Goroutines)
// As a signal to stop, start, or notify.
func Signal() {
	done := make(chan struct{})

	go func() {
		// Do some work
		done <- struct{}{} // Signal done
	}()

	<-done // Wait for signal
}

/*
Why use struct{} here instead of "bool or int"?
You don't need to send data — just a signal.
struct{} takes zero memory and allocates nothing.
It's semantically more clear: “Just signaling”.
*/

//How struct{} is Used in "Goroutines" and "Channels"

//One-Time Signal Between Goroutines - You want one goroutine to send a signal once, and another goroutine to wait until that signal is received before continuing.
/*
Analogy (Real Life)
Imagine:

You are cooking rice.
Your friend is waiting for the rice to be cooked before setting the table.
Once you say, "Rice is ready!" (a one-time signal), your friend begins their work.
You don’t need to pass how much rice or any data — just the signal that it's done.
*/

func One_time_signal_between_goroutines() {
	stop := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		stop <- struct{}{} // send stop signal
	}()

	<-stop // Wait until the signal is received
	fmt.Println("Stopped!")

}

//Broadcast Shutdown Signal (Close Channel)
/*
Analogy(real life)
LIBRARY WITH EARBUDS:
You're in a library with 5 people wearing wireless earbuds. Suddenly, the librarian turns off the main transmitter. All 5 earbuds stop receiving — at the same time — without needing individual messages.

That’s like closing a channel: everyone listening knows instantly
Let’s write code that:
Starts multiple goroutines doing work.
Main function sends a broadcast stop signal by closing a channel.
All goroutines stop at the same time.
*/
func Broadcast_shutdown_signal() {
	done := make(chan struct{}) // broadcast channel

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, done)
	}

	time.Sleep(2 * time.Second) // let workers run
	fmt.Println("Main: sending shutdown signal")
	close(done) // 🔴 broadcast to all workers

	time.Sleep(1 * time.Second) // wait for workers to finish
	fmt.Println("Main: all workers stopped")
}
func worker(id int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Printf("Worker %d: stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
| Type       | Use Case                         | Memory | Suitable? |
| ---------- | -------------------------------- | ------ | --------- |
| `bool`     | True/False signals               | 1 byte | Yes       |
| `int`      | Counting, values                 | ≥4B    | Overkill  |
| `string`   | Descriptive messages             | varies | Overkill  |
| `struct{}` | Only need to **signal** (notify) | 0 byte | ✅ Best    |

*/
