package content

/*
Map types

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type,
called the key type. The value of an uninitialized map is nil.

MapType = "map" "[" KeyType "]" ElementType .
KeyType = Type .

The comparison operators == and != must be fully defined for operands of the key type; thus the key
type must not be a function, map, or slice. If the key type is an interface type, these comparison
operators must be defined for the dynamic key values; failure will cause a run-time panic.

map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}

The number of map elements is called its length.
For a map m, it can be discovered using the built-in function len and may change during execution.
Elements may be added during execution using assignments and retrieved with index expressions;
they may be removed with the delete and clear built-in function.

A new, empty map value is made using the built-in function make, which takes the map type and an optional
capacity hint as arguments:

make(map[string]int)
make(map[string]int, 100)

The initial capacity does not bound its size: maps grow to accommodate the number of items
stored in them, with the exception of nil maps. A nil map is equivalent to an empty map
except that no elements may be added.

Absolutely — let’s break this down **point by point** in clear, detailed language!

---

1️⃣ What is a Map?**

A **map** is a built-in data type in Go, used to store **key-value pairs**.
It’s like a dictionary in Python or an object in JavaScript.

- Each **key** is unique.
- Each **key** maps to a **value**.
- The type of keys and values is fixed when you declare the map.

👉 **Example:**
```go
var m map[string]int
```
This declares `m` as a map with:
- `string` as the **key type**.
- `int` as the **value type**.

---

2️⃣ Syntax**

```go
MapType = "map" "[" KeyType "]" ElementType .
```
- `map` is the keyword.
- `[KeyType]` defines the type of the keys.
- `ElementType` defines the type of the values.

👉 **Example:**
```go
map[string]int  // Key: string, Value: int
```

---

3️⃣ Key Type Rules**

The **key type** must support comparison operations (`==` and `!=`).

✅ Valid key types:
- `int`
- `string`
- `float64`
- `struct` (if all its fields are comparable)
- pointers
- interface (but the actual value must be comparable)

❌ Invalid key types:
- `slice`
- `map`
- `function`

👉 **Why?**
Because Go uses a hash table internally for maps, and the keys must be comparable to compute hashes and detect duplicates.

---

4️⃣ Map Examples**

```go
map[string]int              // Maps string to int.
map[*T]struct{ x, y float64 } // Maps pointer to a struct.
map[string]interface{}      // Maps string to any type.
```

- The first one maps words to numbers.
- The second one maps pointers to structs.
- The third one maps strings to values of any type (`interface{}`).

---

5️⃣ Length of a Map

You can get the number of items with `len()`:

m := map[string]int{"a": 1, "b": 2}
fmt.Println(len(m))  // Output: 2
```

Maps are **dynamic** — you can add or remove elements anytime, so the length can change.

6️⃣ Adding Elements

Use **assignment** to add or update a value:
m["key"] = 42  // Adds or updates "key"

7️⃣ Retrieving Elements
Use an **index expression**:
value := m["key"]
If the key doesn’t exist, `value` will be the **zero value** for that type (e.g., `0` for int).

8️⃣ Removing Elements**

Use `delete()` to remove a key:
delete(m, "key")

Since Go 1.21, you can use `clear()` to remove **all keys**:

clear(m)  // Empties the map


9️⃣ Creating Maps with make()**

To create a new, empty map:
m := make(map[string]int)
You can also provide a capacity hint:
m := make(map[string]int, 100)
This tells Go to optimize for ~100 entries, but the map will grow beyond that if needed.


🔴 nil Maps vs Empty Maps**

A `nil` map:
var m map[string]int
- Is like a "zero" map.
- You cannot add elements to it — it will panic.

An empty map:
m := make(map[string]int)
- Can grow freely.
- Is ready to store key-value pairs.

✅ **Summary:**
- Maps store key-value pairs.
- Keys must be comparable.
- Values can be anything.
- `len()`, `delete()`, `clear()` manage the contents.
- `make()` creates maps.
- `nil` maps cannot hold new data.

---

If you like, I can also give you a **real-world Go example** like a student database or URL shortener using maps!
Want me to do that?
*/
