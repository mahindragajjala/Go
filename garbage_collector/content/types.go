package content

/*
| Data Type / Structure | Stored In Stack                                                | Stored In Heap                          | Why / Notes                                                               |
| ------------------------- | ------------------------------------------------------------------ | ------------------------------------------- | ----------------------------------------------------------------------------- |
| Basic variables       | ✅ Yes (e.g., `int`, `float64`, `bool`, `string` inside a function) | ❌ Only if escaped                           | Stack-allocated by default unless captured by closures or returned as pointer |
| Pointers              | ✅ Pointer value is on stack                                        | ✅ Pointed-to data may be on heap            | Actual pointer lives on stack; heap stores data if it escapes                 |
| Structs (non-pointer) | ✅ Local structs                                                    | ✅ Heap if passed as pointer or escapes      | Escape analysis decides                                                       |
| Structs (via pointer) | ❌                                                                  | ✅ Yes                                       | Useful when returning from a function or sharing across goroutines            |
| Slices                | ✅ Slice header (len, cap, ptr)                                     | ✅ Underlying array is on heap (usually)     | Header on stack, backing array on heap                                        |
| Arrays (non-pointer)  | ✅ Yes (within stack size limits)                                   | ❌ Usually stack-only unless escaped         | Large arrays may be promoted to heap                                          |
| Maps                  | ❌                                                                  | ✅ Always heap                               | Maps are reference types                                                      |
| Channels              | ❌                                                                  | ✅ Always heap                               | Channels involve internal buffers/goroutines                                  |
| Functions (closures)  | ❌                                                                  | ✅ Yes, if they capture outer variables      | Captured variables must be heap-allocated                                     |
| Interfaces            | ✅ Interface value (type + data)                                    | ✅ If data inside is heap-allocated          | Depends on what is stored inside                                              |
| Goroutine stack       | ✅ Separate stack per goroutine                                     | ✅ Grows on demand, allocated from heap pool | Starts small (2 KB), grows up to 1 GB                                         |
| Global variables      | ❌                                                                  | ✅ Static memory, similar to heap            | Not part of stack                                                             |

*/
