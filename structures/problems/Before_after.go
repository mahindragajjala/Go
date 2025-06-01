package problems

import "fmt"

/*
Write a function that accepts a struct Person as a parameter and modifies one of its fields.
Print before and after the modification.
*/
type Student struct {
	Name  string
	Class int
}

func Before_modification() {
	data := Student{Name: "apple", Class: 26}
	fmt.Println("BEFORE MODIFICATION:", data)
	After_modification(&data)
	fmt.Println("AFTER MODIFICATION:", data)

}
func After_modification(data *Student) {
	data.Name = "BANANA"
}
