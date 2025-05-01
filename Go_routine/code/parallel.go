package code

import (
	"runtime"
	"sync"
)

func Parallel() {
	runtime.GOMAXPROCS(2) // Allow parallelism on 2 cores

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		downloadFile("file1")
	}()

	go func() {
		defer wg.Done()
		downloadFile("file2")
	}()

	wg.Wait()
}
