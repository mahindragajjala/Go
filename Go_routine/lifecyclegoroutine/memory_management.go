package lifecyclegoroutine

import (
	"fmt"
	"time"
)

/*
[Spawned]run

	|
	v

[Running]execution

	|
	v

[Terminated]

Go creates a small stack (around 2KB) for the goroutine.
It dynamically grows/shrinks as needed.
The Go runtime scheduler manages when and how goroutines run.
*/
func Memory_management_test() {
	go GO_Routine1()
	go GO_Routine2()
	time.Sleep(1 * time.Second) // Give time for the goroutine to finish
}
func GO_Routine1() {
	var routine1var int = 1
	fmt.Println("go routine 1", routine1var)
}

func GO_Routine2() {
	var routine2var int = 2
	fmt.Println("go routine 2", routine2var)
}
