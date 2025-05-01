package code

import (
	"fmt"
	"time"
)

// Using Goroutines in Go (Concurrent)
func Concurrency() {
	go downloadFile("file1")
	go downloadFile("file2")
	logProgress()
}

func downloadFile(name string) {
	// simulate downloading
	for i := 0; i < 5; i++ {
		fmt.Println(name, "downloading part", i)
		time.Sleep(time.Second)
	}
}

func logProgress() {
	for i := 0; i < 5; i++ {
		fmt.Println("Logging progress...")
		time.Sleep(500 * time.Millisecond)
	}
}
