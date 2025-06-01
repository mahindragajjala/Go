package done

import "fmt"

/*
Pipeline Cancellation using done
When you build pipelines (stages of goroutines), passing a done to each stage ensures early cancellation of the entire pipeline.
*/
func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-done:
				fmt.Println("gen: cancelled")
				return
			case out <- n:
			}
		}
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-done:
				fmt.Println("sq: cancelled")
				return
			case out <- n * n:
			}
		}
	}()
	return out
}

func Pipe_line() {
	done := make(chan struct{})
	defer close(done) // Ensures cleanup when main exits

	numbers := gen(done, 1, 2, 3, 4)
	squares := sq(done, numbers)

	// Only take the first result, then cancel the rest
	fmt.Println(<-squares)
}
