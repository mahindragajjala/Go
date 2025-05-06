package functionliterals

import "fmt"

func Immediately_invoked() {
	message := func(name string) string {
		return "Welcome, " + name
	}("Mahindra") // immediate invocation

	fmt.Println(message)
}
