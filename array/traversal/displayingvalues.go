package traversal

import "fmt"

func Displaying_value(array [5]int) {
	for key, value := range array {
		fmt.Printf("index:%v , value:%v", key, value)
	}

}
