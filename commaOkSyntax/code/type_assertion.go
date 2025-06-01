package code

import "fmt"

func Type_assertion() {
	var x interface{} = "Hello"
	str, ok := x.(string)
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("Not a string!")
	}
}
