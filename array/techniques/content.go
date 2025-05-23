package techniques

/*
In Data Structures and Algorithms (DSA), various problem-solving techniques help simplify and optimize the solution. Here are popular and essential techniques (like the Two Pointer Technique) that every DSA learner should know:



 🔁 1. Two Pointer Technique

* Use Case: Arrays, strings, linked lists
* Idea: Use two pointers to traverse from start and end or move conditionally.
* Examples:

  * Check if an array has a pair with a given sum (sorted array).
  * Reverse an array in-place.
  * Move zeroes to the end of the array.



 🔄 2. Sliding Window Technique

* Use Case: Arrays, strings
* Idea: Use a window (subarray) that moves forward to track a subset.
* Examples:

  * Maximum sum subarray of size `k`.
  * Longest substring without repeating characters.
  * Minimum window substring.



 🔍 3. Binary Search Technique

* Use Case: Sorted arrays, search problems
* Idea: Divide and conquer approach to reduce search space by half.
* Examples:

  * Find element in a sorted array.
  * Search in rotated sorted array.
  * Find square root of a number.



 🧮 4. Prefix Sum / Difference Array

* Use Case: Cumulative sum, range queries
* Idea: Preprocess array to answer range sum queries quickly.
* Examples:

  * Sum of elements in a subarray.
  * Number of occurrences in range.
  * Range update queries.



 🔁 5. Fast & Slow Pointers (Floyd's Cycle)

* Use Case: Linked Lists, cycle detection
* Idea: Two pointers moving at different speeds to detect loops.
* Examples:

  * Detect cycle in a linked list.
  * Find the starting point of a loop.
  * Find middle of a linked list.



 🔁 6. Recursion & Backtracking

* Use Case: Decision problems, permutations, path finding
* Idea: Try all possibilities, undo changes (backtrack) when needed.
* Examples:

  * N-Queens Problem
  * Sudoku Solver
  * Subset/Permutation generation



 🎯 7. Greedy Technique

* Use Case: Optimization problems
* Idea: Take the best choice at each step without reconsidering.
* Examples:

  * Activity Selection
  * Fractional Knapsack
  * Huffman Coding



 📐 8. Divide and Conquer

* Use Case: Recursive problems
* Idea: Break problem into subproblems, solve recursively, combine.
* Examples:

  * Merge Sort
  * Quick Sort
  * Binary Search



 📊 9. Dynamic Programming (DP)

* Use Case: Overlapping subproblems, optimization
* Idea: Store results of subproblems to avoid recomputation.
* Examples:

  * Fibonacci
  * 0/1 Knapsack
  * Longest Common Subsequence



 🧱 10. Bit Manipulation

* Use Case: Optimization, low-level operations
* Idea: Use bitwise operators for efficient calculations.
* Examples:

  * Check if number is power of two
  * Count set bits
  * Swap two numbers without temp variable



 🌲 11. Union Find / Disjoint Set

* Use Case: Graphs (connectivity), clustering
* Idea: Track disjoint sets with union and find operations.
* Examples:

  * Detect cycle in an undirected graph
  * Kruskal’s Algorithm (MST)
  * Connected components



 ⛓ 12. Topological Sorting (DFS / BFS)

* Use Case: Graphs with dependencies (DAG)
* Idea: Linear ordering of vertices where `u -> v`, `u` appears before `v`.
* Examples:

  * Course schedule
  * Task scheduling with dependencies



 🔁 13. Monotonic Stack / Queue

* Use Case: Range queries, Next Greater Element
* Idea: Stack/Queue that keeps elements in increasing/decreasing order.
* Examples:

  * Largest Rectangle in Histogram
  * Next Greater Element
  * Sliding Window Maximum

  | Technique                         | Use Case / When to Use                                  | Common Examples                                           |
|--------------------------------- | ------------------------------------------------------- | --------------------------------------------------------- |
|Two Pointer**                   | Sorted arrays, pair problems, in-place reversal         | Pair sum, reverse array, remove duplicates                |
|Sliding Window**                | Fixed or variable-length subarray/string problems       | Max sum subarray, longest unique substring, anagram check |
|Binary Search**                 | Sorted arrays/lists or answer lies in a search space    | Find element, peak element, square root, min pages        |
|Prefix Sum / Difference Array** | Range queries, cumulative sum/range update              | Subarray sum, count frequency, range updates              |
|Fast & Slow Pointer**           | Cycle detection, middle of linked list                  | Detect loop, find middle, intersection of linked lists    |
|Recursion / Backtracking**      | Try all combinations, undo steps if needed              | N-Queens, permutations, subsets, Sudoku                   |
|Greedy**                        | Choose local optimum for global solution                | Activity selection, Huffman coding, fractional knapsack   |
|Divide and Conquer**            | Break into subproblems and combine                      | Merge sort, quicksort, closest pair                       |
|Dynamic Programming**           | Overlapping subproblems, optimal substructure           | Fibonacci, LCS, LIS, 0/1 knapsack                         |
|Bit Manipulation**              | Binary operations, performance-critical tasks           | Count set bits, power of 2, XOR problems                  |
|Union Find (DSU)**              | Graph connectivity, disjoint set problems               | Cycle detection, Kruskal's MST, connected components      |
|Topological Sort**              | Task scheduling, dependency resolution                  | Course schedule, build order, Kahn’s algorithm            |
|Monotonic Stack/Queue**         | Next greater/smaller, histogram, sliding window max/min | NGE, largest rectangle, sliding window max                |

*/
