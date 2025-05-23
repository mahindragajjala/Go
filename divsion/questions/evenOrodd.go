package questions

import "fmt"

/*
Even or Odd Using Division

Use division (no modulo) to determine if a number is even or odd.
*/
func EvenOdd() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var Even []int
	var Odd []int

	for i := 0; i < len(input); i++ {
		num := input[i]
		if (num/2)*2 == num {
			Even = append(Even, num)
		} else {
			Odd = append(Odd, num)
		}
	}

	fmt.Println("Even:", Even)
	fmt.Println("Odd:", Odd)
}
