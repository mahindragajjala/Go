package code

import "fmt"

type Structure struct {
	value     int
	valuedata int
}
type Nested_Struct struct {
	v Structure
}

func Nested_structure() {
	Printing := Nested_Struct{v: Structure{value: 0, valuedata: 0}}
	fmt.Println(Printing)
}
