package content

import "fmt"

func Types_of_functions() {

}

/*
	No parameters.
	No return value.
*/
// func()

/*
Takes one parameter x of type int.
Returns an int
*/
func OneinpuParameter_oneOutputparameter(x int) int {
	return 1
}

/*
Takes:
a (int),
_ (ignored, int),
z (float32).

Returns a bool.
*/
func checkPositive(a, _ int, z float32) bool {
	return a > 0 && z > 0
}
func Empty_parameter() {
	fmt.Println(checkPositive(5, 100, 3.14)) // Output: true
}

//func(a, b int, z float32) (bool)
func isSumGreaterThanZ(a, b int, z float32) bool {
	sum := float32(a + b)
	return sum > z
}
func Multiple_parameter_with_same_type() {
	fmt.Println(isSumGreaterThanZ(3, 4, 6.5)) // Output: true
}

//func(prefix string, values ...int)
func printSum(prefix string, values ...int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Printf("%s %d\n", prefix, sum)
}

//Variadic_parameter
func Inserting_multiple_inputs() {
	printSum("Total:", 1, 2, 3, 4) // Output: Total: 10
}

//func(a, b int, z float64, opt ...interface{}) (success bool)
func process(a, b int, z float64, opt ...interface{}) (success bool) {
	fmt.Println("Required:", a, b, z)
	fmt.Println("Optional:", opt)
	return true
}

func Unlimited_parameters_interface() {
	status := process(10, 20, 30.5, "extra", 99, true)
	fmt.Println("Success:", status) // Output: Success: true
}

//func(int, int, float64) (float64, *[]int)
func compute(a, b int, factor float64) (float64, *[]int) {
	result := float64(a+b) * factor
	slice := []int{a, b}
	return result, &slice
}
func Slice_Pointer_function() {
	res, nums := compute(2, 3, 1.5)
	fmt.Println("Result:", res) // Output: 7.5
	fmt.Println("Nums:", *nums) // Output: [2 3]
}

//func(n int) func(p *T)
type T struct {
	Name string
}

func greeter(n int) func(p *T) {
	return func(p *T) {
		for i := 0; i < n; i++ {
			fmt.Println("Hello,", p.Name)
		}
	}
}

func Return_function() {
	person := &T{Name: "Mahindra"}
	greetFunc := greeter(3) // returns a function
	greetFunc(person)       // Output: Hello, Mahindra (3 times)
}

//Anonymous Functions
//Functions without a name, often used inline.

func Anonymous_Functions() {
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Go Developer")
}

//Higher-Order Functions
// Functions that take other functions as arguments or return them.
//syntax
func Higher_Order_Functions(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

//example
// Function that takes another function as an argument
func apply(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

// Simple add function
func add(x int, y int) int {
	return x + y
}

// Simple multiply function
func multiply(x int, y int) int {
	return x * y
}

func Higher_Order_Functions_main() {
	result1 := apply(add, 10, 5)      // Calls add(10, 5)
	result2 := apply(multiply, 10, 5) // Calls multiply(10, 5)

	fmt.Println("Sum:", result1)
	fmt.Println("Product:", result2)
}

//Closures
func Closures() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//Recursion
//Functions that call themselves.

func Recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * Recursion(n-1)
}

//Method (Function on Struct)
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

//Defer, Panic, and Recover

//Using defer to delay function execution, useful for cleanup.
func Defer() {
	defer fmt.Println("Executed at the end!")
	fmt.Println("First Line")
}
