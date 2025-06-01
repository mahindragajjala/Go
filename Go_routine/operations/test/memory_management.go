package test

/*
Great question! Let's break down the memory management aspects of both `Mutex` and `WaitGroup` in Go — how they behave in memory, how Go handles their lifecycle, and what you should be careful about.

---

 ✅ 1. Memory Management of `sync.Mutex`

 🔹 Where is it stored?

A `Mutex` is typically stored in:
* Stack (if it's a local variable).
* Heap (if it's a field of a struct allocated with `new()` or `make()`).

 🔹 What does it manage?
`Mutex` does not protect data automatically — you must lock and unlock manually. It only protects access to shared data from concurrent modification.

 🔹 Is it garbage collected?
Yes — the `Mutex` itself is a Go object, so:
* If no references to the `Mutex` remain, Go's garbage collector will clean it up.
* It does not lock system-level resources forever — once it's unused, memory is reclaimed.

 🔹 Important Notes:
* A `Mutex` should not be copied after use — this is a runtime error.
* If you pass it by value, you may end up with two separate copies, leading to bugs.


type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}
In the example above, `mu` is part of the struct — Go manages memory properly as long as you don’t copy the struct around.



 ✅ 2. Memory Management of `sync.WaitGroup`

 🔹 Where is it stored?

Like `Mutex`, a `WaitGroup`:

* Can be on stack or heap, depending on how you declare it.
* If used as a field of a struct, it's heap-allocated with the struct.

 🔹 What does it manage?

It internally keeps a counter (int32) that:

* Increments with `Add(n)`
* Decrements with `Done()` (which calls `Add(-1)`)
* Blocks with `Wait()` until the counter reaches 0

This counter is stored in the `WaitGroup`'s internal memory.

 🔹 Is it garbage collected?

Yes — like `Mutex`, it's a regular Go object, so:

* When it’s no longer referenced, Go’s GC will reclaim its memory.
* But you must not copy a `WaitGroup` after it’s in use (same reason as `Mutex`).

---

 ❗ Memory Safety Rules (VERY IMPORTANT)

| Rule                                                                  | Applies To       | Why It Matters                                        |
| --------------------------------------------------------------------- | ---------------- | ----------------------------------------------------- |
| Do not copy a `Mutex` or `WaitGroup` once in use                  | Both             | Copying causes undefined behavior or panics.          |
| Use pointers if passing structs containing `Mutex` or `WaitGroup` | Both             | Ensures only one instance is referenced.              |
| Use `defer` for `Unlock()` and `Done()`                               | Mutex, WaitGroup | Prevent memory leaks, stuck goroutines, or deadlocks. |
| Never `Unlock()` a `Mutex` you haven’t `Lock()`ed                     | Mutex            | Runtime panic.                                        |
| Don't call `Add()` after `Wait()` has started                     | WaitGroup        | Race condition; can lead to panic.                    |

---

 ✅ What Happens Under the Hood?

 🔹 `sync.Mutex`:

* Uses futexes (fast user-space mutex) on Linux internally.
* It's lightweight until contention happens.
* If multiple goroutines try to lock, Go manages a queue of waiters.

 🔹 `sync.WaitGroup`:

* Just maintains a counter (atomic operations) and waiters (goroutines waiting on `Wait()`).
* Uses low-level primitives (like atomic counters and condition variables).
* Very lightweight and efficient for managing 1000s of goroutines.

---

 ✅ When Memory Leaks Can Happen (Common Mistakes)

1. Goroutines never finish because you forget to call `wg.Done()`.
2. You deadlock a `Mutex` by locking and never unlocking.
3. You copy a struct with `Mutex` or `WaitGroup` — causing undefined behavior.

---

 ✅ Example: Struct with Mutex and WaitGroup


type SafeMap struct {
	mu sync.Mutex
	wg sync.WaitGroup
	m  map[string]int
}

func (s *SafeMap) Add(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = val
}


Here:

* `mu` protects the shared map.
* `wg` could be used if you're tracking when multiple goroutines finish updating the map.

---

 ✅ Summary

| Feature          | Memory Managed by Go GC | Stores in Stack/Heap | Needs Manual Cleanup? | Dangerous if Copied? |
| ---------------- | ----------------------- | -------------------- | --------------------- | -------------------- |
| `sync.Mutex`     | ✅ Yes                   | Stack or Heap        | ❌ No                  | ✅ Yes                |
| `sync.WaitGroup` | ✅ Yes                   | Stack or Heap        | ❌ No                  | ✅ Yes                |

---

If you want, I can show you memory profiling using the `pprof` tool to see how much memory `Mutex` and `WaitGroup` use in a real Go application. Let me know!

*/
