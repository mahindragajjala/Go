package questions

import (
	"fmt"
)

func ReverseString(input string) string {
	// Convert the string to a rune slice to handle Unicode characters
	runes := []rune(input)
	fmt.Println("runes--->", runes)

	n := len(runes)
	fmt.Println("length--->", n)

	// Use a for loop to swap characters
	for i := 0; i < n/2; i++ {
		// Swap the characters
		fmt.Println(runes[i], runes[n-1-i])
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	// Convert the rune slice back to a string
	return string(runes)
}

func Reverse_String() {
	original := "Hello, GoLang!"
	reversed := ReverseString(original)
	fmt.Println("Original:", original)
	fmt.Println("Reversed:", reversed)
}
