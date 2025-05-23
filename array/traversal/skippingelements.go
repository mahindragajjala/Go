package traversal

import "fmt"

func SkippingElements(array [5]int) {
	for i := 0; i < len(array); i += 2 { // i := i + 2
		fmt.Println(array[i])
	}
}
func OddNumbers(array [5]int) {
	for i := 0; i < len(array); i += 2 { // i := i + 2
		fmt.Println("odd numbers :", array[i])
	}
}
func EvenNumbers(array [5]int) {
	for i := 1; i < len(array); i += 2 { // i := i + 2
		fmt.Println("even numbers :", array[i])
	}
}
