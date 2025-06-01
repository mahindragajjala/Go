package code

import "fmt"

type data struct {
	Data interface{}
}

func Data() {
	d := data{
		Data: "string",
	}
	fmt.Println(d)
}
