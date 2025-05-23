package twopointertechnique

/*
The Two Pointer Technique is a fundamental algorithmic strategy in Data Structures and Algorithms (DSA). It is especially useful when dealing with sorted arrays or linked lists. The technique uses two pointers (or indices) to traverse the data structure, either from both ends or in a sliding manner, to solve a problem more efficiently.

 🧠 When to Use Two Pointer Technique
* The array or list is sorted.
* You are searching for pairs or triplets satisfying a condition.
* You want to remove duplicates, reverse, or partition data in-place.
* You need to implement sliding windows.

 ✅ Common Problems Solved Using Two Pointers
1. Pair Sum in a Sorted Array
2. Remove Duplicates from Sorted Array
3. Reverse an Array or Linked List
4. Container With Most Water
5. Trapping Rain Water (advanced use)
6. Merging Two Sorted Arrays

 🧭 How It Works — Intuition
* Initialize two pointers, typically:
  * `left = 0`
  * `right = n - 1` (or another initial position depending on the problem)
* Use a while loop until the pointers meet (`left <= right`)
* Depending on the logic (sum check, comparison, etc.), move one or both pointers
* This avoids nested loops and often reduces O(n²) to O(n)

 🔁 Flowchart
									Start
										|
										V
									[Initialize pointers]
									left = 0
									right = n-1
										|
										V
									[Check loop condition: left <= right?]
										|
									Yes
										|
										V
									[Check condition for current pointers]
										|
										|> Condition Met? > Yes > [Store Result / Take Action]
										|
										No
										|
										V
									[Move left or right pointer accordingly]
										|
										V
									[Repeat Loop]
										|
									No (when left > right)
										|
										V
									End

 🧾 Example: Find Pair with Target Sum
# ❓Problem
Given a sorted array of integers, find if there exists a pair whose sum is equal to a target.
# 🔍 Input
arr = [1, 2, 3, 4, 6]
target = 6


# ✔️ Output

`true` → because 2 + 4 = 6

 🧪 Pseudocode

function hasPairWithSum(arr, target):
    left = 0
    right = length(arr) - 1

    while left < right:
        current_sum = arr[left] + arr[right]

        if current_sum == target:
            return true

        else if current_sum < target:
            left = left + 1  // Need a bigger sum

        else:
            right = right - 1  // Need a smaller sum

    return false

 ⚙️ Dry Run Example

arr = [1, 2, 3, 4, 6]
target = 6

Initial: left = 0 (1), right = 4 (6)

1 + 6 = 7 → too big → move right ←
left = 0 (1), right = 3 (4)

1 + 4 = 5 → too small → move left →
left = 1 (2), right = 3 (4)

2 + 4 = 6 → match found → return true




 🔄 Variants of Two Pointer Technique

| Type                   | Description                    | Example               |
| - |  |  |
| Opposite Direction | Pointers start at two ends     | Sorted array pair sum |
| Same Direction     | Both move from left to right   | Remove duplicates     |
| Fixed + Moving     | One pointer fixed, other moves | Triplet sum           |



 💡 Tips

* Often used with sorted arrays.
* Helps avoid brute force (O(n²)) solutions.
* Can be applied to strings, arrays, or linked lists.
* Can combine with other techniques like sliding window or binary search.



If you’d like, I can also explain code implementation in Go or give visual simulation. Would you like to proceed with that?

*/
