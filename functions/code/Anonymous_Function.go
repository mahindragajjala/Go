package code

import "fmt"

/*
An anonymous function in Go is a function without a name.
It is usually defined inline and can be:
Immediately invoked (called right away), or
Assigned to a variable for later use.
Anonymous functions are useful for short operations, callbacks, closures, or function literals when a full function declaration is unnecessary.
*/
func Anonymous_Functions() {
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Go Developer")
}
