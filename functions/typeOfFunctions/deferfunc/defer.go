package deferfunc

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
defer delays the execution of a function until the surrounding function returns.

It's commonly used for cleanup, closing resources, logging, or error handling.
*/
func Example() {
	defer fmt.Println("Deferred call")
	fmt.Println("Normal call")
}

// Key Properties of defer
// LIFO Order (Last-In, First-Out)
func Last_In_Frist_Out() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

/*
3
2
1
*/

// Arguments are evaluated immediately.
func Test() {
	x := 5
	defer fmt.Println(x) // 5 is captured now
	x = 10
}

// Types of Use Cases (Common Patterns)
// Closing Files or Resources
func Close_function() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // Automatically closes file when function exits

}

// Unlocking Mutexes
func Unlocking_mutexes() {
	// mu.Lock()
	// defer mu.Unlock()
	// // Do some work with protected resource
}

// Recovering from Panic
func SafeDivide() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	// code that might panic
}

// Logging Entry/Exit
func LogExample() {
	log.Println("Entered function")
	defer log.Println("Exiting function")
}

// Timers or Profiling
func Timers() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution time:", time.Since(start))
	}()
}

/*
Types of Deferred Functions
| Type               | Example                       | Description                      |
| ------------------ | ----------------------------- | -------------------------------- |
| Simple function    | `defer fmt.Println("done")`   | Executes when the function exits |
| Anonymous function | `defer func() { ... }()`      | Allows complex logic             |
| With parameters    | `defer fmt.Println("val", x)` | Arguments evaluated immediately  |
| With function call | `defer closeConn(conn)`       | Calls function with current args |

Where You Should Use defer
| Use Case             | Why Use `defer`             |
| -------------------- | --------------------------- |
| File operations      | Ensures file is closed      |
| Network sockets      | Prevents resource leaks     |
| Database connections | Ensures proper closure      |
| Goroutine safety     | Releases mutex or channel   |
| Error handling       | Useful for logging/recovery |

*/
