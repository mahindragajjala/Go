package operatorspunctuationdata

import "fmt"

func DivisionOperator() {
	//division
	a := 10
	b := 2
	result := a / b
	fmt.Println(result) // Output: 5

	//Float Division
	c := 7.0
	d := 2.0
	var Float = c / d
	fmt.Println(Float) //output: 3.5

	//10 / 0 - Invalid	Runtime panic
}
