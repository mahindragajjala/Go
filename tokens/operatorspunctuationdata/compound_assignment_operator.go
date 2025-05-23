package operatorspunctuationdata

import "fmt"

//compound assignment operator
/*
Operator	Meaning
-=			x = x - y
*=			x = x * y
/=			x = x / y
%=			x = x % y
*/
func Compound_assignment_operator() {
	var Go string = "GO"
	var Lang string = "Lang"

	Go += Lang
	// Go = Go + Lang
	fmt.Println("golang:", Go)

}
