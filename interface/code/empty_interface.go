package code

import "fmt"

func Empty_interface(a interface{}) {
	fmt.Println(a)
}
func Main_Empty_interface() {
	Empty_interface(100)
	Empty_interface("aaa")
	Empty_interface(100.00)

}
