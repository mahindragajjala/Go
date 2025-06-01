package questions

import "fmt"

var Map1 = map[string]int{}
var Map2 = map[string]int{}
var sum = map[string]int{}

func Init_func() {
	Map1 = make(map[string]int)
	Map1["one"] = 1
	Map1["two"] = 2
	Map1["three"] = 3
	Map1["four"] = 4

	Map2 = make(map[string]int)
	Map2["one"] = 1
	Map2["two"] = 2
	Map2["three"] = 3
	Map2["four"] = 4

	sum = make(map[string]int) // Initialize the sum map
	Merge_maps(Map1, Map2)

	fmt.Println("Merged Map:", sum)
}

// Merge_maps merges two maps of string:int by summing values for common keys
/*
Algorithm: Merge_maps
1. Iterate through all keys in the first map (x).
2. If the same key exists in the second map (y), sum both values and store in the result map (sum).
3. If the key does not exist in the second map, store the original value from the first map.
4. Iterate through all keys in the second map (y).
5. If a key from the second map is not present in the first map, add it to the result map (sum).
*/
func Merge_maps(x map[string]int, y map[string]int) {
	for key, val := range x {
		if value, exists := y[key]; exists {
			sum[key] = val + value
		} else {
			sum[key] = val
		}
	}

	for key, val := range y {
		if _, exists := x[key]; !exists {
			sum[key] = val
		}
	}
}
