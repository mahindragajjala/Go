package casekeyword

import "fmt"

func CaseKeyword() {
	var data = []int{2, 5, 6, 8, 9, 10}
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case 2:
			fmt.Println(data[i])
		case 5:
			fmt.Println(data[i])
		case 6:
			fmt.Println(data[i])
		case 8:
			fmt.Println(data[i])
		case 9:
			fmt.Println(data[i])
		}
	}
}
