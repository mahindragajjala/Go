package content

import "fmt"

//shared data in the slice

func Shared_Data() {
	Shared_Data := []int{1, 2, 3} //012
	a := Shared_Data[0:2]
	b := Shared_Data[1:]

	fmt.Printf("BEFORE a:%v \n b :%v\n", a, b)

	Shared_Data[1] = 100

	fmt.Printf("AFTER a:%v \n b :%v", a, b)
}
