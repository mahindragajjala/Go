package passingargumentstogoroutines

import "fmt"

func printValue(n int) {
	fmt.Println("Value:", n)
}

//You pass a copy of the data
func PassByValue() {
	n := 10
	go printValue(n)
}

/*
It runs asynchronously (in the background), meaning the rest of your code can continue running while the goroutine does its work.

The function gets a copy of n.

Changes made to n inside printValue do not affect the original n in main.
*/
