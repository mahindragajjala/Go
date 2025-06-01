package questions

import "fmt"

type Rectangle struct {
	Length  int
	Breadth int
}

func (a *Rectangle) AreaOfRectangle() {
	area_of_rectangle := a.Length * a.Breadth
	fmt.Println(area_of_rectangle)
}
func Main_test() {
	input1 := new(Rectangle)
	input1.Length = 5
	input1.Breadth = 6
	input2 := &Rectangle{Length: 20, Breadth: 5}
	input2.AreaOfRectangle()
}
