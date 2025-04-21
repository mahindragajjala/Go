package code

/*
In Go, type assertion is a way to extract the concrete value from an interface type.

When you have a variable of interface{} type, you don’t know what its actual (underlying) type is.
A type assertion helps you assert that the value inside the interface is of a specific type.

value := myInterface.(ConcreteType)
myInterface — variable of type interface{}.
ConcreteType — the real type you expect.
value — the extracted value of ConcreteType.
*/
import "fmt"

func Type_assertion() {
	var i interface{} = "Hello, Go!"

	str := i.(string) // asserting that i holds a string
	fmt.Println(str)

	str, ok := i.(string)
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("Type assertion failed!")
	}
}
