package content

import (
	"fmt"
	"strings"
)

/*
In Go, function types themselves cannot implement interfaces directly, but you can wrap functions inside types that implement interface methods. This lets you pass behavior (functions) flexibly, like using function pointers in C or delegates in C#.
*/
// Define an interface
type Processor interface {
	Process(input string) string
}

// Function type
type FuncProcessor func(string) string

// Implement the interface using a function type with a method
func (f FuncProcessor) Process(input string) string {
	return f(input)
}

func Using_Function_Types_as_Interfaces() {
	// Create functions
	toUpper := FuncProcessor(func(s string) string {
		return "Upper: " + strings.ToUpper(s)
	})

	reverse := FuncProcessor(func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return "Reverse: " + string(runes)
	})

	// Use them as Processor interface
	processInput(toUpper, "golang")
	processInput(reverse, "golang")
}

func processInput(p Processor, input string) {
	fmt.Println(p.Process(input))
}
