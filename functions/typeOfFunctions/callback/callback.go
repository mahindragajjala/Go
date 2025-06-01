package callback

import "fmt"

/*
In Go (Golang), a callback function is a function that is passed as an argument to another function and is called (executed) inside that function. It allows you to customize behavior without changing the actual function logic.
*/
func Process(callback func(string)) {
	callback("Processing done")
}
func Callback_function() {
	Process(func(msg string) {
		fmt.Println("Callback received:", msg)
	})
}

/*
process is a function that accepts another function as a parameter.
This parameter is named "callback".

The type of this parameter is: func(string) — which means the callback must be a function that takes a string as input and returns nothing (void in other languages, nil in Go).

Inside process, we call the callback function, passing it the string "Processing done".
*/
/*
1. No input, no output
func(callback func()) {
    callback()
}

2. Takes input, no output
func(callback func(int, int)) {
    callback(10, 20)
}

3. Takes input, returns output
func(callback func(int, int) int) {
    result := callback(5, 3)
    fmt.Println("Result:", result)
}

4. Returns error
func(callback func(string) error) {
    err := callback("Hello")
    if err != nil {
        fmt.Println("Error:", err)
    }
}

5. Multiple return values
func(callback func(int, int) (int, error)) {
    val, err := callback(10, 5)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Value:", val)
    }
}

*/
