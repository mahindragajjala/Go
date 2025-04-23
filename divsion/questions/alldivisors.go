package questions

import "fmt"

/*
List All Divisors

Given a number n, return all numbers that divide n without a remainder.
*/
func All_divisors() {
	var input int = 522
	var divisors []int
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			divisors = append(divisors, i)
		}

	}
	fmt.Println(divisors)

}
