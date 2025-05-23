package anonymousfunctions

import "fmt"

/*
Returned from a Function (Closure)

A closure is a function that "captures variables from its surrounding lexical scope", even after that scope has finished execution.

In simple terms:

A closure is an inner function that remembers the variables from the outer function, even after the outer function is done.

| Use Case                      | Description                                      |
| ----------------------------- | ------------------------------------------------ |
| ✅ Stateful counters           | Each closure has its own `count`                 |
| ✅ Encapsulating configuration | Return a customized function based on config     |
| ✅ Memoization                 | Store previously computed values                 |
| ✅ Event handlers / callbacks  | Retain context during async or delayed execution |
| ✅ Middleware patterns         | Useful in web servers like Gin or Echo           |

*/

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

//Function Factory
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func Function_factory() {
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Println(double(4)) // 8
	fmt.Println(triple(4)) // 12
}

//Encapsulated Cache
func Encapsulated_Cache() {
	f := memoFib()
	fmt.Println(f(10)) // 55, but computed with memoization
}
func memoFib() func(int) int {
	cache := map[int]int{}
	var fib func(n int) int
	fib = func(n int) int {
		if val, ok := cache[n]; ok {
			return val
		}
		if n <= 1 {
			cache[n] = n
		} else {
			cache[n] = fib(n-1) + fib(n-2)
		}
		return cache[n]
	}
	return fib
}
