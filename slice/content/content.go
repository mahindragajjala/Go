package content

/*
What is a Slice?**
A slice in Go is like a view or window into an array.

- Think of an array like a full pizza.
- A slice is like taking a few slices from that pizza — you can hold only a part of the full array.
- Slices do not store data themselves — they reference the underlying array where the actual data lives.


Length and Capacity
Every slice has:
1. Length (`len`)
   - The number of elements in the slice.
   - Always `>= 0`.

2. Capacity (`cap`)
   - The total space from the slice's starting point to the end of the underlying array.
   - `len <= cap`.
Example

arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:4]  // slice from index 1 to 3
arr : Full array → `[10, 20, 30, 40, 50]`
Slice → `[20, 30, 40]`

`len(s)` = 3
`cap(s)` = 4 (`from index 1 to the end of arr`)


Uninitialized Slice
- A slice that has not been assigned any value is `nil`.

var s []int
fmt.Println(s == nil)  // true

This is different from an empty slice (`[]int{}`) — which is not nil.


Why Use Slices Instead of Arrays?
- Arrays have a fixed size at compile-time.
- Slices can grow and shrink during runtime.
- Easier and more flexible to use.

Creating Slices From an Array

arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]
`[2,3,4]` will be stored in `s`.


Using `make()`

s := make([]int, 5)  // length 5, capacity 5
s := make([]int, 5, 10)  // length 5, capacity 10
This creates a new hidden array behind the scenes.

Shared Storage
When you create multiple slices from the same array:
- Changing one slice's element affects the underlying array.
- Other slices referencing that array will see the change!

Example:
arr := [3]int{1, 2, 3}
a := arr[:2]  // [1,2]
b := arr[1:]  // [2,3]

a[1] = 100
fmt.Println(b[0])  // Outputs: 100

Capacity Explained
Let’s say:
arr := [10]int{1,2,3,4,5,6,7,8,9,10}
s := arr[2:5]  // elements [3,4,5]
len(s) = 3 (indexes 2,3,4)
cap(s) = 8 (from index 2 to index 9)

You can extend a slice within its capacity:
s = s[:7]  // now covers [3,4,5,6,7,8,9]

Slices of Slices (2D, 3D, ...)
- Go does not have true multi-dimensional slices.
- But you can create slices of slices.

Example:
matrix := [][]int{
    {1, 2, 3},
    {4, 5},
    {6},
}
```
Inner slices can be of different lengths. Unlike arrays of arrays, this is flexible.

---

### 💡 **Summary**
| Feature         | Slice                           | Array                           |
|-----------------|---------------------------------|---------------------------------|
| Size            | Dynamic (length can change)    | Fixed (set at creation)         |
| Memory          | Points to underlying array     | Owns its data                   |
| Sharing         | Multiple slices can share data | Each array is independent       |
| Creation        | `make([]T, len, cap)` or slicing | `var arr = [n]T{}`              |
| Multi-dimension | Slice of slices possible       | Arrays of arrays (fixed sizes)  |

---

✅ **Conclusion:**
- Slices = more flexible and dynamic than arrays.
- They are essential for real-world Go programming because you rarely know your data size at compile-time.
- They are lightweight references to arrays and can save memory when handled properly.

---

If you want, I can give you a few real-world examples or quiz-style questions to lock this in.
Want me to?
*/
