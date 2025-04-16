package code

import "fmt"

func Goroutines() {
	go func() {
		fmt.Println("go routines...")
	}()
}
