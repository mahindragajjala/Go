package variadicfunctions

import "fmt"

/*
Arguments can be any type, often used with interface{}
*/
func Heterogeneous_variadic(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
