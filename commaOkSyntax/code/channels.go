package code

import "fmt"

func Channels_comma_ok() {
	channel := make(chan int, 2)
	channel <- 1
	value, ok := <-channel
	if ok {
		fmt.Println("executed", value)
	} else {
		fmt.Println("not executed")
	}
	channel <- 1
	value2, ok := <-channel
	if ok {
		fmt.Println("executed", value2)
	} else {
		fmt.Println("not executed")
	}

}
