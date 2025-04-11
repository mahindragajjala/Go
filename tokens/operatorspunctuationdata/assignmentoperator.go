package operatorspunctuationdata

import "fmt"

func Assignment_operator() {
	//Basic Assignment
	//Assigning One Variable to Another
	var a int = 5
	var b int = 6
	a = b
	fmt.Println(a)
	fmt.Println(b)

	// Short Declaration with Assignment (:=)
	data := 56
	datastring := "data"
	fmt.Println(data, datastring)

	// Reassignment (Overwriting)
	data = 65
}
