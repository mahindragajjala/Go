package operatorspunctuationdata

import "fmt"

/*
Syntax	  Usage

:=	      Short variable declaration
label:	  Custom loop/goto labels
case:	  In switch statements
default:  Default branch in switch

*/
func Colon() {

	//Short Variable Declaration (very common)

	name := "GoLang"

	fmt.Println(name)

	//Labels (with :) — for goto, break, continue

outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				break outer // breaks the outer loop
			}
			fmt.Println(i, j)
		}
	}

	//case: and default: — in switch statements
	switch day := "Mon"; day {
	case "Mon":
		fmt.Println("Monday")
	default:
		fmt.Println("Another day")
	}

}
