package pointer

import "fmt"

//A pointer can also store the address of another pointer.
func Pointer_to_Pointer() {
	var a int = 10
	var p *int = &a
	var pp **int = &p

	fmt.Println(**pp) // Output: 10

}
func Practice() {
	var variable int = 10
	var frist *int = &variable
	var second **int = &frist
	fmt.Println(second)
}
