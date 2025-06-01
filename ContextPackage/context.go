package main

/*
| Function/Type                          | Description                                                                               |
| -------------------------------------- | ----------------------------------------------------------------------------------------- |
| `context.Background()`                 | Returns an empty context. Typically used at the top level (e.g., main or root handler).   |
| `context.TODO()`                       | Placeholder when you’re not sure which context to use yet.                                |
| `context.WithCancel(ctx)`              | Returns a copy of the context that can be canceled. Useful to stop a group of goroutines. |
| `context.WithTimeout(ctx, timeout)`    | Adds a timeout; the context is automatically canceled after the duration.                 |
| `context.WithDeadline(ctx, time.Time)` | Like `WithTimeout`, but uses a fixed time instead of a duration.                          |
| `context.WithValue(ctx, key, val)`     | Attaches key-value data to a context. Use sparingly (not for passing large/complex data). |

*/
