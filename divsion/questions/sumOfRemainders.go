package questions

import "fmt"

/*
Sum of Remainders

Given two numbers n and k, calculate the sum of all remainders of i % k for i from 1 to n.
*/

func Sum_of_Remainders() {
	n := 5
	k := 8
	var remainders []float64
	var SumOfRemainders int
	for i := 1; i <= n; i++ {
		data := i % k
		remainders = append(remainders, float64(data))
	}
	for _, value := range remainders {
		SumOfRemainders += int(value)
	}
	fmt.Println(SumOfRemainders)
}
