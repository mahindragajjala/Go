package questions

import "fmt"

func Loop_string() {
	var Stringdata string = "Hi siri"
	NewMap := make(map[rune]int)
	for _, value := range Stringdata {
		fmt.Println(value)
		NewMap[value]++
	}
	for key, value := range NewMap {
		fmt.Printf("key:%c,value:%d\n", key, value)
	}
}

/*
-  unicode = loop(string)
- map[unicode]int
- range[unicode]
-
*/
