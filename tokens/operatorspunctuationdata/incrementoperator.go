package operatorspunctuationdata

import "fmt"

/*
In Go, the ++ operator is commonly
used in loops and in places where
you need to increment a counter
or index variable.
*/
func Increment_operator() {
	// In for loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Incrementing a counter
	count := 0
	for _, char := range "hello world" {
		if char == 'l' {
			count++
		}
	}
	fmt.Println("Number of l's:", count)

	//While-style loops with for
	i := 0
	for i < 5 {
		fmt.Println("i =", i)
		i++
	}

	//Indexing through slices or arrays
	arr := []int{10, 20, 30, 40}
	ij := 0
	for ij < len(arr) {
		fmt.Println(arr[ij])
		ij++
	}

	/*
		Where you can’t use ++

		Not allowed: x := i++ ❌

		Not allowed: ++i ❌

		You can’t use it inside expressions, only as a standalone statement.
	*/
}
