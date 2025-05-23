package closure

/*
Imagine you're a manager (outer function) setting up a task tracker (closure) for each of your team members.

Scenario:
You assign a whiteboard to each employee to track their task count.

Even after you (the manager) leave the room, your employees keep using their whiteboards to update their tasks.

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func Closure() {
	next := counter()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
}


| Real-world Component      | Go Equivalent                        |
| ------------------------- | ------------------------------------ |
| Manager                   | `counter()` function                 |
| Whiteboard                | `count` variable                     |
| Employee                  | Closure function (returned function) |
| Task Update               | `count++`                            |
| Whiteboard kept after mgr | `count` escapes to heap              |


“Hey, this variable is still needed even after the function finishes — I’ll move it from stack (temporary) to heap (persistent).”


check the heap data by the command - go run -gcflags="-m" main.go

Without closure:
Local variables live on the stack.
They’re destroyed once the function ends.

With closure:
If a local variable is accessed by a returned function, it escapes.
It’s stored on the heap so it can live as long as the closure does.


In Go, you don't explicitly move a variable from the stack to the heap.

Instead, the Go compiler automatically decides where to allocate memory based on escape analysis.

However, you can influence this decision by writing code that causes a variable to escape — typically by:

Returning a pointer to a local variable
Capturing a variable in a closure
Passing a pointer to a goroutine
Storing a pointer in a longer-lived structure
*/

//Returning a pointer to a local variable
func CreatePointer() *int {
	x := 10   // x is on the stack
	return &x // x escapes to heap because its address is returned
}

/*
🧠 Explanation: x would normally be on the stack, "but since its address is returned", it must survive after createPointer() returns, so Go moves it to the heap.
*/

//Capturing in a closure
func Capture_counter() func() int {
	count := 0 // captured by the closure
	return func() int {
		count++
		return count
	}
}

//Assigning a local pointer to a global variable
var global *int

func SetGlobal() {
	x := 42
	global = &x // x escapes to heap because global lives longer
}
