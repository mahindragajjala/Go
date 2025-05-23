package iotakeyword

import "fmt"

const (
	a = iota * 5
	b
	c
	d
)

func Iotakeyword() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
