package functionliterals

import "fmt"

/*
Anonymous functions are functions without a name.
You can assign them to variables or invoke them immediately.
*/
func Assigned_to_a_variable() {
	greet := func(name string) string {
		return "Hi, " + name
	}

	fmt.Println(greet("Mahindra"))
}
