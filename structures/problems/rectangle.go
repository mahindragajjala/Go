package problems

import "fmt"

/*
Create a struct Rectangle with Width and Height.
Write a method Area() that calculates and returns the area.
*/
type Rectangle struct {
	Width  int
	Height int
}

func Parameters() {
	area := Rectangle{Width: 2, Height: 5}
	a := AreaOfRectangle(&area)
	fmt.Println(a)
}
func AreaOfRectangle(area *Rectangle) int {
	areadata := area.Height * area.Width
	return areadata
}
