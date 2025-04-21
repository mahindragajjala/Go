package content

/*
Awesome — here’s a simple **visual diagram** to help you understand the difference between `array`, `slice`, and `map` memory layouts in Go!

---

1️⃣ Array Memory Layout**

When you declare:
arr := [3]int{10, 20, 30}
In Memory:
arr
│
├── [10] ── [20] ── [30]  (continuous memory block)

- Fixed size.
- Data is directly in the variable.
- No descriptor, the name `arr` is the data.


2️⃣ Slice Memory Layout**
When you declare:

s := []int{10, 20, 30}
In Memory:
s (Descriptor)
│
├── Data Pointer ──> [10] ── [20] ── [30]  (actual data on heap)
├── Length = 3
└── Capacity = 3
- Slice holds just a pointer, length, and capacity.
- Actual data lives elsewhere.
- Grows dynamically when needed.

3️⃣ Map Memory Layout**

When you declare:

m := map[string]int{"a": 10, "b": 20}
In Memory:
m (Descriptor)
│
└── Pointer ──> "Hash Table Structure" (buckets) in heap
                     ├─ "a": 10
                     └─ "b": 20
- Map holds a pointer to a hash table.
- Actual keys and values live inside buckets in the heap.
- Grows and reorganizes as more keys are added.

✅ **Summary Table**

| Type       | Data Location            | Descriptor (Meta Info) |
|------------|---------------------------|-------------------------|
| `array`    | Inline / Stack            | ❌ No descriptor         |
| `slice`    | Data on Heap or Stack     | ✅ Pointer, Len, Cap     |
| `map`      | Keys & Values in Hash Table | ✅ Pointer to hash table |

*/
