package code

import "fmt"

type Structure_code struct {
	Value  int
	ValueA int
}

func (v *Structure_code) PassByReference() {
	fmt.Println("PassByReference---->", v)
	v.Value = 23
	v.ValueA = 25
}
func Main_PassByReference() {
	data := new(Structure_code)
	fmt.Println("before", data.Value)
	fmt.Println("after", data.ValueA)
	data.PassByReference()
	fmt.Println("before", data.Value)
	fmt.Println("after", data.ValueA)
}
