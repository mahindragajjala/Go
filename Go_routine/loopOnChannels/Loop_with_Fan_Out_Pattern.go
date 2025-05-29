package looponchannels

import "fmt"

//Loop with Fan-Out Pattern
func Loop_with_Fan_Out_Pattern() {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go func(id int) {
			for val := range ch {
				fmt.Printf("Worker %d received %d\n", id, val)
			}
		}(i)
	}

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

/*
Use case: Distribute workload to multiple goroutines.
*/
