package onearray

import "fmt"

/*
size and capacity are terms commonly associated with slices, and understanding their difference is key when working with slice data structures.

Size (or Length)
Represented by len(slice)
It is the number of elements currently in the slice.
You can access elements from 0 to len(slice)-1.


Capacity
Represented by cap(slice)
It is the number of elements the slice can hold in total without reallocating.
Capacity is always ≥ size.
It is measured from the start index of the slice to the end of the underlying array.
*/

func SizeAndCapcity() {
	s := make([]int, 3, 5)
	fmt.Println("Slice:", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))

	s = append(s, 10, 20)
	fmt.Println("After append:", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))

	s = append(s, 30) // exceeds original capacity
	fmt.Println("After exceeding capacity:", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s)) // capacity will grow, typically doubled
}

/*
Slice: [0 0 0]
Length: 3
Capacity: 5

After append: [0 0 0 10 20]
Length: 5
Capacity: 5

After exceeding capacity: [0 0 0 10 20 30]
Length: 6
Capacity: 10

*/
