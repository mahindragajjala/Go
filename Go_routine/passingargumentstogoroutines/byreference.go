package passingargumentstogoroutines

import "fmt"

/*
You pass the memory address, so the function can modify the original variable.
*/
func printingValue(n *int) {
	fmt.Println("Value:", *n)
}

func ByReference() {
	n := 10
	go printingValue(&n)
}

/*
The function can access and modify the original n.
You must be careful with "race conditions" (when two goroutines access the same variable).
*/
