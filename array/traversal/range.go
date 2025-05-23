package traversal

import "fmt"

func Using_range(array [5]int) {
	for key, value := range array {
		fmt.Println(key, value)
	}
}
