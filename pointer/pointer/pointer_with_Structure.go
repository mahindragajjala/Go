package pointer

import "fmt"

type Structure struct {
	Value string
	Date  string
}

func Original_data() {
	data := Structure{
		Value: ":(",
		Date:  ":(",
	}
	Pointer_with_Structure2(data)
	fmt.Println(data)
	Pointer_with_Structure(&data)
	fmt.Println(data)

}
func Pointer_with_Structure(x *Structure) {
	x.Value = ":) in 1"
	x.Date = ":)"
}
func Pointer_with_Structure2(s Structure) {
	s.Value = ":( in 2"
	fmt.Println(s.Value)
}
