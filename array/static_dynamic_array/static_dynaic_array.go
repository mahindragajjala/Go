package staticdynamicarray

/*
Great question! Understanding static vs dynamic arrays in Go will clarify how memory management, flexibility, and performance differ when dealing with collections of data. Go does not use the terms "static array" or "dynamic array" explicitly, but it supports both concepts through arrays and slices.

---

 🧊 1. What is a Static Array in Go?

 ✅ Definition:

A static array in Go is a fixed-size sequence of elements of the same type, allocated contiguously in memory. Its size is known at compile time and cannot be changed.

 🔹 Syntax:


var arr [5]int // array of 5 integers


 🔹 Example:


package main

import "fmt"

func main() {
    var arr [4]int = [4]int{10, 20, 30, 40}
    fmt.Println(arr)         // [10 20 30 40]
    fmt.Println(arr[2])      // 30
}


 🔹 Key Features:

* Size is fixed.
* Stored in contiguous memory.
* Access is fast: `O(1)` for any index.
* Does not grow/shrink.

 🔹 Memory View:


+---------+---------+---------+---------+
| arr[0]  | arr[1]  | arr[2]  | arr[3]  |
|   10    |   20    |   30    |   40    |
+---------+---------+---------+---------+


---

 ⚙️ 2. What is a Dynamic Array in Go?

Go provides slices, which are dynamic, resizable views over arrays.

 ✅ Definition:

A slice is a flexible, dynamically-sized wrapper over an array. It can grow or shrink in size as needed and provides powerful functionality while being lightweight.

 🔹 Syntax:


var slice []int


 🔹 Example:


package main

import "fmt"

func main() {
    slice := []int{10, 20, 30}
    slice = append(slice, 40, 50)
    fmt.Println(slice)         // [10 20 30 40 50]
}


 🔹 Key Features:

* Backed by an underlying array.
* Has a length and a capacity.
* Can grow beyond its initial size using `append()`.
* Go automatically allocates new memory and copies when capacity is exceeded.

---

 🧠 3. Slice Memory Internals

A slice is a data structure with:

* A pointer to the underlying array.
* A length (number of elements).
* A capacity (max size before new allocation).

 🧾 Behind the Scenes:


type slice struct {
    ptr *T   // pointer to array
    len int  // current length
    cap int  // max capacity
}


 🔹 Example:


s := []int{1, 2, 3} // len=3, cap=3
s = append(s, 4)    // len=4, cap=6 (new array allocation happens)


---

 🔍 4. How Append Works Internally (Advanced)

When you do `append(slice, newElements...)`, if `cap(slice)` is not enough:

1. A new underlying array is created (usually with double capacity).
2. Old elements are copied to the new array.
3. New elements are added.
4. The slice now points to this new array.


s := []int{1, 2, 3} // cap = 3
s = append(s, 4)    // cap increases (new array allocated)


---

 🧪 5. Static vs Dynamic Comparison

| Feature           | Static Array `[N]T`     | Slice `[]T`               |
| ----------------- | ----------------------- | ------------------------- |
| Size              | Fixed (at compile time) | Dynamic (can grow/shrink) |
| Memory Allocation | Stack or heap           | Heap (most cases)         |
| Flexibility       | Rigid                   | Highly flexible           |
| Built-in Methods  | None                    | `append()`, `copy()`, etc |
| Performance       | Slightly faster         | May involve reallocation  |
| Usage             | Low-level, fixed-size   | Real-world dynamic usage  |

---

 🧱 6. Create Slice from Static Array


arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4] // points to 20, 30, 40


* `slice[0]` → `arr[1]`
* Slices are views, not copies.

---

 🧠 7. Slice Capacity Control (Advanced)


s := make([]int, 3, 5) // len = 3, cap = 5
s = append(s, 10, 20)  // still within cap
s = append(s, 30)      // exceeds cap → new array allocated


---

 📦 8. Best Practices

| Scenario                        | Use           |
| ------------------------------- | ------------- |
| Fixed-size buffer               | Array `[N]T`  |
| Growing collection (e.g. logs)  | Slice `[]T`   |
| Function arguments              | Prefer slices |
| Performance-critical fixed data | Arrays        |

---

 🧪 Quiz for Practice

1. What happens if you append to a slice beyond its capacity?
2. Can a slice grow beyond its original array?
3. How are slices internally implemented?
4. Can you pass a static array to a function expecting a slice?

---

Would you like me to provide diagrams or Go programs that demonstrate these concepts in action (e.g., showing how memory grows)?

*/
/*
In Go, slices have two important built-in properties:
* `len(slice)` → Length: number of elements the slice currently holds.
* `cap(slice)` → Capacity: maximum number of elements the slice can hold before it needs to allocate a new underlying array.

🔹 1. What is Length (`len`)?
* The actual number of elements present in the slice.
* Changes when you add or remove elements.
s := []int{10, 20, 30}
fmt.Println(len(s)) // 3

 🔹 2. What is Capacity (`cap`)?

* The size of the underlying array the slice is using.
* Determines how many elements you can append before a new array is allocated.
s := make([]int, 3, 5) // len = 3, cap = 5

Means:
* Go creates an array of size 5.
* Your slice (`s`) points to the first 3 elements of that array.
* So you can grow the slice up to 5 elements without needing a new array.

 🔸 3. Visual Example
s := make([]int, 3, 5)
fmt.Println(s)        // [0 0 0]
fmt.Println(len(s))   // 3
fmt.Println(cap(s))   // 5
s = append(s, 100)    // len = 4, still within cap
s = append(s, 200)    // len = 5, still within cap
s = append(s, 300)    // len = 6, exceeds cap → new array created
fmt.Println(s)        // [0 0 0 100 200 300]

🧠 Memory Transition:

| Operation           | len | cap  | Notes                            |
| ------------------- | --- | ---- | -------------------------------- |
| `make([]int, 3, 5)` | 3   | 5    | Underlying array of size 5       |
| `append(100)`       | 4   | 5    | No reallocation                  |
| `append(200)`       | 5   | 5    | Still within capacity            |
| `append(300)`       | 6   | 10\* | New array allocated, often 2×cap |

 🔍 4. Why Is This Important?

 ✅ Efficiency:

Avoids repeated memory allocations during `append()` if you preallocate a larger capacity:

s := make([]int, 0, 1000) // efficient for large batch appends

 📦 5. Summary

| Property | Description                                  | Example                      |
| -------- | -------------------------------------------- | ---------------------------- |
| `len(s)` | Number of elements currently in the slice    | `len([]int{1,2,3}) → 3`      |
| `cap(s)` | Max number of elements before resize happens | `cap(make([]int, 3, 5)) → 5` |

 ✅ Bonus Tip

Use the `len()` and `cap()` functions to monitor memory usage and optimize slice operations, especially in performance-sensitive or large-scale applications.

*/
