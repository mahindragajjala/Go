package operatorspunctuationdata

import "fmt"

/*
the % operator is called the modulus or remainder operator.
It gives the remainder after dividing two integers.
*/
func ModulesOperator() {
	var a int = 5
	var b int = 2
	var c int = a % b
	fmt.Println(c)
}
