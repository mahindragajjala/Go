package questions

import "fmt"

var Array = []int{15, 16, 6, 8, 5}
var n = len(Array)

func Normal_sorting_with_bubble_sort() {

	for i := 0; i < n-1; i++ {
		fmt.Println("PASS ", i)
		for j := 0; j < n-1; j++ {
			if Array[j] > Array[j+1] {
				temp := Array[j]
				Array[j] = Array[j+1]
				Array[j+1] = temp
			}
		}
	}
	fmt.Println(Array)
}

//To avoid unnessary comparsions we have to enhance the innser loop
func Bubble_sort_without_unnessary_loops() {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if Array[j] > Array[j+1] {
				temp := Array[j]
				Array[j] = Array[j+1]
				Array[j+1] = temp
			}
		}
	}
	fmt.Println(Array)
}

//using flags

func Using_flag_bubble_sort() {
	for i := 0; i < n-1; i++ {
		Flag := 0
		for j := 0; j < n-1-i; j++ {
			if Array[j] > Array[j+1] {
				Array[j], Array[j+1] = Array[j+1], Array[j]
				Flag = 1
			}
		}
		if Flag == 0 {
			break
		}
	}
	fmt.Println(Array)
}

/*
Version	                          Inner Loop Optimization	Early Exit (Flag)	Efficient
Normal_sorting_with_bubble_sort	      ❌	                      ❌	                ❌
Bubble_sort_without_unnessary_loops	  ✅	                      ❌	                ⚠️
Using_flag_bubble_sort (Fixed)	      ✅	                      ✅	                ✅
*/
