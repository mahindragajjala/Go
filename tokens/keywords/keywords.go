package keywords

import "fmt"

/*
break - for/switch/select
default -  switch, select
func - functions
interface - To store the any data in the interface data type
select - channels/goroutines
case - Switch/Select
defer - functions
go - the go keyword is used to start a new goroutine.
map - In Go, map is a built-in data type that represents a collection of key-value pairs.
struct - the struct keyword is used to define a structure, which is a composite data type
		 that groups together zero or more fields (variables) under one name.
chan - creation of channel
else - or else
goto - the goto keyword is used to jump to a labeled statement within the same function.
	   It's like saying: “Hey, skip all the lines in between and go directly to that label.”
package - declare the file name
		Executable programs	package main	Starts execution from main()
		Libraries/packages	package <name>	Groups reusable code
switch - conditional branching—it's a cleaner alternative to multiple if-else statements
const - Cannot change value	Constants are immutable
        Must be compile-time computable	You can't use a function like rand() for const
        No := shorthand	Use const name = value, not :=
iota - iota is a predeclared identifier representing successive untyped
	   integer constants within a const block.
fallthrough -   In Go, fallthrough is a keyword used in a switch statement to force the execution
				to continue to the next case, regardless of whether that case matches the condition.
if - conditional statements
range - the range keyword is used in loops to iterate over elements in arrays, slices, maps, strings, and channels.
type - struct/interface/named type/type alias/function type
continue - it's used inside loops and is pretty straightforward once you get the hang of it. or for skipping the loop
for - looping statements/as while/infinite/Looping Over Arrays, Slices, Maps(using range)
import - basic/multiple/alias/blankIdentifier/third-party packages
return - back from a function - Basic Usage of return/Multiple Return Values/Named Return Values/Use with Early Exit
var - variable intialization
*/
func Keywords() {
	var nums = []int{1, 2, 3, 4, 5}

	for i, val := range nums {
		if val%2 == 0 {
			fmt.Println("Even:", val)
		} else {
			fmt.Println("Odd:", val)
		}

		if i == 3 {
			break // stops the loop at index 3
		}
	}
}
