package looponchannels

import "fmt"

func Using_range() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // required to exit range
	}()

	for val := range ch {
		fmt.Println(val)
	}

}
