package operatorspunctuationdata

import "fmt"

/*
Used in function definitions to accept a variable number of arguments.
*/
func Variadic_Function_Parameters() {
	fmt.Println(sum(1, 2, 3))        // 6
	fmt.Println(sum(10, 20, 30, 40)) // 100
}
func sum(nums ...int) int {
	fmt.Println("nums--->", nums)
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

//Used when passing a slice to a variadic function.
func Slice_Expansion_in_Function_Calls() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(sum(numbers...))
}

//Array_Size_Inference
func Array_Size_Inference() {
	array := [...]int{10, 20, 30, 56}
	fmt.Println(len(array)) // Output: 3
}
