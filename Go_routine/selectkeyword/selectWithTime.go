package selectkeyword

import (
	"fmt"
	"time"
)

var ch1 = make(chan string)
var ch2 = make(chan string)

func Main_Routine() {
	go Goroutine1()
	//go Goroutine2()
	go Goroutine3()
	//go Goroutine4()

	select {
	case ch2 := <-ch1:
		fmt.Println(ch2)
	case ch4 := <-ch2:
		fmt.Println(ch4)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: operation took too long")
	}
}
func Goroutine1() {
	time.Sleep(2 * time.Second)
	ch1 <- "a message from the go routine 1 to 2"
}

func Goroutine2() {
	came := <-ch1
	fmt.Println(came)

}

func Goroutine3() {
	time.Sleep(1 * time.Millisecond)
	ch2 <- "a message from the go routine 3 to 4"

}

func Goroutine4() {
	came := <-ch2
	fmt.Println(came)
}

// PATTERNS
// Timeout waiting for a response
func WaitForResponse() (string, error) {
	ch := make(chan string)

	go func() {
		// Simulate long-running task
		time.Sleep(3 * time.Second)
		ch <- "data"
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-time.After(2 * time.Second):
		return "", fmt.Errorf("timeout")
	}
}

// Timeout inside a loop
// You can use time.After repeatedly or create a time.Timer for more control:
func TIMEOUT_INSIDE_A_LOOP() {
	timer := time.NewTimer(5 * time.Second)
	/*
				timer := time.NewTimer(5 * time.Second)
				What it does:Creates a new timer that will fire exactly once after 5 seconds.
				Internally:

		time.NewTimer returns a *Timer struct. This struct has a channel called C:
		type Timer struct {
		    C <-chan Time  // receives the current time when the timer expires
		}
		So, timer.C will receive the current time value after 5 seconds.
	*/
	for {
		/*
			for { ... }
			Enters an infinite loop.

			We'll stay in this loop until we return from inside it (e.g., when timeout occurs).
		*/
		select {
		case msg := <-ch1:
			fmt.Println("Received:", msg)
		case <-timer.C:
			fmt.Println("Timeout expired, breaking loop")
			return
		}
	}

}

//3. Using time.After to limit total execution time of a goroutine
