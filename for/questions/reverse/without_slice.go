package questions

import "fmt"

/*
Reverse a String Without Using Slices
How can you reverse a string in Go using a for loop and string concatenation?
*/
func Reverse_Without_slice() {
	Stringdata := "hello"
	fmt.Println(len(Stringdata))

	reversed := ""

	for i := len(Stringdata) - 1; i >= 0; i-- {
		reversed += string(Stringdata[i])
	}
}
