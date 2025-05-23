package code

import "fmt"

type MethodDate struct {
	Value1 int
	Value2 int
}

func (m MethodDate) MethodFunc() {
	fmt.Println(m)
}
