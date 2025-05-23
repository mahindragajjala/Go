package parentchild

import "fmt"

func Basic_Spawing() {
	fmt.Println("Parent started")

	go func() {
		fmt.Println("Child goroutine running")
	}()

	fmt.Println("Parent finished")
}
