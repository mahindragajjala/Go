package multiplegoroutinesscheduling

import (
	"fmt"
	"sync"
)

/*
Create a goroutine that counts up to 10 and a main routine that waits for it to finish.

Create a goroutine that counts from 1 to 10.
Ensure the main routine (main function) waits for the goroutine to complete before exiting.
*/
var count int = 0

func Counting() {
	var wgc sync.WaitGroup
	wgc.Add(1)
	go countFunc(&wgc)
	wgc.Wait()
	fmt.Println(count)
}
func countFunc(wgc *sync.WaitGroup) {
	fmt.Println("its coming..")
	defer wgc.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		count++
	}
}
