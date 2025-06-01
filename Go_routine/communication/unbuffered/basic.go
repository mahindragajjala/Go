package unbuffered

import "fmt"

/*
Write a program that creates an unbuffered channel.
Start a goroutine that sends an integer (say 42) to the channel.
In the main goroutine, receive the value from the channel and print it.
*/
func Unbuffered() {
	Unbuffered := make(chan int)
	go func() {
		Unbuffered <- 42
	}()
	fmt.Println(<-Unbuffered)
	close(Unbuffered)
}
