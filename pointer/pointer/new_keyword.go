package pointer

import "fmt"

//new allocates memory and returns a pointer.
func New_keyword() {
	newdata := new(int)
	fmt.Println(newdata)

	slicedata := new([]int)
	mapdata := new(map[int]string)
	fmt.Println(slicedata, mapdata)
}
