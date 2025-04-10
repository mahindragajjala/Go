package main

import (
	"fmt"
	"time"
)

//Start 5 Goroutines, each printing its index (from 1 to 5). Ensure the main function waits for all Goroutines to finish.

func main() {
	for i := 0; i < 5; i++ {
		go Index(i)
		time.Sleep(2 * time.Second)
	}
}
func Index(index int) {
	fmt.Println("go routine: %D ", index)
	var data int
	data = index + 1
	fmt.Println("index is : %V", data)
}
