package goscheduler

import (
	"fmt"
	"runtime"
	"time"
)

/*
Definition : In Go (Golang), GOMAXPROCS is a setting that determines the maximum number of operating system threads that can execute Go code simultaneously. It controls how many CPU cores your Go program can use at once.

Imagine you're in an office:

You have 8 employees (CPU cores).
You get many small tasks (goroutines) to complete.
But you tell your office: “Only 2 employees can work at the same time.” That’s GOMAXPROCS=2.
Even though you have 8 people, only 2 are allowed to work at a time. This limits how fast you can complete tasks — that's exactly what GOMAXPROCS does in Go: it limits how many CPU threads (cores) your program can use simultaneously.

Using Environment Variable
Set it when running your program:

GOMAXPROCS=4 go run main.go
*/
func GOMAXPROCS() {
	previous := runtime.GOMAXPROCS(2) // Set to use 2 OS threads
	fmt.Println("Previous GOMAXPROCS value:", previous)
}
func Test_GOMAXPROCS() {
	runtime.GOMAXPROCS(1) // Force Go to use only 1 CPU core

	go func() {
		for {
			fmt.Print("A")
		}
	}()

	go func() {
		for {
			fmt.Print("B")
		}
	}()

	time.Sleep(time.Second)
}
