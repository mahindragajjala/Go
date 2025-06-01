package looponchannels

import "fmt"

//Loop with Fan-In Pattern
func Fan_in() {
	fanIn := func(ch1, ch2 <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for {
				select {
				case val := <-ch1:
					out <- val
				case val := <-ch2:
					out <- val
				}
			}
		}()
		return out
	}
	fmt.Println(fanIn)
}

/*
Use case: Merge multiple sources of input.
*/
