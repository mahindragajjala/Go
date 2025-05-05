package goroutinewithloop

import "fmt"

/*
goroutines inside loops are a common pattern but need special care due to variable capture issues and synchronization
*/
func Incorrect_Capture() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func Capture_by_Value() {
	for i := 0; i < 5; i++ {
		i := i // create a new local copy of i
		go func() {
			fmt.Println(i)
		}()
	}
}
func Using_parameter() {
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
}
