package multiplegoroutinesscheduling

import (
	"fmt"
	"time"
)

var turn = "ping"

func Ping_Pong() {
	go Ping()
	go Pong()

	// Run for a few seconds before exiting
	time.Sleep(5 * time.Second)
}

func Ping() {
	for {

		if turn == "ping" {
			fmt.Println("Ping...")
			turn = "pong"
			time.Sleep(500 * time.Millisecond)
			/*
				time.Sleep(500 * time.Millisecond)
				After printing and switching turns, the Ping goroutine sleeps for 0.5 seconds (500ms).
				This introduces a delay between alternate prints (for human readability).
				Without this, the output could be too fast and flood the screen.
			*/
		} else {
			/*
				} else {
				This is the else branch of the if turn == "ping" condition.

				It means that it is not Ping’s turn, so this goroutine must wait.
			*/
			time.Sleep(10 * time.Millisecond) // polling delay
		}
	}
}

func Pong() {
	for {
		if turn == "pong" {
			fmt.Println("Pong...")
			turn = "ping"
			time.Sleep(500 * time.Millisecond)
		} else {
			time.Sleep(10 * time.Millisecond) // polling delay
		}
	}
}

// --------------------------------------------------------------------------
var Sharedmemory = "Ping"

func Ping_pong_Prac() {
	go ping()
	go pong()
	time.Sleep(5 * time.Second)
}
func ping() {
	for {
		if Sharedmemory == "Ping" {
			fmt.Println("PING...")
			Sharedmemory = "Pong"
			time.Sleep(500 * time.Millisecond)
		} else {
			time.Sleep(10 * time.Millisecond) // polling delay
		}
	}
}
func pong() {
	for {
		if Sharedmemory == "Pong" {
			fmt.Println("PONG...")
			Sharedmemory = "Ping"
			time.Sleep(500 * time.Millisecond)
		} else {
			time.Sleep(10 * time.Millisecond) // polling delay
		}
	}
}

/*
The Ping() function runs in an infinite loop (for { ... }).

It checks if turn == "ping":
If yes, it prints "Ping...", changes the turn to "pong" and sleeps for 500ms.
If no, it does nothing meaningful—just sleeps for a short time (10ms) and checks again.
This repeated checking without doing actual work is what we call polling.
*/
