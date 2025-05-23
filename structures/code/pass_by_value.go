package code

import "fmt"

//Pass by value

/*
When you pass a variable to a function, a copy of the variable is made —
the function works on the copy, not the original.
*/
func changeValue(x int) {
	x = 20
}

/*
Even though changeValue() sets x to 20, the original variable a in Pass_by_value()
remains 10 because x is just a copy.
*/
func Pass_by_value() {
	a := 10
	changeValue(a)
	fmt.Println("a =", a) // Output: a = 10
}

// slices, maps, and pointers?
func changeValue_pointer(x *int) {
	*x = 20
}
func For_Pointer_Pass_by_value() {
	a := 10
	fmt.Println(a)
	changeValue_pointer(&a)
	fmt.Println(a)
}

//Slices — Pass by Value but Shared Underlying Array
func modifySlice(s []int) {
	s[0] = 99
}

func Slice_Pass_by_value() {
	nums := []int{1, 2, 3}
	modifySlice(nums)
	fmt.Println("nums =", nums) // Output: nums = [99 2 3]
}

//Maps
func modifyMap(m map[string]int) {
	m["apple"] = 100
}

func Map_Pass_by_value() {
	fruits := map[string]int{"apple": 1, "banana": 2}
	modifyMap(fruits)
	fmt.Println("fruits =", fruits) // Output: map[apple:100 banana:2]
}

//channels
func send(ch chan int) {
	ch <- 10
}

func Channel_Pass_by_value() {
	ch := make(chan int, 1)
	send(ch)
	fmt.Println(<-ch) // Output: 10
}
