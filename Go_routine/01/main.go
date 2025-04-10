package main

import (
	"fmt"
	"time"
)

//Write a program that starts a Goroutine to print
//"Hello from Goroutine" while the main function prints "Hello from main".

func main() {
	fmt.Println("Hello from before_main")
	go Go_routing()
	time.Sleep(2 * time.Second)
	fmt.Println("Hello from after_main")
}
func Go_routing() {
	fmt.Println("Hello from Goroutine")
}

//Start 5 Goroutines, each printing its index (from 1 to 5). Ensure the main function waits for all Goroutines to finish.
