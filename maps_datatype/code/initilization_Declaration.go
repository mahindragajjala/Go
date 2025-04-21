package code

import "fmt"

// ----------------------------nil map------------------------

var Nil_Map map[string]int //nil by default
// Nil_Map["adfd"] = 25 // can't work

//Function Return When No Data
func FetchUsers() map[string]string {
	return nil // No users found or DB error
}

//Signaling ‘Unset’ or ‘Not Ready’ State
func Unset_NotReady() {
	var myMap map[string]int
	if myMap == nil {
		fmt.Println("Map not initialized.")
	}
	if myMap == nil {
		myMap = make(map[string]int)
	}
	myMap["help"] = 100
	fmt.Println(myMap)
}

//Efficient Zero Memory Until Needed
func Efficient_Zero_Memory_Function() {
	var bigMap map[string]int     // nil, no memory allocated
	bigMap = make(map[string]int) // when needed
	bigMap["one"] = 1
	fmt.Println(bigMap)
}

//TO insert data while initilization or to insert the data in the nil map
func Initialization_Maps() {
	var NewMap = make(map[string]int)

	NewMap["one"] = 1
	NewMap["one"] = 2
	NewMap["one"] = 3
	fmt.Println(NewMap["one"])
}

/*

✅ Valid key types:
- `int`
- `string`
- `float64`
- `struct` (if all its fields are comparable)
- pointers
- interface (but the actual value must be comparable)

❌ Invalid key types:
- `slice`
- `map`
- `function`
*/
var M = map[string]int{"a": 1, "b": 2}

// Declare globally
var myMap map[string]int

func Declare_global() {
	// Initialize inside a function
	myMap = make(map[string]int)

	// Add elements
	myMap["apple"] = 10
	myMap["banana"] = 20

	fmt.Println(myMap) // Output: map[apple:10 banana:20]
}

func Length_Map() {
	fmt.Println(len(M)) // Output: 2
}

func Adding_Map() {
	M["c"] = 3
}

func Retrieving_Elements() {
	fmt.Println(M["a"])
}
func Removing_Elements() {
	//For one element
	delete(M, "a")
	//Clear all
	clear(M)
}

//Capacity hint
func Capacity_hint() {

	m := make(map[string]int, 100)

	fmt.Println(m)
}
