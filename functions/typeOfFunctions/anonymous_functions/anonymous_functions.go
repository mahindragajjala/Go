package anonymousfunctions

/*
An anonymous function is a function without a name. It can be assigned to a variable, passed as an argument, or executed immediately.


Anonymous functions are used when you:

Need a quick, inline function (no need to reuse elsewhere).
Want to pass logic as a parameter (e.g., callbacks or functional patterns).
Want to capture variables from the outer scope (closures).
Want to avoid polluting the global namespace with temporary function names.



| Use Case                               | Description                                  | Example                                |
| -------------------------------------- | -------------------------------------------- | -------------------------------------- |
| ✅ Callbacks                            | Used in APIs like sort, event handlers, etc. | `sort.Slice()`                         |
| ✅ Closures                             | Capture outer variables and maintain state   | Counter example                        |
| ✅ Goroutines                           | Quick, inline concurrent execution           | `go func() { ... }()`                  |
| ✅ Immediately Invoked Functions (IIFE) | To execute logic instantly and isolate scope | `result := func() int { return 42 }()` |
| ✅ Tests and Mocks                      | Temporary logic in unit tests                | `t.Run()` with inline function         |
| ✅ Functional Patterns                  | Map, filter, reduce-like behaviors           | Writing higher-order functions         |

*/
