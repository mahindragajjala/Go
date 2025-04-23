package questions

import "fmt"

/*
Count Divisible Numbers in a Range

Count how many numbers between L and R are divisible by k.
*/

func Count_divisible_numbers() {
	var L int = 1
	var R int = 10
	var K int = 2
	var count int = 0
	for i := L; i <= R; i++ {
		if i%K == 0 {
			count++
		} else {
			fmt.Println("not divisible")
		}
	}
	fmt.Println(count)
}
