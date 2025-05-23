package content

/*
When Does a Variable Escape to the Heap in Go?

✅ 1. Returned as a pointer
If you return a pointer to a local variable, it escapes to the heap because the function’s stack frame will be gone after it returns.

func create() *int {
    x := 10
    return &x // x escapes to heap
}

✅ 2. Captured by a closure
If a variable is used inside a closure (a function defined within another function), and the closure is used outside that scope, the variable escapes.

func counter() func() int {
    x := 0
    return func() int {
        x++
        return x
    } // x escapes to heap
}

✅ 3. Stored in a heap-allocated object
If a value is assigned to a field of a struct, slice, or map that is on the heap, the value escapes.

type User struct {
    name *string
}

func newUser() *User {
    n := "Mahindra"
    return &User{name: &n} // n escapes
}

✅ 4. Used by a goroutine
If a variable is accessed inside a goroutine, it escapes to heap because goroutines may run after the function returns.
func run() {
    msg := "hello"
    go func() {
        fmt.Println(msg)
    }() // msg escapes to heap
}


✅ 5. Stored in an interface
If a concrete type is stored in an interface, and its size or structure is unknown at compile time, it might escape.

func save(i interface{}) {
    fmt.Println(i)
}
func run() {
    val := 123
    save(val) // val may escape
}


🧪 How to Check Escape Analysis in Go?
Run this command to see escape decisions:

go build -gcflags="-m"
Example:

func test() *int {
    x := 100
    return &x
}
Run:

$ go build -gcflags="-m"
# command-line-arguments
./main.go:3:13: &x escapes to heap
*/
