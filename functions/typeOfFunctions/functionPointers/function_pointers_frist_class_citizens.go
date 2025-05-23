package functionpointers

/*
In Go, functions are first-class citizens, meaning:

You can assign a function to a variable
You can pass a function as a parameter
You can return a function from another function
You can store functions in data structures like slices or maps
This concept is similar to function pointers in C/C++, but with type safety and ease of use in Go.

Great! You're diving into an important Go concept: functions as first-class citizens, or what’s commonly known as function pointers in other languages like C/C++.

Your Code Explained

func greet(name string) {
    fmt.Println("Hello", name)
}
This defines a named function called `greet` that takes a string `name` and prints a greeting.

var sayHello func(string) = greet

What happens here?

* `sayHello` is declared as a variable of type `func(string)`, which means it can point to any function that takes a `string` and returns nothing (`void` in C terms, or `func(string)` in Go).
* You then assign the `greet` function to `sayHello`.

So, `sayHello` now acts like a "function pointer" to `greet`.

sayHello("Mahi")
This calls the function via the pointer (alias) `sayHello`. Output will be:
Hello Mahi

🧠 Behind the Scenes

In memory terms, the function `greet` exists at a specific address. By assigning `sayHello = greet`, you’re essentially telling Go:

> “Let `sayHello` point to the same code block that `greet` points to.”

But unlike C, Go doesn’t expose raw memory addresses — it does this with type-safe function references.

 🧪 Real-Time Example

Let’s say you want to create a function that takes another function as a parameter:
func execute(fn func(string), name string) {
    fn(name)
}

func greet(name string) {
    fmt.Println("Hello", name)
}

func main() {
    execute(greet, "Mahi")  // Output: Hello Mahi
}


Here, the function `greet` is passed as a parameter to `execute`, which then calls it.



 🔄 Summary

| Concept                  | Description                                                |
|  | - |
| Function as variable     | Store functions in variables: `var f func(string) = greet` |
| Function as argument     | Pass functions to other functions                          |
| Function as return value | Return functions from other functions                      |
| Type-safe                | Go enforces function signatures, unlike raw C pointers     |

*/
