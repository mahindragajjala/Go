package goroutinewithloop

import "fmt"

/*
Too Many Goroutines (Memory Leak)
Problem: Spawning goroutines in a large loop without control may cause memory issues or crash.
*/
func Memory_Leak() {
	for i := 0; i < 1e9; i++ {
		go doSomething()
	}
}

func doSomething() {
	fmt.Print("memory leak...")
}

/*
Use a worker pool or limit goroutines using buffered channels/semaphores.
*/
