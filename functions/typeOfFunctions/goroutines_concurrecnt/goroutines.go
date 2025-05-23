package goroutinesconcurrecnt

import "fmt"

func Goroutines() {
	/*
		Goroutine-Invoked Functions (Concurrency)
		Functions run asynchronously using goroutines.
	*/
	go func() {
		fmt.Println("Running in a goroutine")
	}()
}

//read in the folder:)
