package routing

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Parallelism requires multiple CPU cores.
runtime.GOMAXPROCS(n)
This tells Go to use n OS threads for parallel execution.
*/
func heavyTask(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1e7; i++ {
		// simulate CPU work
	}
	fmt.Println(name, "done")
}

func Parallelism() {
	runtime.GOMAXPROCS(4) // use 4 cores

	var wg sync.WaitGroup
	wg.Add(2)

	go heavyTask("Task A", &wg)
	go heavyTask("Task B", &wg)

	wg.Wait()
	fmt.Println("All tasks finished")
}
