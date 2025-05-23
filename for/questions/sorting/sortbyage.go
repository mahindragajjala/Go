package questions

import "fmt"

/*
Sort Struct by Field
Given a slice of structs like []Person{name string, age int}, sort by age using Bubble Sort.
*/
type Person struct {
	Name string
	age  int
}

func Sort_struct_by_field() {
	data := []Person{{Name: "a", age: 52}, {Name: "b", age: 21}, {Name: "c", age: 72}}
	fmt.Println(data)
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j].age < data[j+1].age {
				data[j].age, data[j+1].age = data[j+1].age, data[j].age
			}
		}
	}
	fmt.Println(data)
}
