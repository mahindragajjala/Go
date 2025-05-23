package code

import "fmt"

func Describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is an int:", v)
	case string:
		fmt.Println("This is a string:", v)
	case bool:
		fmt.Println("This is a bool:", v)
	default:
		fmt.Println("Unknown type")
	}
}
