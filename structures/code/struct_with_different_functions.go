package code

import "fmt"

type Counter struct {
	Value int
}

func (c Counter) Increment() Counter {
	c.Value++
	return c
}

// Method 2
func (c *Counter) Reset() {
	c.Value = 0
}
func Struct_with_different_functions() {

	// Method 1
	c := Counter{Value: 10}
	fmt.Println("Original:", c)

	c = c.Increment() // returns a new copy
	fmt.Println("After Increment:", c)

	c.Reset() // modifies the original using pointer
	fmt.Println("After Reset:", c)
}
