package traversalpatterns

/*
 🔹 1. Forward Traversal

 ✅ Definition:

In forward traversal, you iterate through a data structure from the beginning to the end — usually from index `0` to `length - 1`.

* Printing elements in order.
* Searching for a value.
* Applying transformations or filters.

for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

 🔹 2. Reverse Traversal

In reverse traversal, you iterate from the end to the beginning — from index `length - 1` to `0`.

* Reversing a string/array.
* Processing suffixes.
* Finding the last occurrence of an element.

for i := len(arr) - 1; i >= 0; i-- {
    fmt.Println(arr[i])
}


 🧠 Coding Questions: Forward and Reverse Traversal

Here's a categorized list of questions using these two patterns.



 🔹 Forward Traversal – Coding Questions

1. Print all elements of an array.
2. Find the sum of all elements in an array.
3. Count the number of even and odd elements.
4. Find the maximum and minimum in an array.
5. Linear search for a given element.
6. Print characters of a string one by one.
7. Count the number of vowels in a string.
8. Check if a string is a palindrome using forward traversal (build a reversed string separately).
9. Find the first repeating element in an array.
10. Count occurrences of a target number.



 🔹 Reverse Traversal – Coding Questions

1. Print all elements of an array in reverse order.
2. Reverse a string manually using loop.
3. Reverse an array in-place.
4. Find the last occurrence of a target element.
5. Remove trailing zeroes from an integer array.
6. Print string characters in reverse order.
7. Reverse words in a sentence (without using built-in functions).
8. Detect palindrome using reverse traversal only.
9. Implement reverse iteration without using built-in `reverse()` function.
10. Print suffixes of a string in reverse order.



 🔁 Combined/Advanced Use Cases

1. Compare elements from both ends (two-pointer method).
2. Palindrome check using forward and reverse traversal together.
3. Find pairs that sum to a given value from both ends.
4. Zig-zag traversal (forward one row, reverse next).
5. Spiral traversal of a matrix.



Would you like detailed explanations and solutions for each of the questions above in Go (Golang)?

*/
