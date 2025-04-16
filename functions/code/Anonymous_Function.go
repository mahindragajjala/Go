package code

import "fmt"

func Anonymous_Functions() {
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Go Developer")
}
