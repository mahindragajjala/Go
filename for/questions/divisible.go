package questions

import "fmt"

/*
	Write a program to count how many numbers between 1 and 100 are divisible by 3.

	➡️ Goal: Practice condition checking (%) and use a counter variable.
*/
func DivisibleByNumber(x int) {
	var count int
	for i := 0; i <= 100; i++ {
		if i%x == 0 {
			count++
		}
	}
	fmt.Println(count)
}
func DivisibleByNumber_userinput() {
	DivisibleByNumber(3)
}
