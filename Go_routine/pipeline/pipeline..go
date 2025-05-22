package pipeline

import (
	"fmt"
)

func Pipeline() {
	//input
	nums := []int{2, 3, 4, 7, 1}

	//stage 1 - convert slice into channel
	datachannel := SliceToChannel(nums)

	//stage 2 - square the each number
	finalchannel := Square(datachannel)

	//stage3 - printing the square from the channel
	for n := range finalchannel {
		fmt.Println(n)
	}
}

func SliceToChannel(nums []int) <-chan int { //retuning the read-only channel
	fmt.Println("SliceToChannel asynchoronus goroutine")
	out := make(chan int)
	//asynchoronus goroutine
	go func() {
		for _, value := range nums {
			out <- value
		}
		close(out)
	}()

	return out
}

func Square(in <-chan int) <-chan int {
	fmt.Println("Square asynchoronus goroutine")
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

/*
| Stage             | Goroutine? | Blocking Behavior                 | Sync With                         |
| ----------------- | ---------- | --------------------------------- | --------------------------------- |
| SliceToChannel    | Yes        | `out <- value` blocks             | Sync with Square's input channel  |
| Square            | Yes        | `range in`, `out <- result` block | Sync with upstream and downstream |
| Main loop (Print) | No         | `range finalchannel` blocks       | Sync with Square output           |

SliceToChannel(nums)
go func() {
	for _, value := range nums {
		out <- value
	}
	close(out)
}()
This runs asynchronously because of go func().
BUT! The out <- value line is synchronous with the receiver:
It blocks (waits) until some other goroutine receives the value.
🟢 Verdict:
The goroutine is asynchronous (runs independently),
But each send out <- value is synchronized with the receiver (the Square stage).


Square(datachannel)
go func() {
	for n := range in {
		out <- n * n
	}
	close(out)
}()
Also runs asynchronously due to go func().
It blocks on range in (waits for values from SliceToChannel).
Then sends n * n to out, which also blocks until read by the final loop.
🟢 Verdict:
Again, asynchronous goroutine,
But synchronized with upstream and downstream via channels.
*/
