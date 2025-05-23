package operatorspunctuationdata

import "fmt"

/*
Use Case	              Example

Function call             fmt.Println("Hi")
Function definition	      func add(a, b int)
Grouping expressions	  (a + b) * c
Logical expressions	      (x > 0 && y > 0)
Type conversion	          int(3.14)
Multiple return values	  func f() (int, string)
*/
func Parentheses() {
	//functionCall
	Parentheses_call()
	var a int = 1
	var b int = 2

	//Grouping expressions
	var c int = (a + b)
	fmt.Println(c)

	//Logical expressions
	var d bool = (a > 0 && b > 0)
	fmt.Println(d)

	//Type conversion
	var x float64 = 10.5
	var y int = int(x) // convert float64 to int
	fmt.Println(y)

	//Multiple return values
	e, f := Parentheses_Multiple_return_Value(a, b)
	fmt.Println(e, f)

}

//Function definition
func Parentheses_call() {}
func Parentheses_Multiple_return_Value(int, int) (int, int) {
	return 1, 2
}
