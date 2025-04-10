package identifiers

import "fmt"

//Names you create: e.g., main, myVar, doSomething.

//main word is not keyword its an identifier becoz package main/func main
/*
Main is not one of them. You could technically name other variables or functions main in other packages, though it's discouraged.
*/
func Main_identifiers() {
	var Variable_identifiers int
	var slice_identifiers = []int{}
	var map_identifiers = map[int]string{}

	fmt.Println("Variable_identifiers:", Variable_identifiers)
	fmt.Println("slice_identifiers:", slice_identifiers)
	fmt.Println("map_identifiers:", map_identifiers)
}
