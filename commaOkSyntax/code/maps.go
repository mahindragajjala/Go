package code

import "fmt"

func Using_maps() {
	data := make(map[string]int)
	data["one"] = 1
	data["two"] = 2
	data["three"] = 3
	var x string = "one"
	not, ok := data[x]
	if ok {
		fmt.Println(not)
	} else {
		fmt.Println("Not found")
	}
}
