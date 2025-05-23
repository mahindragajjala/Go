package content

/*
Great — I love that you're thinking visually! Let's break this down like a diagram so you can clearly **see how Go's `map` works under the hood**.

Concept: Hash Table in Go**
When you create this map:

Data := map[int]string{
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
}
Go stores it in a **hash table** structure.

Visualization of a Hash Table**
 ---------------------------
| Bucket 0 | [ ]            |
 ---------------------------
| Bucket 1 | [2 => "two"]   |
 ---------------------------
| Bucket 2 | [6 => "six"]   |
 ---------------------------
| Bucket 3 | [4 => "four"]  |
 ---------------------------
| Bucket 4 | [1 => "one"]   |
 ---------------------------
| Bucket 5 | [5 => "five"]  |
 ---------------------------
| Bucket 6 | [3 => "three"] |
 ---------------------------
| Bucket 7 | [ ]            |
 ---------------------------

 How it works:

1. Hashing:
   When you insert `Data[1] = "one"` — Go uses an internal hash function like:

hash = hashFunction(1)
bucket = hash % tableSize
For example, if `tableSize` is `8`, the hash decides which bucket the key-value pair goes into.

2. Buckets:
   Buckets hold the actual key-value pairs. Each bucket might store more than one item, especially if there’s a hash collision (multiple keys mapping to the same bucket).

3. Random Order on Iteration:
   When you `range` over the map:

for key, value := range Data {
    fmt.Println(key, value)
}


Go does not read them in `1, 2, 3...` order.
Instead, it walks through the **buckets** — and the order can vary **every time** because:
- Go randomizes the starting point and the walking order for security and reliability.
- The order is not related to insertion.


Why Random?
Before Go 1.12, maps sometimes showed the same order (accidentally) — this led people to write unsafe code assuming that order.
Since Go 1.12, the runtime adds randomization so that the order **changes on each run** — this reminds you that maps are unordered.

Key Takeaway:**

- Maps are designed for **fast access** not for ordering.
- If you want ordering, you must use **slices + sorting**.

---

I can also sketch this as an **image diagram** if you’d like — want me to draw that for you? 🎨🧑‍💻
*/
