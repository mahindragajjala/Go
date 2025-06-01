package questions

import "fmt"

/*

Write a program that deletes a key from a map and prints the remaining keys.

*/
func Delete() {
	var mapdata = make(map[int]string)
	mapdata[1] = "one"
	mapdata[2] = "two"
	mapdata[3] = "three"
	mapdata[4] = "four"
	mapdata[5] = "five"
	mapdata[6] = "six"

	delete(mapdata, 1)
	fmt.Println(mapdata)
}
