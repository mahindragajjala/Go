package operations

import "fmt"

func Start_goroutine() {
	go someFunction()
}
func someFunction() { fmt.Println("go routine") }
