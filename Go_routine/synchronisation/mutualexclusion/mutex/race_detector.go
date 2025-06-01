package mutex

import "fmt"

/*
Race Detector and Mutex

* Using `go run -race` to detect race conditions.
* How Mutex helps to fix race conditions.
*/
var counter int

func increment() {
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func Race_detector() {
	go increment()
	go increment()
	fmt.Println("Final counter:", counter)
}

/*
Go provides a built-in race detector to detect race conditions at runtime.
go run -race your_file.go

PS C:\Users\mahindra.gajjala\Desktop\GO\Go_routine> go run -race .\main.go
go: -race requires cgo; enable cgo by setting CGO_ENABLED=1
PS C:\Users\mahindra.gajjala\Desktop\GO\Go_routine>
*/
