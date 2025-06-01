package questions

import "fmt"

var ClassA = map[string]int{}
var ClassB = map[string]int{}
var NewClass = map[string]int{}

func Merge_Classes() {
	NewClass = make(map[string]int)
	ClassA = make(map[string]int)
	ClassA["one"] = 1
	ClassA["two"] = 2
	ClassA["three"] = 3
	ClassA["four"] = 4

	ClassB = make(map[string]int)
	ClassB["one"] = 1
	ClassB["two"] = 2
	ClassB["three"] = 3
	ClassB["four"] = 4

	for key, values := range ClassA {
		valuesd, ok := ClassB[key]
		if ok {
			NewClass[key] = values + valuesd
		} else {
			NewClass[key] = values
		}
	}
	for x, y := range NewClass {
		fmt.Println(x, y)
	}
}
