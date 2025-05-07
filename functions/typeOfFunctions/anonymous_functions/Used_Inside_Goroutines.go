package anonymousfunctions

import "fmt"

func Goroutines() {
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from goroutine")

}
