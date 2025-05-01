package code

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func CreateGoRoutines() {
	go printNumbers()       // Creating and running the goroutine
	time.Sleep(time.Second) // Wait to allow goroutine to complete
	fmt.Println("Main function exiting...")
}
