package questions

import (
	"fmt"
	"sort"
)

/* Write a function that checks whether a key exists in a map or not.*/

var Data map[int]string

func Inserting_data() {
	Data = make(map[int]string)
	Data[1] = "one"
	Data[2] = "two"
	Data[3] = "three"
	Data[4] = "four"
	Data[5] = "five"
	Data[6] = "six"

}
func Checking() {
	for key, value := range Data {
		fmt.Println(key, value)
	}
}

/*
output: PS C:\Users\mahindra.gajjala\Desktop\GO\maps> go run main.go
1 one
2 two
3 three
4 four
5 five
6 six
PS C:\Users\mahindra.gajjala\Desktop\GO\maps> go run main.go
6 six
1 one
2 two
3 three
4 four
5 five

why ?
*/
/*
Maps in Go are unordered collections.

Go does not guarantee any order of the keys when you iterate over a map — even if you insert them in a specific order.

What's the algorithm behind this?
Under the hood:

Go maps are implemented as "hash tables".
The keys (int in your case) are hashed and placed into "buckets."
When you range over the map, Go walks through these buckets.
The order depends on factors like:

Hash function.
Map resizing (internal reallocation when elements are added).
Runtime's randomization for security.
Go's map iteration order is deliberately randomized for each run (since Go 1.12+). This prevents developers from accidentally depending on a consistent iteration order.

*/

//If you want predictable order

func Actual_order() {
	keys := make([]int, 0, len(Data))
	for k := range Data {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, Data[k])
	}

}
