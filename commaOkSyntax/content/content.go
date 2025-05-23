package content

/*
when you access a map or work with channels or type assertions
Exactly!

Definition:

In Go, when you access a map or work with channels or type assertions, you can use:
value, ok := something

The `ok` is a boolean — short for "okay" — that tells you whether the operation was successful.

Examples in Go

Map Lookup:

capital, ok := Country["India"]
if ok {
    fmt.Println("The capital is:", capital)
} else {
    fmt.Println("Country not found!")
}

Channel Receive:

ch := make(chan int, 1)
ch <- 10
value, ok := <-ch
fmt.Println(value, ok)  // Output: 10 true

value2, ok2 := <-ch
fmt.Println(value2, ok2)  // Output: 0 false  (channel empty)

Type Assertion:

var x interface{} = "Hello"
str, ok := x.(string)
if ok {
    fmt.Println("String value:", str)
} else {
    fmt.Println("Not a string!")
}

Summary Table:

| Use case          | Syntax                        | Purpose                                 |
|-------------------|--------------------------------|-----------------------------------------|
| Map lookup        | `value, ok := map[key]`        | Checks if key exists.                  |
| Channel receive   | `value, ok := <- channel`      | Checks if channel is closed.           |
| Type assertion    | `value, ok := x.(Type)`        | Checks if interface holds this type.   |

*/
