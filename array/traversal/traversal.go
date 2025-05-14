package traversal

/*
Traversing an array is one of the most fundamental and powerful concepts in programming. Understanding how to traverse arrays in Go, from basic to advanced levels, will help you master loops, memory access, algorithm building, and even optimizations.

 🧭 1. What is Array Traversal?
Array Traversal means visiting each element in an array one by one (in sequence or custom order) to perform an operation like:
* Displaying values
* Calculating a sum or average
* Searching
* Updating/modifying
* Comparing

 🧊 2. Basic Traversal (Forward Loop)
 ✅ Using `for` loop:
arr := [5]int{10, 20, 30, 40, 50}
for i := 0; i < len(arr); i++ {
    fmt.Println("Index:", i, "Value:", arr[i])
}
 🔍 Output:
Index: 0 Value: 10
Index: 1 Value: 20

 🔁 3. Traversal Using `range` (Idiomatically in Go)
 ✅ Example:
for index, value := range arr {
    fmt.Printf("arr[%d] = %d\n", index, value)
}
* `range` returns index and value
* Clean and preferred in Go

 🔄 4. Reverse Traversal
for i := len(arr) - 1; i >= 0; i-- {
    fmt.Println("Reverse:", arr[i])
}

 🔂 5. Traversal by Steps (Skipping Elements)
for i := 0; i < len(arr); i += 2 {
    fmt.Println("Every 2nd Element:", arr[i])
}

 🎯 6. Conditional Traversal

 ✅ Print only even numbers:
for _, val := range arr {
    if val%2 == 0 {
        fmt.Println("Even:", val)
    }
}

 🔍 7. Advanced Traversal Techniques

 A. Two-Pointer Traversal (Start and End)
start := 0
end := len(arr) - 1

for start < end {
    fmt.Println("Pair:", arr[start], arr[end])
    start++
    end--
}

 B. Traversing Arrays of Structs
type Student struct {
    name string
    marks int
}

students := [2]Student{
    {"Alice", 80},
    {"Bob", 90},
}

for _, student := range students {
    fmt.Println(student.name, student.marks)
}

 🧠 8. Operations During Traversal

| Operation     | Purpose                           | Example                      |
| ------------- | --------------------------------- | ---------------------------- |
| Sum           | Add all elements                  | `sum += arr[i]`              |
| Modify        | Update elements                   | `arr[i] *= 2`                |
| Search        | Find a value                      | `if arr[i] == key`           |
| Count         | Count values matching a condition | `if arr[i] > 50 { count++ }` |
| Find Max/Min  | Track highest/lowest              | `if arr[i] > max` etc.       |
| In-place Swap | Reverse or rotate                 | `arr[i], arr[n-i-1] = ...`   |

 🔍 9. Topic-wise Traversal Patterns

| Pattern          | Description                      | Example Use Case                  |
| ---------------- | -------------------------------- | --------------------------------- |
| Full Linear      | Visit each element               | Sum, average                      |
| Conditional      | Visit only matching elements     | Print odd/even, filter values     |
| Reverse          | From end to start                | Reverse display, palindromes      |
| Two-pointer      | Start from both ends             | Palindrome check, pair sums       |
| Nested Traversal | Loop inside loop                 | 2D arrays, comparisons            |
| Sliding Window   | Maintain subarrays of fixed size | Subarray sums, max sliding window |
| In-place         | Modify without extra space       | Reversals, sorting                |

 🔁 10. 2D Array Traversal
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
    }
}

 🔍 11. Recursive Traversal
Advanced: Use recursion instead of a loop
func traverse(arr []int, i int) {
    if i == len(arr) {
        return
    }
    fmt.Println(arr[i])
    traverse(arr, i+1)
}
Call:
traverse([]int{10, 20, 30}, 0)

 🧪 12. Common Interview Questions on Traversal

1. Reverse an array in-place
2. Find second largest element
3. Move all zeros to end
4. Count unique elements
5. Remove duplicates from array
6. Rotate array by K steps
7. Check if array is palindrome
8. Find pair with sum = X
 🧠 Summary Table

| Topic                 | Loop Type         | Notes                           |
| --------------------- | ----------------- | ------------------------------- |
| Basic Traversal       | `for i := 0`      | Manual index                    |
| Idiomatic Go          | `for _, val`      | Clean, readable                 |
| Reverse Traversal     | `for i := n`      | Useful for palindromes, reverse |
| Condition-based       | `if` inside       | Filtering                       |
| Nested Traversal (2D) | `for i` + `for j` | Matrices, graphs, etc.          |
| Recursion             | Function call     | Memory-depth control            |
| Two-pointer/Window    | Pattern-based     | For optimized algorithms        |
*/
