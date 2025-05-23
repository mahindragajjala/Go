package nestedfunctions

import "fmt"

/*
Nested function in Golang
In Go, we can create a function inside another function. This is known as a nested function.
*/
func Outer_function() {
	fmt.Println("OUTER FUNCTION")

	// inner function
	//anonymous function
	var displayName = func() int {
		fmt.Println("Hi", "inner loop")
		return 1
	}

	// call inner function
	int := displayName()
	fmt.Println(int)
}

/* In the above example, we have created an anonymous function inside the Outer_function() function.
Here, var displayName = func() {...} is a nested function.
The nested function works similar to the normal function.
It executes when displayName() is called inside the function greet().
*/

/*
Returning a function in Go
We can create a function that returns an anonymous function. For example,
*/

func greet() func() {

	return func() {
		fmt.Println("Hi John")
	}

}

func Returning_anonymous_function() {
	g1 := greet()
	g1()
}

/*
 Here, func() before the curly braces indicates that the function returns another function.

Also, if you look into the return statement of this function, we can see the function is returning a function.

From mReturning_anonymous_functionain(), we call the greet() function.

Here, the returned function is now assigned to the g1 variable. Hence, g1() executes the nested anonymous function.
*/

/* -----------------------Go Closures-------------------------- */
/*
Go closure is a nested function that allows us to access variables of the outer function even after the outer function is closed.

Before we learn about closure, let's first revise the following concepts:

Nested Functions
Returning a function

As we have already discussed, closure is a "nested function" that helps us "access the outer function's variables even after the outer function is closed". Let's see an example.



*/
func greet_closure() func() string {

	// variable defined outside the inner function
	name := "John"

	// return a nested anonymous function
	return func() string {
		name = "Hi " + name
		return name
	}
}

/*
A closure is a function (often anonymous) that captures and uses variables from its enclosing (outer) function scope.
Storing such a function in a variable allows you to use the closure later, along with its captured state.
*/
func Closure() {

	// call the outer function
	message := greet_closure()

	// call the inner function
	fmt.Println(message())
}
