package code

import "fmt"

type Car struct {
	Name  string
	Price int
}

func (C Car) Drive() {
	fmt.Println("we are driving...")
}
func (C Car) Race() {
	fmt.Println("we are ready to race")
}
func GO() {
	data := new(Car)
	data.Drive()
	data.Race()
}
