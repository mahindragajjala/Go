package patterns

import (
	"fmt"
	"sync"
)

/* func Stream() {
	//channels
	var wg sync.WaitGroup
	channels := make([]chan int, 5)
	processchannel := make(chan int)
	singlechannel := make(chan int)
	for i := 0; i < len(channels); i++ {
		channels[i] = make(chan int)
	}
	generate(&wg, processchannel)                         //fan-out
	process(&wg, channels, processchannel, singlechannel) //processing
	combine(&wg, singlechannel)                           //fan-in
	wg.Wait()
}

func generate(wg *sync.WaitGroup, processchannel chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			processchannel <- i
		}
		close(processchannel)
	}()
}
func process(wg *sync.WaitGroup, channels []chan int, processchannel chan int, singlechannel chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range processchannel {
			go GoRoutine_func(wg, channels[v], singlechannel)
		}
		close(singlechannel)
	}()
}
func combine(wg *sync.WaitGroup, singlechannel chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range singlechannel {
			fmt.Println(v)
		}
	}()
}
func GoRoutine_func(wg *sync.WaitGroup, data chan int, singlechannel chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range data {
			v = v * 2
			singlechannel <- v
		}
	}()
} */

func Stream() {
	var wg sync.WaitGroup
	numWorkers := 5

	// Fan-out: Create a buffered channel to hold generated data
	processchannel := make(chan int)

	// Fan-in: Single output channel
	singlechannel := make(chan int)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker_routine(&wg, processchannel, singlechannel)
	}

	// Generate data
	go generate(processchannel)

	// Collect results
	go func() {
		wg.Wait()
		close(singlechannel) // Only close after all workers are done
	}()

	// Print results
	combine(singlechannel)
}

func generate(processchannel chan int) {
	for i := 0; i < 10; i++ {
		processchannel <- i
	}
	close(processchannel)
}

func worker_routine(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer wg.Done()
	for val := range in {
		out <- val * 2
	}
}

func combine(singlechannel <-chan int) {
	for val := range singlechannel {
		fmt.Println(val)
	}
}
