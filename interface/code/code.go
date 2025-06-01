package code

import "fmt"

type Cardata struct {
	Name  string
	Price int
}

func (C Cardata) Drive() {
	fmt.Println("we are driving...")
}
func (C Cardata) Race() {
	fmt.Println("we are ready to race")
}
func GO() {
	data := new(Cardata)
	data.Drive()
	data.Race()
}
