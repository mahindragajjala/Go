package questions

import "fmt"

/*
Write a program to find the sum of all even numbers between 1 and 100 using a for loop.

➡️ Goal: Combine loop + condition + accumulator logic.

*/
func Sum_of_all_even_numbres() {
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
