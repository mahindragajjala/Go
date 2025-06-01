package methodexpressionandmethodvalues

/*
A method expression refers to the method without binding to a specific receiver. Instead, you pass the receiver explicitly as the first argument.

Syntax:

areaFunc := Circle.Area
fmt.Println(areaFunc(c)) // method takes 'c' as an argument
Circle.Area refers to the method as a function: func(Circle) float64

Now areaFunc is a function that expects a Circle value as the first parameter.

Receiver is explicit, unlike method values.
*/

/* func Method_expression() {
	// Method Expression
	areaExpr := Circle.Area                        // Not bound to c
	fmt.Println("Method Expression:", areaExpr(c)) // c passed as arg
}
*/

/*
Works for both pointer and value receivers.
Useful for callbacks, closures, higher-order functions.
Makes method handling flexible like functional programming.
*/
