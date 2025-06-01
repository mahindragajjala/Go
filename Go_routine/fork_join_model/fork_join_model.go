package forkjoinmodel

import (
	"fmt"
	"sync"
	"time"
)

/*
DEFINTION :
The Fork-Join model in Go (Golang) is a way of breaking a task into smaller subtasks (fork), running them in parallel (usually using goroutines), and then waiting for all of them to finish before continuing (join). It’s a common pattern for parallel or concurrent programming.

Fork: Split the work and start multiple goroutines.
Join: Wait for all the goroutines to complete.
This helps in speeding up work when tasks can be done at the same time.
*/
//----------------------------------------------------------------------------------------------
/*


"Don't communicate by sharing memory, share memory by communicating." - Rob Pike

Go uses the idea of the fork-join model of concurrency behind goroutines.
The fork-join model essentially implies that a child process splits from its parent process to run concurrently with the parent process.
After completing its execution, the child process merges back into the parent process.
The point where it joins back is called the "join point".



Okay, so this works but it's not Perfect. So how do we improve this?

Well, the most tricky part about using goroutines is knowing when they will stop.

It is important to know that goroutines run in the same address space, so access to shared memory must be synchronized.



									+---------------------+
									|    Main Program     |
									+---------------------+
											|
											| (Start)
											v
									+---------------------+
									|   Fork the tasks    |
									+---------------------+
									/         |         \
									v          v          v
							+-----------+ +-----------+ +-----------+
							| Task 1    | | Task 2    | | Task 3    |
							| (Goroutine)| (Goroutine)| (Goroutine)|
							+-----------+ +-----------+ +-----------+
									\         |         /
									\        |        /
									v       v       v
									+---------------------+
									|   Join the tasks    |  ← `wg.Wait()`
									+---------------------+
											|
											v
									+---------------------+
									| Continue Execution  |
									+---------------------+

*/
// Simulate downloading a file
func downloadFile(fileName string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when function exits
	fmt.Printf("Downloading %s...\n", fileName)
	time.Sleep(2 * time.Second) // Simulate time delay
	fmt.Printf("Finished downloading %s\n", fileName)
}

func Fork_join_model() {
	var wg sync.WaitGroup

	files := []string{"file1.jpg", "file2.jpg", "file3.jpg"}

	// Fork: Start a goroutine for each file
	for _, file := range files {
		wg.Add(1) // Increase counter for each goroutine
		go downloadFile(file, &wg)
	}

	// Join: Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All files downloaded. Continue processing...")
}

/*
Fork: go downloadFile(...) starts 3 goroutines in parallel.
Join: wg.Wait() blocks the main function until all 3 are done.
sync.WaitGroup is used to track when all goroutines finish.
*/
