package initfunction

import "fmt"

/*
init() is a special function in Go that runs automatically before the main() function or any test.

It’s used for initialization — setting up state, loading configs, connecting to databases, etc.
*/
func init() {
	fmt.Println("Init called")
}

/*
| Property                 | Description                                      |
| ------------------------ | ------------------------------------------------ |
| Automatically called     | No need to call `init()` manually.               |
| No parameters, no return | Signature is always `func init()`                |
| Runs once per file       | Each Go file can have **its own** `init()`       |
| Runs before `main()`     | Used to prepare the program state                |
| Used with `import`       | If a package has `init()`, it runs when imported |

*/
/*
Multiple init() functions

func init() {
    fmt.Println("Init 1")
}

func init() {
    fmt.Println("Init 2")
}

output :
Init 1
Init 2

Order of Execution (Important!)

Imported packages' init() functions run first (in the import order).
Then init() functions in the current package.
Then the main() function.
*/

/*
Real-Life Use Cases for init()


| Use Case                 | Example                                                |
| ------------------------ | ------------------------------------------------------ |
| Load config from file    | `init()` reads a JSON/YAML file and sets global config |
| Establish DB connections | `init()` sets up a global DB handle                    |
| Register plugins/modules | Common in frameworks                                   |
| Setup logging            | `init()` configures logger once                        |
| Auto-run setup on import | Useful for libraries/packages                          |

*/

/*
What init() Should NOT Be Used For
Business logic
Heavy computation
Anything order-dependent across packages (can lead to unexpected results)
*/

/*
 Runs once per file
What it means:
Every .go file in a package can have its own init(), and Go will run each one exactly once — in file order determined by the compiler (not necessarily in the order you expect).

Example:

📁 main.go

package main

import "fmt"

func init() {
    fmt.Println("Init from main.go")
}

func main() {
    fmt.Println("Main")
}

📁 config.go

package main

import "fmt"

func init() {
    fmt.Println("Init from config.go")
}
Output:
Init from config.go
Init from main.go
Main

🧠 Both init() functions are executed, each only once, before main().

*/

/*
Used with import
What it means:
When you import a package, its init() function runs, even if you don’t call anything from that package.

This is useful in libraries that auto-register, like database drivers or plugins.

📁 mathlib/math.go
package mathlib

import "fmt"

func init() {
    fmt.Println("mathlib init called")
}

📁 main.go
package main

import (
    _ "yourmodule/mathlib" // Only importing for its init()
    "fmt"
)

func main() {
    fmt.Println("Main starts")
}
Output:

mathlib init called
Main starts
🧠 Even though nothing from mathlib is called in main.go, the init() inside it still runs.


| Use `init()` For              | Why                             |
| ----------------------------- | ------------------------------- |
| Load environment/config       | Needed before `main()` runs     |
| Initialize connections        | e.g., DB, cache, API setup      |
| Register plugins/modules      | Auto-initialization             |
| Logging setup                 | Consistent formatting/log level |
| Auto-code execution on import | Used in frameworks, drivers     |

*/
