package questions

import "fmt"

/*
Write a Go program that creates a map where keys are words and values are their lengths, from a given slice of strings.
*/

func StringToMap() map[string]int {
	NewMap := make(map[string]int)
	slice := []string{"one", "two", "three"}
	for _, value := range slice {
		NewMap[value] = len(value)
	}
	fmt.Println(NewMap)
	return NewMap
}
