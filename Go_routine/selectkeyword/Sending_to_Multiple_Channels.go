package selectkeyword

import "fmt"

func Sending_to_Multiple_Channels() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	select {
	case ch1 <- "Sending to ch1":
		fmt.Println("Sent to ch1")
	case ch2 <- "Sending to ch2":
		fmt.Println("Sent to ch2")
	}
}
