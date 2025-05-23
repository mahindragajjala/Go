package array

/*
Great! Since you’re exploring arrays deeply, a memory visualization will absolutely help reinforce how arrays work under the hood — especially when you're dealing with low-level languages like C or system-aware languages like Go.

🔍 Let's Visualize an Array in Memory

✅ C Code Example:

int arr[4] = {10, 20, 30, 40};
Let’s assume `int` = 4 bytes and base address is `1000`.

🧠 Memory Layout (Contiguous Storage)

| Index   | Value | Memory Address |
| ------- | ----- | -------------- |
| arr\[0] | 10    | 1000           |
| arr\[1] | 20    | 1004           |
| arr\[2] | 30    | 1008           |
| arr\[3] | 40    | 1012           |

📊 Diagram

Memory:
+--------+--------+--------+--------+
| 1000   | 1004   | 1008   | 1012   |
|  10    |  20    |  30    |  40    |
+--------+--------+--------+--------+
   ↑        ↑        ↑        ↑
 arr[0]  arr[1]   arr[2]   arr[3]

 🔍 Now Let’s Do the Same in Go

 ✅ Go Code:
package main

import "fmt"

func main() {
    arr := [4]int{10, 20, 30, 40}
    fmt.Printf("%p %p %p %p\n", &arr[0], &arr[1], &arr[2], &arr[3])
}

🧠 Output Example:

You may see something like:
0xc000014090 0xc000014094 0xc000014098 0xc00001409c

Each element is `4 bytes` apart (Go also uses contiguous memory layout like C).
📊 Diagram (Go)
Heap/Stack (depending on allocation):
+-------------+-------------+-------------+-------------+
| 0x...4090   | 0x...4094   | 0x...4098   | 0x...409C   |
|    10       |    20       |    30       |    40       |
+-------------+-------------+-------------+-------------+
   ↑             ↑             ↑             ↑
 arr[0]       arr[1]        arr[2]        arr[3]

> Note: Go's memory addresses might be on the heap or stack depending on escape analysis. But the layout is always contiguous.

 ⚠️ Key Takeaways:

* Arrays = contiguous memory.
* Direct access using formula:
  `address = base + (index × size_of_element)`
* Go & C store arrays similarly in memory.
* Go may store arrays on stack or heap based on escape analysis.
* Understanding this helps with performance, especially in system-level coding.
*/
