package multiplegoroutinesscheduling

import (
	"fmt"
	"sync"
)

/*
Write a program that starts 5 goroutines, each printing a number from 1 to 5.
*/
var wga sync.WaitGroup

func Nnumbers_Goroutine() {

	for i := 1; i <= 5; i++ {
		wga.Add(1)
		go Printing(i, &wga)
	}
	wga.Wait()
}
func Printing(n int, wga *sync.WaitGroup) {
	defer wga.Done()
	fmt.Println("GO ROUTINE :", n)
	fmt.Println(n)
}

/*
// Function that accepts a number and WaitGroup pointer
func printNumber(n int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when finished
	fmt.Println("Number:", n)
}

func main() {
	var wg sync.WaitGroup

	// Start 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the counter before starting a goroutine

		// Pass the loop variable and waitgroup to the goroutine
		go printNumber(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All goroutines finished.")
}
*/
