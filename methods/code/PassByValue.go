package code

import "fmt"

type Structure struct {
	Value  int
	ValueA int
}

func (v Structure) PassByValue() {
	fmt.Println("PassByValue---->", v)
	v.Value = 25
	v.ValueA = 60
}

func Main_PassByValue() {
	data := new(Structure)
	fmt.Println("before", data.Value)
	fmt.Println("after", data.ValueA)
	data.PassByValue()
	fmt.Println("before", data.Value)
	fmt.Println("after", data.ValueA)
}
