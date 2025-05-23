package operatorspunctuationdata

import "fmt"

func BitwiseOperator() {
	//single bitwise_operator
	var ab int = 1111
	var bb int = 2222

	bbb := ab | bb

	fmt.Printf("ab:%b\nbb:%b\nsingle bitwise_operator:%b\n%v\n", ab, bb, bbb, bbb)

	//|= -  compound assignment operator used for bitwise OR and assign
	var one int = 55
	var two int = 66
	fmt.Printf("one:%b\n", one)
	fmt.Printf("two:%b\n", two)
	one |= two
	fmt.Printf("one:%b\n", one)

	//dual logical_operator
	var a bool = true
	var b bool = true
	var c bool = false
	var d bool = false
	var e bool = a || b
	var f bool = a || c
	var g bool = c || a
	var h bool = c || d

	fmt.Println("dual logical_operator--->", e, f, g, h)

}
